package logger

import (
	"log"
	"os"
)

// Log Levels
const (
	InfoLevel = iota
	WarningLevel
	ErrorLevel
)

type Logger struct {
	Level int
	infoLogger *log.Logger
	warnLogger *log.Logger
	errorLogger *log.Logger
}

var logger *Logger

// init function
func init()  {
	logger = &Logger{
		Level: InfoLevel, // just the default value
		infoLogger: log.New(os.Stdout, "INFO: ", log.LstdFlags),
		warnLogger: log.New(os.Stdout, "WARN: ", log.LstdFlags),
		errorLogger: log.New(os.Stdout, "ERROR: ", log.LstdFlags | log.Lshortfile),																												
	}
}

// Set Log Levels
func SetLevel(level int)  {
	logger.Level = level
}	


// Methods to logs
func Info(message string) {
	if logger.Level <= InfoLevel {
		logger.infoLogger.Println(message)
	}
}

func Warn(message string) {
	if logger.Level <= WarningLevel {
		logger.warnLogger.Println(message)
	}
}

func Error(message string) {
	if logger.Level <= ErrorLevel {
		logger.errorLogger.Println(message)
	}
}