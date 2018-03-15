package formatter

import (
	"fmt"

	"github.com/sanksons/reflorest/src/common/logger/message"
)

type consoleFormat struct {
}

const greenColor = "\x1b[32m"
const redColor = "\x1b[91m"
const defaultStyle = "\x1b[0m"
const lightGrayColor = "\x1b[37m"

//formatString is format of the log in string formattype configuration
//var formatString = "[level : %s, message : %s, tId : %s, reqId : %s, appId : %s, sessionId : %s, userId : %s, stackTraces : %s, timestamp : %s, uri : %s]"

//var formatString = ""

//GetFormattedLog returns formatted log as a string interface
func (sf *consoleFormat) GetFormattedLog(msg *message.LogMsg) interface{} {

	fmt.Sprintf("%s", msg.TimeStamp, msg.Level)

	return fmt.Sprintf(formatString, msg.Level,
		msg.Message,
		msg.TransactionID,
		msg.RequestID,
		msg.AppID,
		msg.SessionID,
		msg.UserID,
		msg.StackTraces,
		msg.TimeStamp,
		msg.URI)
}

func (this *consoleFormat) getTime(time string) string {
	return fmt.Sprintf("%s %s %s", lightGrayColor, time, defaultStyle)
}

func (this *consoleFormat) getLevel(level string) string {
	if level == "error" {
		return fmt.Sprintf("%s ERR %s", redColor, defaultStyle)
	} else if level == "warning" {
		return fmt.Sprintf("%s WRN %s", redColor, defaultStyle)
	}

}
