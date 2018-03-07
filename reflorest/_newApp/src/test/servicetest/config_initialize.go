package servicetest

import (
	"github.com/sanksons/reflorest/src/core/service"
)

func initTestConfig() {
	cm := new(service.ConfigManager)
	cm.InitializeGlobalConfig("testdata/testconf.json")
}
