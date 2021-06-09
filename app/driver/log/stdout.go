package log

import "fmt"

type StdoutLogger struct {
}

func NewStdoutLogger() *StdoutLogger {
	return &StdoutLogger{}
}

func (s StdoutLogger) Write(o ...interface{}) {
	fmt.Println(o)
}
