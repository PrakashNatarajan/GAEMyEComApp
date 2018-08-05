// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Polmorphism is a piece of code or a program behave in different ways according to the different datas.
// Sample datastore demonstrates use of the cloud.google.com/go/datastore package from App Engine flexible.
package ordersitems

import (
    "fmt"
    //"io"
	  "time"
    //"strings"
    "strconv"
	  "net/http"
	  //"encoding/json"
    "github.com/gin-gonic/gin"
	  "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
)


func OrdersListV1(ginctx *gin.Context) {
  aectx := appengine.NewContext(ginctx.Request)
  var err error
  UsrId := ginctx.Param("UsrId")
  UserId, err := strconv.ParseInt(UsrId, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  query := datastore.NewQuery("Order").Filter("UserId =", UserId).Limit(10)
  var orders []Order
  keys, err := query.GetAll(aectx, &orders)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  for inx, key := range keys {
    orders[inx].Id = key.IntID()
  }
 
  ginctx.JSON(http.StatusOK, gin.H{"keys": keys, "orders": orders})
}

func CreateOrderV1(ginctx *gin.Context) {
  // [START new_context]
  aectx := appengine.NewContext(ginctx.Request)
  // [END new_context]

  // [START get_Id]
  //path := req.URL.Path
  //Id := strings.TrimPrefix(path, "/updateorder/")
  UsrId := ginctx.Param("UsrId")
  UserId, err := strconv.ParseInt(UsrId, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END get_Id]

  Amt := ginctx.PostForm("Amount")
  Amnt, err := strconv.ParseFloat(Amt, 32)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  Amount := float32(Amnt)

  Disct := ginctx.PostForm("Discount")
  Discnt, err := strconv.ParseFloat(Disct, 32)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  Discount := float32(Discnt)

  TtlAmount := Amount - Discount
  // It's a order request, so handle the form submission.
  // [START new_order]
  order := Order{
    UserId:  UserId,
    Amount: Amount,
    Discount: Discount,
    TtlAmount: TtlAmount,
    OrdStatus: "Placed",
    CreatedAt: time.Now(),
  }
  // [END new_order]

  // [START new_key]
  key := datastore.NewIncompleteKey(aectx, "Order", nil)
  // [END new_key]
  // [START add_order]
  ordKey, err := datastore.Put(aectx, key, &order)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END add_order]
  order.Id = ordKey.IntID()
  ginctx.JSON(http.StatusOK, gin.H{"order": order})
}

func OrderTransactionV1(ginctx *gin.Context) {
  // [START new_context]
  aectx := appengine.NewContext(ginctx.Request)
  // [END new_context]

  // [START get_Id]
  //path := req.URL.Path
  //Id := strings.TrimPrefix(path, "/updateorder/")
  Id := ginctx.Param("OrdId")
  ordId, err := strconv.ParseInt(Id, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END get_Id]

  // [START exist_key]
  key := datastore.NewKey(aectx, "Order", "", ordId, nil)
  // [END exist_key]

  var order Order
  err = datastore.Get(aectx, key, &order)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // It's a order request, so handle the form submission.
  // [START exist_order]
  order.Id = ordId
  order.TraxnId = ginctx.PostForm("TraxnId")
  order.TraxnStatus = ginctx.PostForm("TraxnStatus")
  order.TrxnDate = time.Now()
  // [END exist_order]

  if order.TraxnId == "" {
    ginctx.JSON(http.StatusOK, gin.H{"Error": "TraxnactionId must be there"})
    return
  }
  if order.TraxnStatus == "" {
    ginctx.JSON(http.StatusOK, gin.H{"Error": "TraxnactionStatus must be there"})
    return
  }

  // [START add_order]
  fmt.Println(order)
  ordKey, err := datastore.Put(aectx, key, &order)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END add_order]
  order.Id = ordKey.IntID()
  ginctx.JSON(http.StatusOK, gin.H{"order": order})
}

func ChangeOrderStatusV1(ginctx *gin.Context) {
  // [START new_context]
  aectx := appengine.NewContext(ginctx.Request)
  // [END new_context]

  // [START get_Id]
  //path := req.URL.Path
  //Id := strings.TrimPrefix(path, "/updateorder/")
  Id := ginctx.Param("OrdId")
  ordId, err := strconv.ParseInt(Id, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END get_Id]

  // [START exist_key]
  key := datastore.NewKey(aectx, "Order", "", ordId, nil)
  // [END exist_key]

  var order Order
  err = datastore.Get(aectx, key, &order)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // It's a order request, so handle the form submission.
  // [START exist_order]
  order.Id = ordId
  order.OrdStatus = ginctx.PostForm("OrdStatus")
  order.UpdatedAt = time.Now()
  // [END exist_order]

  if order.OrdStatus == "" {
    ginctx.JSON(http.StatusOK, gin.H{"Error": "Order Status must be there"})
    return
  }

  // [START add_order]
  fmt.Println(order)
  ordKey, err := datastore.Put(aectx, key, &order)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END add_order]
  order.Id = ordKey.IntID()
  ginctx.JSON(http.StatusOK, gin.H{"order": order})
}

func DeleteOrderV1(ginctx *gin.Context) {
  // [START new_context]
  aectx := appengine.NewContext(ginctx.Request)
  // [END new_context]

  // [START get_Id]
  //path := req.URL.Path
  //Id := strings.TrimPrefix(path, "/deleteorder/")
  Id := ginctx.Param("OrdId")
  ordId, err := strconv.ParseInt(Id, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END get_Id]

  // [START new_key]
  key := datastore.NewKey(aectx, "Order", "", ordId, nil)
  // [END new_key]
  // [START add_order]
  err = datastore.Delete(aectx, key)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END add_order]
  ginctx.JSON(http.StatusOK, gin.H{"Message": "Successfully deleted"})
}

