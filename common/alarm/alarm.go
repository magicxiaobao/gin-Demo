package alarm

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"goDemo/common/function"
	"goDemo/config"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

type errorInfo struct {
	Time         string `json:"time"`
	Alarm        string `json:"alarm"`
	FileName     string `json:"fileName"`
	Line         int    `json:"line"`
	FunctionName string `json:"functionName"`
	Message      string `json:"message"`
}

func Info(errorMsg string) error {
	alarm("INFO", errorMsg, 2)
	return &errorString{errorMsg}
}

func Error(errorMsg string) error {
	alarm("ERROR", errorMsg, 2)
	return &errorString{errorMsg}
}

func Panic(errorMsg string) error {
	alarm("PANIC", errorMsg, 2)
	return &errorString{errorMsg}
}

func alarm(level string, msg string, skip int) {

	now := function.GetCurrentTimeStr()
	fileName, line, functionName := "?", 0, "?"
	pc, fileName, line, ok := runtime.Caller(skip)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName := filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}
	errorStruct := errorInfo{
		now,
		level,
		fileName,
		line,
		functionName,
		msg,
	}
	jsonBytes, errs := json.Marshal(errorStruct)
	if errs != nil {
		fmt.Println("json Marshal error:", errs)
	}
	jsonStr := string(jsonBytes)
	if level == "INFO" {
		filePath := config.LOG_FILE_PATH
		fileName := config.LOG_FILE_NAME
		absoluteFilePath := path.Join(filePath, fileName)
		openFile, err := os.OpenFile(absoluteFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("openFile err", err)
		}
		logger := logrus.New()
		logger.Out = openFile
		logger.SetLevel(logrus.InfoLevel)
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		logger.Info(jsonStr)
	} else if level == "ERROR" {
		filePath := config.LOG_FILE_PATH
		fileName := config.LOG_FILE_NAME
		absoluteFilePath := path.Join(filePath, fileName)
		openFile, err := os.OpenFile(absoluteFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("openFile err", err)
		}
		logger := logrus.New()
		logger.Out = openFile
		logger.SetLevel(logrus.ErrorLevel)
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		logger.Error(jsonStr)
	} else if level == "PANIC" {
		filePath := config.LOG_FILE_PATH
		fileName := config.LOG_FILE_NAME
		absoluteFilePath := path.Join(filePath, fileName)
		openFile, err := os.OpenFile(absoluteFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("openFile err", err)
		}
		logger := logrus.New()
		logger.Out = openFile
		logger.SetLevel(logrus.PanicLevel)
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		logger.Panic(msg)
	} else if level == "EMAIL" {

	} else if level == "SMS" {

	}
}
