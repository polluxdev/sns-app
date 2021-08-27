package infrastructure

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/polluxdev/sns-app/app/usecases"
)

type Logger struct{}

func NewLogger() usecases.Logger {
	return &Logger{}
}

func (l *Logger) LogError(format string, v ...interface{}) {
	PrintLog("error", format, v...)
}

func (l *Logger) LogAccess(format string, v ...interface{}) {
	PrintLog("access", format, v...)
}

func PrintLog(status string, format string, v ...interface{}) {
	filepath := path.Join(path.Base("./log"), fmt.Sprintf("%s.log", status))

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("%s", err)
	}
	defer file.Close()

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format, v...)
}
