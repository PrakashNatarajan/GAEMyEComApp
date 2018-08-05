// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Polmorphism is a piece of code or a program behave in different ways according to the different datas.
// Sample datastore demonstrates use of the cloud.google.com/go/datastore package from App Engine flexible.
package gaemyecomapp

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
)

// This function's name is a must. App Engine uses it to drive the requests properly.
func init() {

// Run App at 'release' mode in production.
  gin.SetMode(gin.ReleaseMode)

// Starts a new Gin instance with no middle-ware
  route := gin.New()
  
  // Define your handlers
  route.GET("/", func(ctx *gin.Context) {
    ctx.String(http.StatusOK, "Hello World!")
  })
  route.GET("/ping", func(ctx *gin.Context) {
    ctx.String(http.StatusOK, "pong")
  })

  route.GET("/kinds", KindsList)

  // Handle all requests using net/http
  http.Handle("/", route)
}


func KindsList(ginctx *gin.Context) {
  aectx := appengine.NewContext(ginctx.Request)  
  kinds, err := datastore.Kinds(aectx)
  if err != nil {
    //ginctx.String(http.StatusOK, "Error")
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    //http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }
  fmt.Println(kinds)
  ginctx.JSON(http.StatusOK, gin.H{"kinds": kinds})
  //ginctx.String(http.StatusOK, "kinds")
}