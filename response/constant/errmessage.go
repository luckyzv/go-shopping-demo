package constant

import "shopping/response/constant/errorcode"

var MessageFlags = map[string]string{
  // 基本信息
  errorcode.SUCCESS: "success",
  errorcode.ERROR:   "fail",
  errorcode.ErrorRequiredParamFail:  "必传参数为空",
  errorcode.ErrorRequiredHeaderFail: "必传请求头为空",
  errorcode.ErrorTokenCheckFail:         "Token认证失败",
  errorcode.ErrorTokenTimeOut:           "Token过期",

  // 用户相关错误
  errorcode.ErrorUserHashedPasswordFail: "密码加密失败",
  errorcode.ErrorUserPasswordCheckFail:  "密码输入错误",

  errorcode.ErrorUserTokenReleaseFail:   "Token发放失败",
  errorcode.ErrorUserCreateUserFail:     "创建用户失败",
  errorcode.ErrorUserNonExisted:         "该用户不存在",
  errorcode.ErrorUserExisted:            "该用户已经存在",
  errorcode.ErrorUserFindFail:           "查找用户失败",
  // 产品相关错误
  errorcode.ErrorProductSkuIdDuplicated:   "SkuId已经存在",
  errorcode.ErrorProductCreateProductFail: "创建产品失败",
  errorcode.ErrorProductPriceNotAllowed:   "价格设置不合理",
  // 订单相关错误
  errorcode.ErrorOrderIdDuplicated: "请勿重复提交订单",
  errorcode.ErrorOrderTotalPriceWrong: "价格提交错误",
  errorcode.ErrorOrderUpdate: "订单更新失败",
}

func GetMessage(code string) string {
  msg, ok := MessageFlags[code]
  if ok {
    return msg
  }
  return MessageFlags[errorcode.ERROR]
}
