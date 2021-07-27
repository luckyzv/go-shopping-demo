package order

import (
  "github.com/gin-gonic/gin"
  "shopping/common"
  "shopping/engine"
  "shopping/model"
  "shopping/response"
  "shopping/response/constant/errorcode"
  "shopping/src/order/dto"
)

type Controller struct {

}

var orderService = &Service{}

func (c *Controller) AddNewOrder(ctx *gin.Context)  {
  db := engine.GetMysqlClient()
  redisClient := engine.GetRedisClient()

  var addNewOrderDto dto.AddNewOrderDto

  if err := ctx.ShouldBindJSON(&addNewOrderDto); err != nil {
    response.ClientFailedResponse(ctx, errorcode.ErrorRequiredParamFail)
    return
  }

  if redisClient.Get(redisClient.Context(), "orderId:" + addNewOrderDto.OrderId).Val() != "1" {
    response.ClientFailedResponse(ctx, errorcode.ErrorOrderIdWrong)
    return
  }

  existed, _:= model.OrderIsExistedByOrderId(db, addNewOrderDto.OrderId)
  if existed {
    response.ClientFailedResponse(ctx, errorcode.ErrorOrderIdDuplicated)
    return
  }

  orderService.AddNewOrder(ctx, db, addNewOrderDto)
}

func (c *Controller) AddNewSecKillOrder(ctx *gin.Context)  {
  db := engine.GetMysqlClient()
  //redisClient := engine.GetRedisClient()
  var addNewOrderDto dto.AddNewOrderDto

  if err := ctx.ShouldBindJSON(&addNewOrderDto); err != nil {
    response.ClientFailedResponse(ctx, errorcode.ErrorRequiredParamFail)
    return
  }

  //if redisClient.Get(redisClient.Context(), "orderId:" + addNewOrderDto.OrderId).Val() != "1" {
  //  response.ClientFailedResponse(ctx, errorcode.ErrorOrderIdWrong)
  //  return
  //}

  isCouldBuy := checkLuaResult(ctx)
  if isCouldBuy == -1 {
    return
  }

  existed, _:= model.OrderIsExistedByOrderId(db, addNewOrderDto.OrderId)
  if existed {
    response.ClientFailedResponse(ctx, errorcode.ErrorOrderIdDuplicated)
    return
  }

  orderService.AddNewOrder(ctx, db, addNewOrderDto)
}

func (c *Controller) GetOrderId(ctx *gin.Context) {
  db := engine.GetMysqlClient()
  redisClient := engine.GetRedisClient()

  orderId := GetUniqueOrderId(ctx, db)
  redisClient.Set(redisClient.Context(), "orderId:" + orderId,"1", 0)

  response.Response(ctx, errorcode.SUCCESS, map[string]string{"orderId": orderId})
}

func checkLuaResult(ctx *gin.Context) int {
  userId, _ := ctx.Get("userId")
  luaResult := common.EvalSecondScript(userId.(uint))

  if luaResult == int64(common.LuaSuccess) {
    return 1
  }

  var code string
  if luaResult == int64(common.LuaHadBuy) {
    code = errorcode.ErrorOrderUserHadBuy
  }

  if luaResult == int64(common.LuaNumLessStock) {
    code = errorcode.ErrorOrderUserHadBuy
  }

  if luaResult == int64(common.LuaStockEmpty) {
    code = errorcode.ErrorOrderUserHadBuy
  }

  common.Logger("Common", "EvalSecondScript", code, nil, map[string]interface{}{"userId": userId})
  response.ClientFailedResponse(ctx, code)

  return -1
}
