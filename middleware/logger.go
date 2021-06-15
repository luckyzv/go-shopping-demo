package middleware

import (
  "bytes"
  "encoding/json"
  "github.com/gin-gonic/gin"
  rotatelogs "github.com/lestrrat-go/file-rotatelogs"
  "github.com/rifflock/lfshook"
  "github.com/sirupsen/logrus"
  "os"
  "path"
  "shopping/response"
  "shopping/config"
  "shopping/lib"
  "shopping/util"
  "strconv"
  "time"
)

// TODO:疑问代码片段
type bodyLogWriter struct {
  gin.ResponseWriter
  body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error)  {
  w.body.Write(b)
  return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error)  {
  w.body.WriteString(s)
  return w.ResponseWriter.WriteString(s)
}

var logger *logrus.Logger

func Logger() *logrus.Logger {
  viperConfig := config.GetLoggerConfig()
  logFilePath := viperConfig.FilePath
  logFileName := viperConfig.FileName

  // 日志文件
  fileName := path.Join(logFilePath, logFileName)
  initFile(fileName)

  // 实例化logger
  loggerClient := logrus.New()

  // 设置熟悉、新增Hook
  loggerClient.SetLevel(logrus.DebugLevel)
  lfHook, writer := rotateFile(fileName)
  loggerClient.Out = writer
  loggerClient.SetFormatter(&logrus.JSONFormatter{
    TimestampFormat: "2006-01-02 15:04:23",
  })
  loggerClient.AddHook(lfHook)

  return loggerClient
}

func initFile(fileName string) *os.File {
  // 创建日志文件
  //if _, err := os.Stat(fileName); err != nil {
  //  if _, err := os.Create(fileName); err != nil {
  //    util.FailOnError(err, "创建日志文件失败")
  //  }
  //}

  src, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
  if err != nil {
    lib.FailOnError(err, "打开日志文件失败")
  }
  return src
}

func rotateFile(fileName string) (*lfshook.LfsHook, *rotatelogs.RotateLogs) {
  // 分割文件
  logWriter, err := rotatelogs.New(
    fileName + ".%Y%m%d.log", // 分割后的文件名称
    rotatelogs.WithLinkName(fileName), // 软链
    rotatelogs.WithMaxAge(7 * 24 * time.Hour), // 最大保存时间（7天）
    rotatelogs.WithRotationTime(24 * time.Hour), // 按照天来分割
  )
  lib.FailOnError(err, "分割日志失败")
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

func LoggerToFile() gin.HandlerFunc {
  return LoggerFormat()
}

func LoggerFormat() func(c *gin.Context)  {
  loggerClient := Logger()
  logger = loggerClient

  return func(c *gin.Context) {
    var responseCode string
    var responseMsg string
    var responseData interface{}

    // TODO:疑问代码片段
    bodyLogWriter := &bodyLogWriter{
      body: bytes.NewBufferString(""),
      ResponseWriter: c.Writer,
    }
    c.Writer = bodyLogWriter

    // 开始时间
    startTime := time.Now()
    // 处理请求
    c.Next()
    // 请求结束
    responseBody := bodyLogWriter.body.String()
    if responseBody != "" {
      res := response.Body{}
      err := json.Unmarshal([]byte(responseBody), &res)
      if err == nil {
        responseCode = strconv.Itoa(res.Code)
        responseMsg = res.Message
        responseData = res.Data
      }
    }

    // 结束时间
    endTime := time.Now()
    if c.Request.Method == "POST" {
      c.Request.ParseForm()
    }
    clientIp := c.ClientIP()
    loggerClient.WithFields(logrus.Fields{
      "RequestID": util.GetMd5String(strconv.Itoa(int(time.Now().Unix()))),
      "req_method": c.Request.Method,
      "req_uri": c.Request.RequestURI,
      "req_post_data": c.Request.PostForm.Encode(),
      //"req_user_agent": c.Request.UserAgent(),
      "req_client_ip": clientIp,

      "res_status_code": c.Writer.Status(),
      "res_code": responseCode,
      "res_message": responseMsg,
      "res_data": responseData,

      "operation_time": strconv.Itoa(int(endTime.Sub(startTime).Milliseconds())) + "ms",
    }).Info()
  }
}

func GetLogrusLogger() *logrus.Logger {
  return logger
}
