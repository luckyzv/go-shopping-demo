package lib

import (
  "fmt"
  "github.com/go-redis/redis/v8"
  "sync"
)

func CreateScript(script string) *redis.Script {
  luaScript := redis.NewScript(script)
  return luaScript
}

func SecondKillScript() *redis.Script  {
  return CreateScript(`
    local goodsStock
    local flag
    local hadBuyUserIds = tostring(KEYS[1])
    local userId = tonumber(ARGV[1])
    local goodsStockKey = tostring(KEYS[2])
    local hadBuy = redis.call("sIsMember", hadBuyUserIds, userId)

    -- 该用户已经购买商品
    if hadBuy ~= 0 then
      return 0
    end

    -- 库存键不存在
    goodsStock = redis.call("GET", goodsStockKey)
    if goodsStock == false then
      return 0
    end

    -- 库存不足
    goodsStock = tonumber(goodsStock)
    if goodsStock <= 0 then
      return 0
    end

    flag = redis.call("SADD", hadBuyUserIds, userId)
    flag = redis.call("DECR", goodsStockKey)
    return 1
  `)
}


func EvalSecondScript(client *redis.Client, userId string, wg *sync.WaitGroup) {
  secondKillScript := SecondKillScript()
  sha, err := secondKillScript.Load(client.Context(), client).Result()
  FailOnError(err, "Load lua脚本失败")

  ret := client.EvalSha(client.Context(), sha, []string{
    "hadBuyUserIds",
    "goodsStock",
  }, userId)
  result, err := ret.Result()
  if err != nil {
    FailOnError(err, "执行lua脚本失败")
  }

  fmt.Println("")
  fmt.Printf("userId: %s, result: %d", userId, result)
}
