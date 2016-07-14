package raft

import (
	"log"
	"os"
)

const (
	Debug = 1
	Trace = 2
)

var logLevel int = 0
var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[raft]", log.Lmicroseconds)
}

func LogLevel() int {
	return logLevel
}

func SetLogLevel(level int) {
	logLevel = level
}

func warn(v ...interface{}) {
	logger.Print(v...)
}

func warnf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func warnln(v ...interface{}) {
	logger.Println(v...)
}

func debug(v ...interface{}) {
	if logLevel >= Debug {
		logger.Print(v...)
	}
}

func debugf(format string, v ...interface{}) {
	if logLevel >= Debug {
		logger.Printf(format, v...)
	}
}

func debugln(v ...interface{}) {
	if logLevel >= Debug {
		logger.Println(v...)
	}
}

func trace(v ...interface{}) {
	if logLevel >= Trace {
		logger.Print(v...)
	}
}

func tracef(format string, v ...interface{}) {
	if logLevel >= Debug {
		logger.Printf(format, v...)
	}
}

func traceln(v ...interface{}) {
	if logLevel >= Debug {
		logger.Println(v...)
	}
}
