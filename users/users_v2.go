// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Polmorphism is a piece of code or a program behave in different ways according to the different datas.
// Sample datastore demonstrates use of the cloud.google.com/go/datastore package from App Engine flexible.
package users

import (
	  "net/http"
    "github.com/gin-gonic/gin"
	  "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
)


func UsersListV2(ginctx *gin.Context) {
  aectx := appengine.NewContext(ginctx.Request)
  var err error
  query := datastore.NewQuery("User").Order("FullName").Limit(10)
  var users []User
  keys, err := query.GetAll(aectx, &users)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  for inx, key := range keys {
    users[inx].Id = key.IntID()
  }  
  ginctx.JSON(http.StatusOK, gin.H{"keys": keys, "users": users})
}

func CreateUserV2(ginctx *gin.Context) {
  // [START new_context]
  aectx := appengine.NewContext(ginctx.Request)
  // [END new_context]

  // It's a User request, so handle the form submission.
  // [START new_User]
  user := User{
    EmailId:  ginctx.PostForm("EmailId"),
    Password: ginctx.PostForm("Password"),
    FullName: ginctx.PostForm("FullName"),
  }
  // [END new_User]

  if user.EmailId == "" {
    ginctx.JSON(http.StatusOK, gin.H{"Error": "EmailId must be there"})
    return
  }
  if user.Password == "" {
    ginctx.JSON(http.StatusOK, gin.H{"Error": "Password must be there"})
    return
  }
  // [START new_key]
  key := datastore.NewIncompleteKey(aectx, "User", nil)
  // [END new_key]
  // [START add_User]
  usrKey, err := datastore.Put(aectx, key, &user)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END add_User]
  user.Id = usrKey.IntID()
  ginctx.JSON(http.StatusOK, gin.H{"User": user})
}

func LoginUserV2(ginctx *gin.Context) {
  // [START new_context]
  aectx := appengine.NewContext(ginctx.Request)
  // [END new_context]

  EmailId := ginctx.PostForm("EmailId")
  Password := ginctx.PostForm("Password")

  if EmailId == "" {
    ginctx.JSON(http.StatusOK, gin.H{"Error": "EmailId must be there"})
    return
  }
  if Password == "" {
    ginctx.JSON(http.StatusOK, gin.H{"Error": "Password must be there"})
    return
  }

  var err error
  query := datastore.NewQuery("User").Filter("EmailId =", EmailId).Filter("Password =", Password).Limit(1)
  var users []User
  _, err = query.GetAll(aectx, &users)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  if len(users) == 0 {
    ginctx.JSON(http.StatusOK, gin.H{"Error": "Invalid EmailId or Password"})
    return
  }

  user := users[0]
  accessToken := GenerateAccessToken(user.EmailId, true)

  ginctx.JSON(http.StatusOK, gin.H{"accesstoken": accessToken})
}

func PingUserV2(ginctx *gin.Context) {
  ginctx.String(http.StatusOK, "V2 pong")
}