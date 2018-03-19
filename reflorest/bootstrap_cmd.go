package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const bootstrapDataSource = "https://raw.githubusercontent.com/sanksons/reflorest/master/reflorest/_newApp"

var ApplicationPath string
var ApplicationName string

var Rmap map[string]string

func BuildBootstrapCommand() *Command {

	flagSet := flag.NewFlagSet("bootstrap", flag.ExitOnError)
	return &Command{
		Name:         "bootstrap",
		FlagSet:      flagSet,
		UsageCommand: "reflorest bootstrap <Application PATH>",
		Usage: []string{
			"Bootstrap a new Application",
		},
		Command: func(args []string, additionalArgs []string) {
			fmt.Printf("%s Bootstrapping New Application.%s\n", greenColor, defaultStyle)
			err := generateBootstrap(args)
			if err != nil {
				fmt.Printf("%s Error Occurred: %s. %s\n", redColor, err.Error(), defaultStyle)
				return
			}
			fmt.Printf("%s Finished.%s\n", greenColor, defaultStyle)
		},
	}
}

func generateBootstrap(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("You need to supply application path.")
	}
	ApplicationPath = args[0]
	ApplicationName = prepareApplicationName(ApplicationPath)

	Rmap = initiateReplacementMap()
	return createDirStructure()
}

func prepareApplicationName(path string) string {
	pathArr := strings.Split(path, string("/"))
	fmt.Printf("Path:%s\n", path)
	return pathArr[len(pathArr)-1]
}

type Path string

type DS struct {
	root *Folder
}

func (this *DS) Fetch() error {

	this.root.Remote = bootstrapDataSource
	wd, _ := os.Getwd()
	this.root.Local = Path(wd)
	return this.root.BootUp(true)
}

func (this *DS) Show() {
	var show func(Folder, string)
	show = func(folder Folder, spaces string) {
		//print files
		if len(folder.Files) > 0 {
			for _, file := range folder.Files {
				file.PrintIt(spaces)
			}
		}
		if len(folder.Folders) > 0 {
			for _, folder := range folder.Folders {
				folder.PrintIt(spaces)
				show(folder, spaces+"    ")
			}
		}

	}
	show(*this.root, "    ")
}

type Folder struct {
	Name       string
	ActualName string
	Files      []File
	Folders    []Folder
	Local      Path
	Remote     Path
	Spaces     int
}

func (this *Folder) BootUp(skip bool) error {

	err := this.create(skip)
	if err != nil {
		return this.error(err)
	}
	err = this.createFiles()
	if err != nil {
		return this.error(err)
	}
	if this.Folders == nil {

		err := os.Chdir("../")
		if err != nil {
			return this.error(err)
		}
		return nil
	}
	for _, folder := range this.Folders {
		folder.Local = this.Local + Path(os.PathSeparator) + Path(folder.ActualName)
		folder.Remote = this.Remote + "/" + Path(folder.ActualName)
		err := folder.BootUp(false)
		if err != nil {
			return err
		}

	}
	err1 := os.Chdir("../")
	if err1 != nil {
		return this.error(err1)
	}

	return nil
}

func (this *Folder) create(skip bool) error {

	if skip {
		return nil
	}
	//wd, _ := os.Getwd()
	//fmt.Printf("making %s in %s\n", this.ActualName, wd)
	err := os.Mkdir(this.ActualName, 0755)
	if err != nil {
		return this.error(err)
	}
	err = os.Chdir(this.ActualName)
	if err != nil {
		return this.error(err)
	}

	return nil
}

func (this *Folder) error(err error) error {
	return err
}

func (this *Folder) createFiles() error {
	if len(this.Files) <= 0 {
		return nil
	}
	for _, file := range this.Files {
		file.Local = this.Local + Path(os.PathSeparator) + Path(file.GetFileName())
		file.Remote = this.Remote + "/" + Path(file.GetFileName())
		err := file.Create()
		if err != nil {
			return this.error(err)
		}
	}
	return nil
}

func (this *Folder) PrintIt(spaces string) {
	fmt.Println(spaces + this.ActualName + "/")
}

type File struct {
	Name       string
	ActualName string
	Extension  string
	Local      Path
	Remote     Path
}

func (this *File) PrintIt(spaces string) {
	fmt.Println(spaces + this.GetFileName())
}

func (this *File) error(err error) error {
	return fmt.Errorf("File:%s, Location:%s, Error:%s\n", this.ActualName, this.Local, err.Error())
}

func (this *File) Fetch() ([]byte, error) {
	var maxretry int = 3
	var retry int
	var err error
	var resp *http.Response
	var returnData []byte

	for retry < maxretry {
		retry++
		resp, err = http.Get(string(this.Remote))
		if err != nil {
			continue
		}
		if resp.StatusCode != 200 {
			resp.Body.Close()
			err = fmt.Errorf("Got HTTP Status code: %d", resp.StatusCode)
			continue
		}
		b := make([]byte, 0)
		buf := bytes.NewBuffer(b)
		_, err := buf.ReadFrom(resp.Body)
		if err != nil {
			resp.Body.Close()
			continue
		}
		resp.Body.Close()
		returnData = buf.Bytes()
	}
	return returnData, err
}

func (this *File) GetFileName() string {
	filename := this.ActualName
	if this.Extension != "" {
		filename = this.ActualName + "." + this.Extension
	}
	return filename
}

func (this *File) Create() error {

	destFile, err := os.Create(this.GetFileName())
	if err != nil {
		return this.error(err)
	}
	defer destFile.Close()

	//fmt.Printf("Hitting: %s\n", this.Remote)
	data, err := this.Fetch()
	if err != nil {

		return this.error(err)
	}
	//{{APP_PATH}}
	for k, v := range Rmap {
		//	fmt.Println(k)
		data = []byte(strings.Replace(string(data), k, v, -1))
	}
	_, err = destFile.Write(data)
	if err != nil {
		return this.error(err)
	}
	destFile.Sync()
	return nil
}

func initiateReplacementMap() map[string]string {

	s := make(map[string]string)

	s["{{APP_PATH}}"] = ApplicationPath
	s["{{APP_NAME}}"] = ApplicationName

	//prepare log path
	var logpath string
	if runtime.GOOS == "windows" {
		//hack for windows
		logpath = "C:\\\\" + ApplicationName + "\\\\"
	} else {
		logpath = "/var/log/" + ApplicationName + "/"
	}

	s["{{LOG_PATH}}"] = logpath

	os.MkdirAll(logpath, 0777)
	return s
}

func createDirStructure() error {
	//Define directory structure
	fmt.Printf("Creating Directory Structure:\n")
	root := Folder{
		Name:       "root",
		ActualName: "root",
		Files: []File{
			File{Name: "main", ActualName: "main", Extension: "go"},
		},
		Folders: []Folder{
			Folder{
				Name:       "conf",
				ActualName: "conf",
				Files: []File{
					File{Name: "conf", ActualName: "conf", Extension: "json"},
					File{Name: "logger", ActualName: "logger", Extension: "json"},
					File{Name: "standard", ActualName: "standard", Extension: "flf"},
				},
			},
			Folder{
				Name:       "src",
				ActualName: "src",
				Folders: []Folder{
					Folder{
						Name:       "common",
						ActualName: "common",
						Folders: []Folder{
							Folder{Name: "appconfig", ActualName: "appconfig", Files: []File{File{Name: "config", ActualName: "application_config", Extension: "go"}}},
							Folder{Name: "appconstant", ActualName: "appconstant", Files: []File{File{Name: "errcodes", ActualName: "error_codes", Extension: "go"}}},
						},
					},
					Folder{
						Name:       "hello",
						ActualName: "hello",
						Files: []File{
							File{Name: "apidef", ActualName: "api_definition", Extension: "go"},
							File{Name: "datastruct", ActualName: "data_structures", Extension: "go"},
							File{Name: "hello", ActualName: "hello_world", Extension: "go"},
							File{Name: "hellohealth", ActualName: "hello_world_health_checker", Extension: "go"},
							File{Name: "swagger", ActualName: "swagger", Extension: "go"},
						},
					},
					Folder{Name: "test", ActualName: "test"},
				},
			},
		},
	}
	//Fetch and put on system
	ds := DS{root: &root}
	err := ds.Fetch()
	if err == nil {
		ds.Show()
	}
	return err
}
