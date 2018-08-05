// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Polmorphism is a piece of code or a program behave in different ways according to the different datas.
// Sample datastore demonstrates use of the cloud.google.com/go/datastore package from App Engine flexible.
package mainproducts

import (
	  "net/http"
	  //"encoding/json"
    "github.com/gin-gonic/gin"
	  "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
)


func ProductsOrdSqnc(ginctx *gin.Context) {
  aectx := appengine.NewContext(ginctx.Request)
  var err error
  query := datastore.NewQuery("Product").Order("Title").Limit(10)
  var products []Product
  keys, err := query.GetAll(aectx, &products)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  for inx, key := range keys {
    products[inx].Id = key.IntID()
  }
  ginctx.JSON(http.StatusOK, gin.H{"keys": keys, "products": products})
}