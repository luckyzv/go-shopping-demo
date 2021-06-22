package middleware

import (
  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
  "shopping/common"
  "shopping/response"
  "shopping/response/constant"
)

func Auth() gin.HandlerFunc {
  return func(ctx *gin.Context) {
    var code int
    prefix := "Bearer "
    token := ctx.GetHeader("Authorization")
    if token == "" || len(token) < len(prefix) || token[:len(prefix)] != prefix {
      response.ClientFailedResponse(ctx, constant.ErrorRequiredHeaderFail)
      return
    }

    claims, err := common.ParseToken(token[len(prefix):])
    if err != nil {
      switch err.(*jwt.ValidationError).Errors {
      case jwt.ValidationErrorExpired:
        code = constant.ErrorTokenTimeOut
      default:
        code = constant.ErrorTokenCheckFail
      }
      response.ClientFailedResponse(ctx, code)
      return
    }

    ctx.Set("userId", claims.UserId)
    ctx.Set("userName", claims.UserName)

    ctx.Next()
  }
}
