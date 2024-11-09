package util

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var LogrusObj *logrus.Logger

func InitLog() {
	if LogrusObj != nil {
		src, _ := setOutputFile()
		LogrusObj.Out = src
		return
	}
	logger := logrus.New()
	src, _ := setOutputFile()
	logger.Out = src

	//	设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	LogrusObj = logger
}

// 设置日志输出的文件路径
func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	src, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	return src, nil
}
