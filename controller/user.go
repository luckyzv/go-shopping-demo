package controller

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "shopping/common"
  "shopping/common/constant"
)

func UserRegister(c *gin.Context) {
  //logger := middleware.GetLogger()
  //logger.Debug("oh 内容%d ")
  common.Response(c, http.StatusOK, constant.SUCCESS, nil)
}
