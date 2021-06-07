package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "shopping/config"
  "shopping/router"
)

func main() {
  fmt.Println("开始了")
  r := gin.Default()

  router.Init(r)

  viperConfig := config.GetServerConfig()
  port := viperConfig.Port
  if port != "" {
    r.Run(":" + port)
  } else {
    r.Run()
  }
}
