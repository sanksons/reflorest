package common

import (
	"fmt"

	"{{APP_PATH}}/src/common/appconfig"
	"{{APP_PATH}}/src/common/appconstant"
	"{{APP_PATH}}/src/hello"
	"github.com/sanksons/reflorest/src/core/service"
)

//main is the entry point of the florest web service

func StartServer() {
	fmt.Println("APPLICATION BEGIN")
	webserver := new(service.Webserver)
	Register()
	webserver.PreStart(func(){}, func(){})
	webserver.Start()
}

func Register() {
	registerConfig()
	registerErrors()
	registerAllApis()
}

func registerAllApis() {
	service.RegisterAPI(new(hello.HelloAPI))
}

func registerConfig() {
	service.RegisterConfig(new(appconfig.ApplicationConfig))
}

func registerErrors() {
	service.RegisterHTTPErrors(appconstant.APPErrorCodeToHTTPCodeMap)
}
