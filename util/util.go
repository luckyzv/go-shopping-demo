package util

import (
  "crypto/md5"
  "fmt"
  "io"
  "math/rand"
  "strconv"
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

func GetUniqueOrderId() string  {
  month := time.Now().Format("01")
  day := time.Now().Format("02")
  timeStamp := time.Now().Unix()
  randomString := GetRandomString(3)

  return month + day + strconv.FormatInt(timeStamp, 10) + randomString
}
