package main

import (
	"{{APP_PATH}}/src/common/appconfig"
	"{{APP_PATH}}/src/common/appconstant"
	"fmt"
	"github.com/sanksons/reflorest/src/core/service"
	"{{APP_PATH}}/src/hello"
)

//main is the entry point of the florest web service
func main() {
	fmt.Println("APPLICATION BEGIN")
	webserver := new(service.Webserver)
	registerConfig()
	registerErrors()
	registerAllApis()
	webserver.Start()
}

func registerAllApis() {
	service.RegisterAPI(new(hello.HelloAPI))
}

func registerConfig() {
	service.RegisterConfig(new(appconfig.AppConfig))
}

func registerErrors() {
	service.RegisterHTTPErrors(appconstant.APPErrorCodeToHTTPCodeMap)
}
