package errorcode

const (
  ErrorUserExisted  = "u40002"
  ErrorUserNonExisted     = "u40003"
  ErrorUserPasswordCheckFail = "u40101"

  ErrorUserCreateUserFail = "u50001"
  ErrorUserHashedPasswordFail = "u50002"
  ErrorUserTokenReleaseFail = "u50003"
  ErrorUserFindFail = "u50004"
)
