package logger

import (
	"path/filepath"
	"runtime"

	"go.uber.org/zap"
)

var log *zap.SugaredLogger
var logKind = "LOCAL"
var logPath string

func init() {
	newLogger()
}

// Logger return log instance
func Logger() *zap.SugaredLogger {
	if log == nil {
		newLogger()
	}
	return log
}

func newLogger() {
	if logKind == "LOCAL" {
		// get log file path
		_, file, _, _ := runtime.Caller(0)
		currPath := filepath.Dir(file)
		logPath = currPath + "/../../../log/api.log"
	}

	// set log file
	logBuilder := zap.NewDevelopmentConfig()
	logBuilder.OutputPaths = []string{
		// you can set more log file here
		logPath,
	}

	// build log instance and change it into sugar type
	logInstance, err := logBuilder.Build()
	if err != nil {
		panic("can't initialize zap logger:" + err.Error())
	}
	log = logInstance.Sugar()
}

// Close close log instance
func Close() {
	if log != nil {
		log.Sync()
	}
}

// // Info log the message with title Info into log file
// func Info(msg interface{}) {
// 	log.Info(msg)
// }

// Debug log the message with title Debug into log file
//func Debug(msg interface{}) {
//	log.Debug(msg)
//}

// // Warn log the message with title Warn into log file
// func Warn(msg interface{}) {
// 	log.Warn(msg)
// }

// Error log the message with title Error into log file
//func Error(msg interface{}) {
//	log.Error(msg)
//}

// // Fatal log the message with title Fatal into log file
// func Fatal(msg interface{}) {
// 	log.Fatal(msg)
// }

// this is only an example, won't be used in anywhere
func example() {
	// get a instance, and remeber to close it in the end
	log := Logger()
	defer Close()

	// put anything you want in func
	log.Info()
	log.Debug()
}
