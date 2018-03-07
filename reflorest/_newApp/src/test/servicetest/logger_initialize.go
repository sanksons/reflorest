package servicetest

import (
	"github.com/sanksons/reflorest/src/common/logger"
)

func initTestLogger() {
	logger.Initialise("testdata/testloggerAsync.json")
}
