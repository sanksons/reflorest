package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/sanksons/reflorest/src/common/config"
	"github.com/sanksons/reflorest/src/common/constants"
	"github.com/sanksons/reflorest/src/common/logger"
	"github.com/sanksons/reflorest/src/common/ratelimiter"
	utilhttp "github.com/sanksons/reflorest/src/common/utils/http"
	workflow "github.com/sanksons/reflorest/src/core/common/orchestrator"
	"github.com/sanksons/reflorest/src/core/common/versionmanager"
)

type Webserver struct {
}

func (ws Webserver) ServiceHandler(w http.ResponseWriter, req *http.Request) {

	io, derr := GetData(req)
	if derr != nil {
		fmt.Fprintf(w, "Error %v", derr)
		return
	}

	serviceVersion, _, _, gerr := versionmanager.Get("SERVICE", "V1", "GET", constants.OrchestratorBucketDefaultValue, "")

	if gerr != nil {
		fmt.Fprintf(w, "Error %v", gerr)
		return
	}

	if serviceOrchestrator, ok := serviceVersion.(workflow.Orchestrator); ok {
		output := serviceOrchestrator.Start(io)
		response, _ := output.IOData.Get(constants.APIResponse)
		if v, ok := response.(utilhttp.APIResponse); ok {
			//logger.Error(fmt.Sprintf("HEllo %+v", v.Headers))
			for key, val := range v.Headers {
				w.Header().Set(key, val)
			}
			w.WriteHeader(int(v.HTTPStatus))
			w.Write(v.Body)
			return
		}
	}

	w.Header().Set("Content-Type", "application/txt")
	w.Write([]byte("Error"))
}

func (ws Webserver) Start() {

	//BootStrap the Application
	initMgr := new(InitManager)
	initMgr.Execute()

	logger.Info(fmt.Sprintln("Web server Initialization done"))

	//All requests will be passed to the service handler
	var httpHandlerFunc = utilhttp.MakeGzipHandler(ws.wrapperHandler)

	if config.GlobalAppConfig.AppRateLimiterConfig != nil {
		//Rate Limit the App
		rl, rerr := ratelimiter.New(config.GlobalAppConfig.AppRateLimiterConfig)
		if rerr != nil {
			logger.Error(fmt.Sprintln("Could not initialise rate limiter ", rerr.Error()))
			panic(rerr)
		}

		//@todo: Its overkill to always gzip when the respons size is very small.
		// this should be on the user to decide.
		httpHandlerFunc = utilhttp.MakeGzipHandler(
			ratelimiter.MakeRateLimitedHTTPHandler(ws.wrapperHandler, rl, "SERVICE"),
		)
	}

	http.HandleFunc("/", httpHandlerFunc)

	//Start the web server
	url := ":" + config.GlobalAppConfig.ServerPort

	logger.Info(fmt.Sprintln("Web server Starting......"))

	ws.displayConfigOnCli()
	ws.displayLogoOnCli()
	fmt.Printf("\nWeb server Starting......on port %s\n", config.GlobalAppConfig.ServerPort)

	serr := http.ListenAndServe(url, nil)
	if serr != nil {
		log.Printf("Could not start web server %s\n", serr.Error())
		logger.Error(fmt.Sprintln("Could not start web server ", serr))
	}
	if serr == nil {
		log.Printf("Web server Started on port %v\n", config.GlobalAppConfig.ServerPort)
		logger.Info(fmt.Sprintln("Web server Started on port : ", config.GlobalAppConfig.ServerPort))
	}

}

// wrapper handler
func (ws Webserver) wrapperHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", swaggerAllowedHeaders)
	w.Header().Set("content-type", "application/json")
	if strings.HasPrefix(r.URL.Path, "/swagger") {
		ws.swaggerHandler(w, r)
	} else {
		ws.ServiceHandler(w, r)
	}
}

// swagger handler
func (ws Webserver) swaggerHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir(".")).ServeHTTP(w, r)
}

func (ws Webserver) displayLogoOnCli() {
	file, _ := os.Open("conf/standard.flf")

	myFigure := figure.NewFigureWithFont(config.GlobalAppConfig.AppName, file, false)
	myFigure.Print()
}

func (ws Webserver) displayConfigOnCli() {
	fmt.Printf("\n\nUsing Configuration: \n")
	fmt.Printf("%s\n", config.GlobalAppConfig.ShowConfig())
}
