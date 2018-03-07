package servicetest

import (
	"github.com/sanksons/reflorest/src/core/common/env"
	"github.com/sanksons/reflorest/src/core/service"
	"github.com/sanksons/reflorest/src/test/api"
)

func InitializeTestService() {

	service.RegisterAPI(new(api.TestAPI))

	reqVNAPI := new(api.ReqVNAPI)
	reqVNAPI.SetVersion("GET", "V1", "REQVD", "")
	service.RegisterAPI(reqVNAPI)

	testRateLimitedAPI := new(api.TestRateLimitedAPI)
	testRateLimitedAPI.SetVersion("GET", "V1", "TESTRATE", "")
	service.RegisterAPI(testRateLimitedAPI)

	env.GetOsEnviron()

	initTestConfig()

	initTestLogger()

	service.InitMonitor()

	service.InitVersionManager()

	service.InitCustomAPIInit()

	service.InitApis()

	service.InitHealthCheck()

	initialiseTestWebServer()

}

func PurgeTestService() {

}
