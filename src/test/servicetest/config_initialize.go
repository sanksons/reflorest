package servicetest

import (
	"github.com/sanksons/reflorest/src/common/config"
	"github.com/sanksons/reflorest/src/core/service"
)

func initTestConfig() {
	service.RegisterConfig(new(TestAPPConfig))

	cm := new(service.ConfigManager)
	cm.InitializeGlobalConfig("testdata/testconf.json")
	cm.UpdateConfigFromEnv(config.GlobalAppConfig, "global")

}
