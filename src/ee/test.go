package ee

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "shopping/engine"
)

func Test(ctx *gin.Context)  {
  db := engine.GetMysqlClient()
  redisClient := engine.GetRedisClient()
  mqConn := engine.GetRabbitmqConn()
  esClient := engine.GetEsClient()
  fmt.Println(db, redisClient, mqConn, esClient)
}
