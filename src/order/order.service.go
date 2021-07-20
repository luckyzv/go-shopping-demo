package order

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "shopping/model"
  "shopping/response"
  "shopping/response/constant/errorcode"
  "shopping/src/order/dto"
  "shopping/util"
)

type Service struct {

}

func (orderService Service) AddNewOrder(ctx *gin.Context, db *gorm.DB, addNewOrderDto dto.AddNewOrderDto)  {
  var productIds []uint
  var totalPrice float64
  var orderProducts []model.OrderProduct
  productInfo := make(map[uint]uint)

  for _, orderProduct := range addNewOrderDto.OrderProducts {
    productIds = append(productIds, orderProduct.ProductId)
    productInfo[orderProduct.ProductId] = orderProduct.Number
  }

  // 根据数据库，计算出实际的总价
  products, _ := model.ProductGetAll(db, productIds)
  for _, product := range products {
    totalPrice = totalPrice + (product.Price * float64(productInfo[product.ID]))
    orderProducts = append(orderProducts, model.OrderProduct{
      Price: product.Price,
      Num: productInfo[product.ID],
      ProductID: product.ID,
    })
  }
  if totalPrice != addNewOrderDto.TotalPrice {
    response.ClientFailedResponse(ctx, errorcode.ErrorOrderTotalPriceWrong)
    return
  }

  userId, _ := ctx.Get("userId")
  user, _ := model.UserGetOneById(db, userId.(uint))

  order := model.Order{
    OrderId: addNewOrderDto.OrderId,
    Price: totalPrice,
    User: *user,
    Products: orderProducts,
    OrderHistories: []model.OrderHistory{{}},
    OrderPayments: []model.OrderPayment{{TotalPrice: totalPrice, Discount: 0, RealPaid: totalPrice }},
  }
  err := model.OrderAddNew(db, order)
  if err != nil {
    response.ServerFailedResponse(ctx, errorcode.ErrorOrderAddNew)
    return
  }

  response.Response(ctx, errorcode.SUCCESS, order)
}

func GetUniqueOrderId(ctx *gin.Context, db *gorm.DB) string  {
  orderId := util.GetUniqueOrderId()

  existed, _ := model.OrderIsExistedByOrderId(db, orderId)
  if existed {
    GetUniqueOrderId(ctx, db)
  }
  return orderId
}

