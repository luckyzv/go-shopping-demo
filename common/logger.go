package common

import (
  "fmt"
  rotatelogs "github.com/lestrrat-go/file-rotatelogs"
  "github.com/rifflock/lfshook"
  "github.com/sirupsen/logrus"
  "os"
  "path"
  "shopping/config"
  "time"
)

// ApiLogger 记录日常请求日志
func ApiLogger() *logrus.Logger {
  viperConfig := config.GetLoggerConfig()
  logger := initLogger(viperConfig.ApiFilePath)
  return logger
}

// 记录系统日志
func appLogger() *logrus.Logger {
  viperConfig := config.GetLoggerConfig()
  logger := initLogger(viperConfig.AppFilePath)
  return logger
}

func initLogger(logFilePath string) *logrus.Logger {
  src := initFile(logFilePath)

  // 实例化logger
  logger := logrus.New()

  // 新增Hook
  //lfHook, src := rotateFileAndNewHook(fileName)
  //loggerClient.AddHook(lfHook)

  logger.Out = src
  logger.SetLevel(logrus.DebugLevel)

  logger.SetFormatter(&logrus.JSONFormatter{
   TimestampFormat: "2006-01-02 15:04:23",
  })
  return logger
}

func initFile(logFilePath string) *os.File {
  now := time.Now()
  if err := os.MkdirAll(logFilePath, 0777); err != nil {
    FailOnError(err, "创建日志目录失败")
  }
  // 创建日志文件
  logFileName := now.Format("2006-01-02") + ".log"
  fileName := path.Join(logFilePath, logFileName)

  if _, err := os.Stat(fileName); err != nil {
   if _, err := os.Create(fileName); err != nil {
     FailOnError(err, "创建日志文件失败")
   }
  }

  src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  if err != nil {
    FailOnError(err, "打开日志文件失败")
  }
  return src
}

func rotateFileAndNewHook(fileName string) (*lfshook.LfsHook, *rotatelogs.RotateLogs) {
  // 分割文件
  logWriter, err := rotatelogs.New(
    fileName + ".%Y%m%d.log", // 分割后的文件名称
    rotatelogs.WithLinkName(fileName), // 软链
    rotatelogs.WithMaxAge(7 * 24 * time.Hour), // 最大保存时间（7天）
    rotatelogs.WithRotationTime(24 * time.Hour), // 按照天来分割
  )
  FailOnError(err, "分割日志失败")
  writeMap := lfshook.WriterMap{
    logrus.InfoLevel: logWriter,
    logrus.FatalLevel: logWriter,
    logrus.DebugLevel: logWriter,
    logrus.WarnLevel: logWriter,
    logrus.ErrorLevel: logWriter,
    logrus.PanicLevel: logWriter,
  }
  lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
    TimestampFormat: "2006-01-02 15:04:23",
  })
  return lfHook, logWriter
}

// Logger 记录系统内部错误接口 结合请求日志排查bug
// @packageName 包名 @funcName 函数名 @errCode 内部错误码 @err 错误内容
func Logger(packageName string, funcName string, errCode string, err error, funcParam interface{})  {
  appLogger().WithFields(logrus.Fields{
    "packageName": packageName,
    "funcName": funcName,
    "innerErrCode": errCode,
    "funcParam": funcParam,
  }).Errorf("系统内部错误：【%v】", err)
}

func FailOnError(err error, msg string)  {
  if err != nil {
    panic(fmt.Errorf("%s, err: %s \n", msg, err))
  }
}
