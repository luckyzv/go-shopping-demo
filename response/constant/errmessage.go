package constant

var MessageFlags = map[int]string{
  // 基本信息
  SUCCESS: "success",
  ERROR: "fail",

  ErrorRequiredParamFail: "必传参数为空",
  ErrorRequiredHeaderFail: "必传请求头为空",

  // 用户相关错误
  ErrorHashedPasswordFail:     "密码加密失败",
  ErrorPasswordCheckFail: "密码输入错误",
  ErrorTokenCheckFail:   "Token认证失败",
  ErrorTokenTimeOut: "Token过期",
  ErrorTokenReleaseFail: "Token发放失败",
  ErrorCreateUserFail: "创建用户失败",
  ErrorUserNonExisted:      "该用户不存在",
  ErrorUserExisted: "该用户已经存在",
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
