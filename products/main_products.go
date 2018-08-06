// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Polmorphism is a piece of code or a program behave in different ways according to the different datas.
// Sample datastore demonstrates use of the cloud.google.com/go/datastore package from App Engine flexible.
package mainproducts

import (
    "time"
	  "net/http"
	  //"encoding/json"
    "github.com/gin-gonic/gin"
)

// This function's name is a must. App Engine uses it to drive the requests properly.
func init() {

// Run App at 'release' mode in production.
  gin.SetMode(gin.ReleaseMode)

// Starts a new Gin instance with no middle-ware
  router := gin.New()

  admin := router.Group("/admin")
  admin.GET("/products", ProductsList)
  admin.POST("/products", CreateProduct)
  admin.PUT("/products/:Id", UpdateProduct)
  admin.DELETE("/products/:Id", DeleteProduct)
  admin.GET("/products/orderscount", UpdateProductOrdCnt)

  v1 := router.Group("/v1")
  v1.GET("/products/:OrdSqnce", ProductsOrdSqnc)
  // Handle all requests using net/http
  http.Handle("/", router)
}

type Product struct {
  Id      int64  `json:"id" datastore:"-"`
  Title   string `json:"Title"`
  Description string `json:"Description"`
  Price float32 `json:"Price"`
  Discount float32 `json:"Discount"`
  ImageUrl  string `json:"ImageUrl"`
  OrderCount int32 `json:"OrderCount"`
  CreatedAt time.Time `json:"CreatedAt"`
  UpdatedAt time.Time `json:"UpdatedAt"`
}



