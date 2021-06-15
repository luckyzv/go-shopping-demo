package response

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "shopping/response/constant"
)

type SuccessResBody struct {
  Code int `json:"code"`
  Message string `json:"message"`
  Data interface{} `json:"data"`
}

type FailedResBody struct {
  Code int `json:"code"`
  Message string `json:"message"`
  Url string `json:"url"`
}

//type MysqlFailedResponseBody struct {
//  Code int `json:"code"`
//  Message string `json:"message"`
//  Url string `json:"url"`
//}


func Response(c *gin.Context, code, errorCode int, data interface{})  {
  c.JSON(code, SuccessResBody{
    Code: errorCode,
    Message: constant.GetMessage(errorCode),
    Data: data,
  })
  return
}

func NotFound(c *gin.Context)  {
  c.JSON(http.StatusNotFound, FailedResBody{
    Code: http.StatusNotFound,
    Message: "Not Found",
    Url: c.Request.URL.Path,
  })
  c.Abort()
}

//func FailedParamResponse(c *gin.Context, errorCode int)  {
//  c.JSON(http.StatusUnauthorized, FailedResponseBody{
//    Code: errorCode,
//    Message: "参数验证失败",
//    Url: c.Request.URL.Path,
//  })
//  c.Abort()
//}

func FailedResponse(c *gin.Context, errCode int) {
  c.JSON(http.StatusInternalServerError, FailedResBody{
    Code: errCode,
    Message: constant.GetMessage(errCode),
    Url: c.Request.URL.Path,
  })
  c.Abort()
}
