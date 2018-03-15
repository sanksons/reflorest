package logger

import (
	"testing"
)

func presetup() {
	conf = &Config{
		AppName:       "TestApp",
		LogLevel:      5,
		Write2Console: true,
	}
	loggerImpls = make(map[string]LogInterface)
	SetConsoleLogger()
}

func TestConsoleLogger(t *testing.T) {
	presetup()
	msg := "I am a message"
	DebugSpecific(ConsoleLoggerKey, msg)
	InfoSpecific(ConsoleLoggerKey, msg)
	TraceSpecific(ConsoleLoggerKey, msg)
	ErrorSpecific(ConsoleLoggerKey, msg)
	WarningSpecific(ConsoleLoggerKey, msg)
}
