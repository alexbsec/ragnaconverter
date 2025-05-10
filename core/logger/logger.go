package logger

import (
	"log"
	"os"
)

type ConsoleLogger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{
		debug: log.New(os.Stdout, "[DEBUG]: ", log.LstdFlags),
		info:  log.New(os.Stdout, "[INFO]: ", log.LstdFlags),
		warn:  log.New(os.Stdout, "[WARN]: ", log.LstdFlags),
		err:   log.New(os.Stderr, "[ERROR]: ", log.LstdFlags),
	}
}

func (cl *ConsoleLogger) Debug(v ...any) {
	cl.debug.Println(v...)
}

func (cl *ConsoleLogger) Info(v ...any) {
	cl.info.Println(v...)
}

func (cl *ConsoleLogger) Warn(v ...any) {
	cl.warn.Println(v...)
}

func (cl *ConsoleLogger) Error(v ...any) {
	cl.err.Println(v...)
}
