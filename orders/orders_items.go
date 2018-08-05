// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Polmorphism is a piece of code or a program behave in different ways according to the different datas.
// Sample datastore demonstrates use of the cloud.google.com/go/datastore package from App Engine flexible.
package ordersitems

import (
    "time"
	  "net/http"
	  //"encoding/json"
    "github.com/gin-gonic/gin"
)

// This function's UserId is a must. App Engine uses it to drive the requests properly.
func init() {

// Run App at 'release' mode in production.
  gin.SetMode(gin.ReleaseMode)

// Starts a new Gin instance with no middle-ware
  router := gin.New()
  v1 := router.Group("/v1")
  v1.GET("/orders/:UsrId/list", OrdersListV1)
  v1.POST("/orders/:UsrId/place", CreateOrderV1)
  v1.PUT("/orders/:OrdId/tranxn", OrderTransactionV1)
  v1.PUT("/orders/:OrdId/ordstatus", ChangeOrderStatusV1)
  v1.DELETE("/orders/:OrdId/remove", DeleteOrderV1)

  v1.POST("/items/:UsrId", CreateItemV1)
  v1.GET("/items/:UsrId/usrlist", UserItemsV1)
  v1.PUT("/items/:UsrId", UpdateItemsV1)
  v1.DELETE("/items/:ItmId", DeleteItemV1)
  v1.GET("/orditems/:OrdId/ordlist", OrderItemsV1)

  // Handle all requests using net/http
  http.Handle("/", router)
}

type Order struct {
  Id      int64  `json:"id" datastore:"-"`
  UserId   int64 `json:"UserId"`
  Amount float32 `json:"Amount"`
  Discount float32 `json:"Discount"`
  TtlAmount float32 `json:"TtlAmount"`
  OrdStatus string `json:"OrdStatus"`
  CreatedAt time.Time `json:"CreatedAt"`
  TraxnId string `json:"TraxnId"`
  TraxnStatus string `json:"TraxnStatus"`
  TrxnDate time.Time `json:"TrxnDate"`
  UpdatedAt time.Time `json:"UpdatedAt"`
}

type Item struct {
  Id      int64  `json:"id" datastore:"-"`
  UserId   int64 `json:"UserId"`
  ProdId  int64 `json:"ProdId"`
  OrderId  int64 `json:"UserId"`
  Amount float32 `json:"Amount"`
  Discount float32 `json:"Discount"`
  TtlAmount float32 `json:"TtlAmount"`
  CreatedAt time.Time `json:"CreatedAt"`
}