package response

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "shopping/response/constant"
  "strconv"
)

type SuccessResBody struct {
  Code string `json:"code"`
  Message string `json:"message"`
  Data interface{} `json:"data"`
}

type FailedResBody struct {
  Code string `json:"code"`
  Message string `json:"message"`
  Url string `json:"url"`
}

//type MysqlFailedResponseBody struct {
//  Code int `json:"code"`
//  Message string `json:"message"`
//  Url string `json:"url"`
//}


func Response(c *gin.Context, errorCode string, data interface{})  {
  c.JSON(http.StatusOK, SuccessResBody{
    Code: errorCode,
    Message: constant.GetMessage(errorCode),
    Data: data,
  })
  return
}

func NotFound(c *gin.Context)  {
  c.JSON(http.StatusNotFound, FailedResBody{
    Code: strconv.Itoa(http.StatusNotFound),
    Message: "Not Found",
    Url:     c.Request.URL.Path,
  })
  c.Abort()
}

func ClientFailedResponse(c *gin.Context, errorCode string)  {
 c.JSON(http.StatusUnauthorized, FailedResBody{
   Code: errorCode,
   Message: constant.GetMessage(errorCode),
   Url: c.Request.URL.Path,
 })
 c.Abort()
}

func ServerFailedResponse(c *gin.Context, errCode string) {
  c.JSON(http.StatusInternalServerError, FailedResBody{
    Code: errCode,
    Message: constant.GetMessage(errCode),
    Url:     c.Request.URL.Path,
  })
  c.Abort()
}
