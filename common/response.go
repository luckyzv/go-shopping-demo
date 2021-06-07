package common

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "shopping/common/constant"
)

type ResponseBody struct {
  Code int `json:"code"`
  Message string `json:"message"`
  Data interface{} `json:"data"`
}

type FailedResponseBody struct {
  Code int `json:"code"`
  Message string `json:"message"`
  Url string `json:"url"`
}

type NotFoundBody struct {
  Code int `json:"code"`
  Message string `json:"message"`
  Url string `json:"url"`
}


func Response(c *gin.Context, code, errorCode int, data interface{})  {
  c.JSON(code, ResponseBody{
    Code: errorCode,
    Message: constant.GetMessage(errorCode),
    Data: data,
  })
  return
}

func NotFound(c *gin.Context)  {
  c.JSON(http.StatusNotFound, FailedResponseBody{
    Code: http.StatusNotFound,
    Message: "Not Found",
    Url: c.Request.URL.Path,
  })
  c.Abort()
}

func FailedParamResponse(c *gin.Context, code, errorCode int)  {
  c.JSON(code, FailedResponseBody{
    Code: errorCode,
    Message: "参数验证失败",
    Url: c.Request.URL.Path,
  })
  c.Abort()
}
