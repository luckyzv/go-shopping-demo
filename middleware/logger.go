package middleware

import (
  "bytes"
  "encoding/json"
  "github.com/gin-gonic/gin"
  "github.com/sirupsen/logrus"
  "io/ioutil"
  "shopping/common"
  "shopping/response"
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

func Logger() gin.HandlerFunc {
  return loggerFormat()
}

func loggerFormat() func(c *gin.Context)  {
  logger := common.ApiLogger()

  return func(c *gin.Context) {
    var responseCode string
    var responseMsg string
    var responseData interface{}

    // 写入响应到自己的结构体
    bodyLogWriter := &bodyLogWriter{
      body: bytes.NewBufferString(""),
      ResponseWriter: c.Writer,
    }
    c.Writer = bodyLogWriter

    // 开始时间
    startTime := time.Now()
    // 解析application/json格式下的请求体。因为gin只允许读取一次请求body，所以特殊处理
    if c.Request.Method == "POST" {
      reqBody := make(map[string]interface{})
      data, _ := c.GetRawData()
      json.Unmarshal(data, &reqBody)
      c.Set("reqBodyString", reqBody)
      c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
    }
    // 处理请求
    c.Next()

    // 请求结束 会将响应体写入c.Writer(通过write方法)
    responseBody := bodyLogWriter.body.String()
    if responseBody != "" {
      res := response.SuccessResBody{}
      err := json.Unmarshal([]byte(responseBody), &res)
      if err == nil {
        responseCode = res.Code
        responseMsg = res.Message
        responseData = res.Data
      }
    }

    // 结束时间
    endTime := time.Now()
    clientIp := c.ClientIP()
    requestBody, _ := c.Get("reqBodyString")

    logger.WithFields(logrus.Fields{
      "RequestID": util.GetMd5String(strconv.Itoa(int(time.Now().Unix()))),
      "req_method": c.Request.Method,
      "req_uri": c.Request.RequestURI,
      "req_post_data": requestBody,
      //"req_user_agent": c.Request.UserAgent(),
      "req_client_ip": clientIp,

      "res_status_code": c.Writer.Status(),
      "res_code": responseCode,
      "res_message": responseMsg,
      "res_data": responseData,

      "latency": strconv.Itoa(int(endTime.Sub(startTime).Milliseconds())) + "ms",
    }).Info("请求已响应")
  }
}
