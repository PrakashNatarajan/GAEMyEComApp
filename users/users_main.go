// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Polmorphism is a piece of code or a program behave in different ways according to the different datas.
// Sample datastore demonstrates use of the cloud.google.com/go/datastore package from App Engine flexible.
package users

import (
    "time"
    "strconv"
    "crypto/hmac"
    "encoding/base64"
    "crypto/sha1"
	  "net/http"
    "github.com/gin-gonic/gin"
)

// This function's name is a must. App Engine uses it to drive the requests properly.
func init() {

// Run App at 'release' mode in production.
  gin.SetMode(gin.ReleaseMode)

// Starts a new Gin instance with no middle-ware
  router := gin.New()

  v1 := router.Group("/v1")
  v1.GET("/users/list", UsersListV1)
  v1.POST("/users/register", CreateUserV1)
  v1.POST("/users/login", LoginUserV1)
  v1.GET("/users/ping", PingUserV1)

  v2 := router.Group("/v2")
  v2.GET("/users/list", UsersListV2)
  v2.POST("/users/register", CreateUserV2)
  v2.POST("/users/login", LoginUserV2)
  v2.GET("/users/ping", PingUserV2)
  
  // Handle all requests using net/http
  http.Handle("/", router)
}

type User struct {
  Id      int64  `json:"id" datastore:"-"`
  EmailId   string `json:"EmailId"`
  Password string `json:"Password"`
  FullName string `json:"FullName"`
  AccessToken string `json:"AccessToken"`
}


func GenerateAccessToken(full_url string, escapesign bool) string {
  secretkey := "2FMOHJPXXLCCDCPUWXQRPJEI&"
  hashfun := hmac.New(sha1.New, []byte(secretkey))
  timestampUrl := CurrentTimeStamp() + full_url
  hashfun.Write([]byte(timestampUrl))
  rawSignature := hashfun.Sum(nil)
  base64signature := base64.StdEncoding.EncodeToString(rawSignature)
  if escapesign {
    base64signature = escape(base64signature)
  }
  return base64signature
}

func CurrentTimeStamp() string {
  timestamp := strconv.FormatInt(time.Now().Unix(), 10)
  return timestamp
}

func escape(s string) string {
  t := make([]byte, 0, 3*len(s))
  for i := 0; i < len(s); i++ {
    c := s[i]
    if isEscapable(c) {
      t = append(t, '%')
      t = append(t, "0123456789ABCDEF"[c>>4])
      t = append(t, "0123456789ABCDEF"[c&15])
    } else {
      t = append(t, s[i])
    }
  }
  return string(t)
}

func isEscapable(b byte) bool {
  return !('A' <= b && b <= 'Z' || 'a' <= b && b <= 'z' || '0' <= b && b <= '9' || b == '-' || b == '.' || b == '_' || b == '~')
}
