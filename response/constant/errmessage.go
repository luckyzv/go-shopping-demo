package constant

var MessageFlags = map[int]string{
  // 基本信息
  SUCCESS: "success",
  ERROR: "fail",

  Error_RequiredParamFail: "必传参数为空",

  // 用户相关错误

  Error_AuthCheckTokenFail: "Token鉴权失败",
  Error_MysqlCreateUserError: "创建失败",
  Error_UserExistedFail: "该用户已经存在",
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
