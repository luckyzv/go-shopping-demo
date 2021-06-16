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


func Response(c *gin.Context, errorCode int, data interface{})  {
  c.JSON(http.StatusOK, SuccessResBody{
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

func ClientFailedResponse(c *gin.Context, errorCode int)  {
 c.JSON(http.StatusUnauthorized, FailedResBody{
   Code: errorCode,
   Message: constant.GetMessage(errorCode),
   Url: c.Request.URL.Path,
 })
 c.Abort()
}

func ServerFailedResponse(c *gin.Context, errCode int) {
  c.JSON(http.StatusInternalServerError, FailedResBody{
    Code: errCode,
    Message: constant.GetMessage(errCode),
    Url: c.Request.URL.Path,
  })
  c.Abort()
}
