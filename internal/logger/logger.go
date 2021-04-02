package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	LevelError = iota + 1
	LevelWarning
	LevelInfo
	LevelDebug
)

// Default to error only
var level = LevelError

func SetLevel(l int) {
	level = l
}

// logger references the used application logger.
var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

func Debug(v ...interface{}) {
	if level >= LevelDebug {
		logger.Printf("[D] %v\n", v)
	}
}

func Info(v ...interface{}) {
	if level >= LevelInfo {
		logger.Printf("[I] %v\n", v)
	}
}

func Warn(v ...interface{}) {
	if level >= LevelWarning {
		logger.Printf("[W] %v\n", v)
	}
}

func Error(v ...interface{}) {
	if level >= LevelWarning {
		logger.Fatalf("[E] %v\n", v)
	}

	// Still print errors and exit early.
	fmt.Printf("Error: %v\n", v)
	os.Exit(1)
}
