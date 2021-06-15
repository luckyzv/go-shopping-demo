package util

import (
  "crypto/md5"
  "fmt"
  "io"
  "math/rand"
  "time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetRandomString(length int) string {
  rand.Seed(time.Now().Unix())
  b := make([]byte, length)
  for i := range b {
    b[i] = letterBytes[rand.Intn(len(letterBytes))]
  }
  return string(b)
}

func GetMd5String(data string) string {
  w := md5.New()
  io.WriteString(w, data)
  md5Str := fmt.Sprintf("%x", w.Sum(nil))
  return md5Str
}
