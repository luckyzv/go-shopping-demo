package constant

var MessageFlags = map[int]string{
  // 基本信息
  SUCCESS: "success",
  ERROR: "fail",

  ErrorRequiredParamFail: "必传参数为空",

  // 用户相关错误

  ErrorAuthCheckTokenFail: "Token鉴权失败",

  // 产品相关错误


  // 订单相关错误
}

func GetMessage(code int) string {
  msg, ok := MessageFlags[code]
  if ok {
    return msg
  }
  return MessageFlags[ERROR]
}
