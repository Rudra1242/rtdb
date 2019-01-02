package loggerutils

import (

	//"github.com/coreos/go-log/log"

	"go.uber.org/zap"
)

// Define your custom logger type.
var loggerObj *zap.SugaredLogger

//GetLogger function
// func GetLogger() *log.Logger {
// 	if loggerObj == nil {
// 		loggerObj = log.New("orderhold", false,
// 			log.WriterSink(os.Stderr,
// 				"%v [%d]:%s %s\n",
// 				[]string{"full_time", "pid", "priority", "message"}))
// 	}
// 	return loggerObj
// }

// GetLogger provides sugaredLogger instance
func GetLogger() *zap.SugaredLogger {
	if loggerObj == nil {
		logger, _ := zap.NewDevelopment()
		defer logger.Sync() // flushes buffer, if any
		loggerObj = logger.Sugar()
		loggerObj.Infof("Initialized Logger: %s", "Suggared Logger")
	}

	return loggerObj
}
