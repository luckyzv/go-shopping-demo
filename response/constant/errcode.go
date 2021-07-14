package constant

const (
  SUCCESS = 0
  ERROR = 1

  ErrorRequiredParamFail   = 40001 // 必传参数为空
  ErrorRequiredHeaderFail   = 40004
  ErrorTokenCheckFail = 40301
  ErrorTokenTimeOut   = 40005

  ErrorUserExisted  = 40002
  ErrorUserNonExisted     = 40003
  ErrorUserPasswordCheckFail = 40101

  ErrorUserCreateUserFail = 50001
  ErrorUserHashedPasswordFail = 50002
  ErrorUserTokenReleaseFail = 50003
  ErrorUserFindFail = 50004

  ErrorProductSkuIdDuplicated = 40006
  ErrorProductCreateProductFail = 50005
)
