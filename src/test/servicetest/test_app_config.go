package servicetest

import (
	"errors"
	"fmt"

	"github.com/sanksons/reflorest/src/common/config"
	"github.com/sanksons/reflorest/src/common/logger"
	"github.com/sanksons/reflorest/src/components/sqldb"
)

type TestAPPConfig struct {
	ResponseHeaders config.ResponseHeaderFields `json:"ResponseHeaders,omitempty"`
	MySQL           *sqldb.SDBConfig            `json:"MySQL,omitempty"`
}

func getTestAPPConfig() (*TestAPPConfig, error) {
	c := config.GlobalAppConfig.ApplicationConfig
	testConf, ok := c.(*TestAPPConfig)
	if !ok {
		msg := fmt.Sprintf("Test APP Config Not correct %+v", c)
		logger.Error(msg)
		return nil, errors.New(msg)
	}
	return testConf, nil
}
