package constant

var MessageFlags = map[int]string{
  // 基本信息
  SUCCESS: "success",
  ERROR: "fail",

  ErrorRequiredParamFail: "必传参数为空",
  ErrorRequiredHeaderFail: "必传请求头为空",

  // 用户相关错误
  ErrorUserHashedPasswordFail:     "密码加密失败",
  ErrorUserPasswordCheckFail: "密码输入错误",
  ErrorTokenCheckFail:   "Token认证失败",
  ErrorTokenTimeOut: "Token过期",
  ErrorUserTokenReleaseFail: "Token发放失败",
  ErrorUserCreateUserFail: "创建用户失败",
  ErrorUserNonExisted:      "该用户不存在",
  ErrorUserExisted: "该用户已经存在",
  ErrorUserFindFail: "查找用户失败",
  // 产品相关错误
  ErrorProductSkuIdDuplicated: "SkuId已经存在",
  ErrorProductCreateProductFail: "创建产品失败",
  ErrorProductPriceNotAllowed: "价格设置不合理",
  // 订单相关错误
}

func GetMessage(code int) string {
  msg, ok := MessageFlags[code]
  if ok {
    return msg
  }
  return MessageFlags[ERROR]
}
