package writers

import (
	"github.com/sanksons/reflorest/src/common/logger/formatter"
	"github.com/sanksons/reflorest/src/common/logger/message"
)

type LogWriter interface {
	Write(msg *message.LogMsg)
	SetFormatter(formatter.FormatInterface)
}
