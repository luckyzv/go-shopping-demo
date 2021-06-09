package middleware

import (
  "github.com/gin-gonic/gin"
  "github.com/sirupsen/logrus"
  "os"
  "path"
  "shopping/config"
  "shopping/util"
  "time"
)

var logger *logrus.Logger

func Logger() *logrus.Logger {
  viperConfig := config.GetLoggerConfig()
  logFilePath := viperConfig.FilePath
  logFileName := viperConfig.FileName

  fileName := path.Join(logFilePath, logFileName)
  if _, err := os.Stat(fileName); err != nil {
    if _, err := os.Create(fileName); err != nil {
      util.FailOnError(err, "创建日志文件失败")
    }
  }

  src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  if err != nil {
    util.FailOnError(err, "打开日志文件失败")
  }

  loggerClient := logrus.New()
  loggerClient.Out = src
  loggerClient.SetLevel(logrus.DebugLevel)
  loggerClient.SetFormatter(&logrus.TextFormatter{
   TimestampFormat: "2006-01-02 15:04:23",
  })
  return loggerClient
}

func LoggerToFile() gin.HandlerFunc {
  return LoggerFormat()
}

func LoggerFormat() func(c *gin.Context)  {
  loggerClient := Logger()
  logger = loggerClient

  return func(c *gin.Context) {
    startTime := time.Now()
    c.Next()
    endTime := time.Now()
    operationTime := endTime.Sub(startTime)
    reqMethod := c.Request.Method
    reqUri := c.Request.RequestURI
    statusCode := c.Writer.Status()
    clientIp := c.ClientIP()
    //logger.WithFields(logrus.Fields{
    //  "status_code": statusCode,
    //  "operation_time": operationTime,
    //  "client_ip": clientIp,
    //  "req_method": reqMethod,
    //  "req_uri": reqUri,
    //})
    
    // 这个是info级别的日志。gin框架本身就有使用logger
    logger.Infof("| %3d | %13v | %15s | %s | %s|",
     statusCode, operationTime, clientIp, reqMethod, reqUri)
  }
}

func GetLogger() *logrus.Logger {
  return logger
}
