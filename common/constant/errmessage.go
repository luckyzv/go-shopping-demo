package constant

var MessageFlags = map[int]string{
  // 基本信息
  SUCCESS: "success",
  ERROR: "fail",

  // 用户相关错误
  ERROR_AUTH_CHECK_TOKEN_FAIL: "Token鉴权失败",

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
