package common

import (
  "github.com/dgrijalva/jwt-go"
  "shopping/config"
  "shopping/model"
  "time"
)

type Claims struct {
  UserId uint
  UserName string
  jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error)  {
  jwtConfig := GetJwtConfig()

  expiration := time.Now().Add(1 * 24 * time.Hour) // 一天
  claims := &Claims{
    UserId: user.ID,
    UserName: user.Name,
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: expiration.Unix(),
      IssuedAt: time.Now().Unix(),
      Issuer: jwtConfig.Issuer,
      Subject: jwtConfig.Subject,
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenStr, err := token.SignedString([]byte(jwtConfig.JwtKey))
  if err != nil {
    return "", err
  }
  return tokenStr, nil
}

func ParseToken(tokenStr string) (*Claims, error) {
  jwtConfig := GetJwtConfig()
  tokenClaims, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
    return jwtConfig.JwtKey, nil
  })
  if tokenClaims != nil {
    if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
      return claims, nil
    }
  }
  return nil, err
}

func GetJwtConfig() config.JwtConfig  {
  viperConfig := config.GetJwtConfig()
  return viperConfig
}
