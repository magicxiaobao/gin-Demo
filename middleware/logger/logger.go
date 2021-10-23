package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goDemo/config"
	"os"
	"path"
	"time"
)

func LoggerToFile() gin.HandlerFunc {

	logFilePath := config.LOG_FILE_PATH
	logFileName := config.LOG_FILE_NAME
	absoluteFilePath := path.Join(logFilePath, logFileName)
	fmt.Printf("absoluteFilePath: %v", absoluteFilePath)
	openFile, err := os.OpenFile(absoluteFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("openFile err", err)
	}
	logger := logrus.New()
	logger.Out = openFile
	logger.SetLevel(logrus.DebugLevel)
	//logger.SetFormatter(&logrus.TextFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//})
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next() // handle request
		endTime := time.Now()
		costTime := endTime.Sub(startTime)
		method := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		httpStatus := ctx.Writer.Status()
		ip := ctx.ClientIP()
		//logger.Infof("| %3d | %13v | %15s | %s | %s |",
		//	httpStatus,
		//	costTime,
		//	ip,
		//	method,
		//	reqUri,
		//	)
		logger.WithFields(logrus.Fields{
			"status_code":  httpStatus,
			"latency_time": costTime,
			"client_ip":    ip,
			"req_method":   method,
			"req_uri":      reqUri,
		}).Info()
	}
}
