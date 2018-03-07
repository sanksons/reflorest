package servicetest

import (
	"common/appconstant"
	"hello"

	"github.com/sanksons/reflorest/src/core/service"
)

func InitializeTestService() {
	service.RegisterHTTPErrors(appconstant.APPErrorCodeToHTTPCodeMap)
	service.RegisterAPI(new(hello.HelloAPI))
	initTestLogger()

	initTestConfig()

	service.InitMonitor()

	service.InitVersionManager()

	initialiseTestWebServer()

	service.InitHealthCheck()

}

func PurgeTestService() {

}
