package log

import (
	"fmt"
	"time"
)

type StdoutLogger struct {
}

func NewStdoutLogger() *StdoutLogger {
	return &StdoutLogger{}
}

func (s StdoutLogger) Write(o ...interface{}) {
	fmt.Print("[" + time.Now().Format("2006-01-02 15:04:05") + "] ")
	fmt.Println(o...)
}
