package controller

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "shopping/common"
  "shopping/common/constant"
)

func UserRegister(c *gin.Context) {
  //group := app.Group("/api/v1/register")
  //group.POST("/doSignIg", user.)
  common.Response(c, http.StatusOK, constant.SUCCESS, nil)
}
