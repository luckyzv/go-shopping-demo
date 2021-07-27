package common

import (
  "fmt"
  "github.com/go-redis/redis/v8"
  "shopping/engine"
)

const (
  LuaNumLessStock = iota - 3

  LuaKeyNonExist

  LuaStockEmpty

  LuaHadBuy

  LuaSuccess
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
    local goodsStockKey = tostring(KEYS[2])

    local userId = tonumber(ARGV[1])
    local hadBuy = redis.call("sIsMember", hadBuyUserIds, userId)

    -- 库存键不存在
    goodsStock = redis.call("GET", goodsStockKey)
    if goodsStock == false then
      return -2
    end

    -- 该用户已经购买商品
    -- if hadBuy ~= 0 then
     -- return 0
    -- end

    -- 售空，库存<=0
    goodsStock = tonumber(goodsStock)
    if goodsStock <= 0 then
      return -1
    end

    -- 库存数小于用户购买数
    -- if goodsStock <= 0 then
     -- return -3
    -- end

    flag = redis.call("SADD", hadBuyUserIds, userId)
    flag = redis.call("DECR", goodsStockKey)
    return 1
  `)
}


func EvalSecondScript(userId uint) interface{} {
  client := engine.GetRedisClient()
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

  fmt.Printf("userId: %d, result: %d\n", userId, result)

  return result
}
