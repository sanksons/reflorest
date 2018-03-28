package common

import (
	"fmt"

	"github.com/sanksons/reflorest-testapp/src/common/appconfig"
	"github.com/sanksons/reflorest-testapp/src/common/appconstant"
	"github.com/sanksons/reflorest-testapp/src/hello"
	"github.com/sanksons/reflorest/src/core/service"
)

//main is the entry point of the florest web service

func StartServer() {
	fmt.Println("APPLICATION BEGIN")
	webserver := new(service.Webserver)
	Register()
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
	service.RegisterConfig(new(appconfig.AppConfig))
}

func registerErrors() {
	service.RegisterHTTPErrors(appconstant.APPErrorCodeToHTTPCodeMap)
}
