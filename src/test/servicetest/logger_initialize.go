package servicetest

import (
	"github.com/sanksons/reflorest/src/common/logger"
)

func initTestLogger() {
	logger.Initialise("testdata/testLoggerSync.json")
}
