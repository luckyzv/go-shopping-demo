package main

import (
  "github.com/gin-gonic/gin"
  "shopping/config"
  _ "shopping/engine"
  "shopping/middleware"
  "shopping/router"
)

func main() {
  r := gin.Default()
  r.Use(middleware.LoggerToFile())

  router.Init(r)

  //var wg sync.WaitGroup
  //for _, v := range []string{"1","2","1","2","2","3","7"} {
  //  wg.Add(1)
  //  go util.EvalScript(engine.GetRedisClient(), v, &wg)
  //}

  viperConfig := config.GetServerConfig()
  port := viperConfig.Port
  if port != "" {
    r.Run(":" + port)
  } else {
    r.Run()
  }
}
