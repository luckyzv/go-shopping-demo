package order

import (
  "github.com/gin-gonic/gin"
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
  var addNewOrderDto dto.AddNewOrderDto

  if err := ctx.ShouldBindJSON(&addNewOrderDto); err != nil {
    response.ClientFailedResponse(ctx, errorcode.ErrorRequiredParamFail)
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

  orderId := GetUniqueOrderId(ctx, db)
  response.Response(ctx, errorcode.SUCCESS, map[string]string{"orderId": orderId})
}
