package constant

const (
  SUCCESS = 0
  ERROR = 1

  ErrorRequiredParamFail   = 400001
  ErrorUserExisted  = 400002
  ErrorUserNonExisted     = 400003
  ErrorPasswordCheckFail = 401001
  ErrorTokenCheckFail = 403001

  ErrorCreateUserFail = 500001
  ErrorHashedPasswordFail = 500002
  ErrorTokenReleaseFail = 500003

)
