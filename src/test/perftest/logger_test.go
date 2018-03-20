package perftest

import (
	"testing"

	"github.com/sanksons/reflorest/src/common/logger"
)

func initTestLogger(confFile string) {
	logger.Initialise(confFile)
}

func BenchmarkFlorestSyncLogger(b *testing.B) {
	initTestLogger("testdata/testloggerAsync.json")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("INFO MESSAGE")
	}
}

func BenchmarkFlorestAsyncLogger(b *testing.B) {
	initTestLogger("testdata/testloggerSync.json")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("INFO MESSAGE")
	}
}
