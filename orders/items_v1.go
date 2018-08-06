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


func OrderItemsV1(ginctx *gin.Context) {
  aectx := appengine.NewContext(ginctx.Request)
  var err error
  ordId := ginctx.Param("OrdId")
  OrdId, err := strconv.ParseInt(ordId, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  query := datastore.NewQuery("Item").Filter("OrderId =", OrdId).Limit(10)
  var items []Item
  keys, err := query.GetAll(aectx, &items)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  for inx, key := range keys {
    items[inx].Id = key.IntID()
  }
 
  ginctx.JSON(http.StatusOK, gin.H{"keys": keys, "items": items})
}

func UserItemsV1(ginctx *gin.Context) {
  aectx := appengine.NewContext(ginctx.Request)
  var err error
  usrId := ginctx.Param("UsrId")
  UsrId, err := strconv.ParseInt(usrId, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  query := datastore.NewQuery("Item").Filter("UserId =", UsrId).Limit(10)
  var items []Item
  keys, err := query.GetAll(aectx, &items)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  for inx, key := range keys {
    items[inx].Id = key.IntID()
  }
 
  ginctx.JSON(http.StatusOK, gin.H{"keys": keys, "items": items})
}

func CreateItemV1(ginctx *gin.Context) {
  // [START new_context]
  aectx := appengine.NewContext(ginctx.Request)
  // [END new_context]

  // [START get_Id]
  //path := req.URL.Path
  //Id := strings.TrimPrefix(path, "/updateitem/")
  UsrId := ginctx.Param("UsrId")
  UserId, err := strconv.ParseInt(UsrId, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  prdId := ginctx.PostForm("ProdId")
  PrdId, err := strconv.ParseInt(prdId, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END get_Id]
  Qnty := ginctx.PostForm("Quantity")
  Qunty, err := strconv.ParseInt(Qnty, 5, 32)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  Quantity := int32(Qunty)

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
  // It's a item request, so handle the form submission.
  // [START new_item]
  item := Item{
    UserId:  UserId,
    ProdId: PrdId,
    Quantity: Quantity,
    Amount: Amount,
    Discount: Discount,
    TtlAmount: TtlAmount,
    CreatedAt: time.Now(),
  }
  // [END new_item]

  // [START new_key]
  key := datastore.NewIncompleteKey(aectx, "Item", nil)
  // [END new_key]
  // [START add_item]
  itmKey, err := datastore.Put(aectx, key, &item)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END add_item]
  item.Id = itmKey.IntID()
  ginctx.JSON(http.StatusOK, gin.H{"item": item})
}

func UpdateItemsV1(ginctx *gin.Context) {
  // [START new_context]
  aectx := appengine.NewContext(ginctx.Request)
  // [END new_context]

  // [START get_Id]
  //path := req.URL.Path
  //Id := strings.TrimPrefix(path, "/updateitem/")
  UsrId := ginctx.Param("UsrId")
  UserId, err := strconv.ParseInt(UsrId, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  ordId := ginctx.PostForm("OrdId")
  if ordId == "" {
    ginctx.JSON(http.StatusOK, gin.H{"Error": "Order Id must be there"})
    return
  }
  OrdId, err := strconv.ParseInt(ordId, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END get_Id]

  query := datastore.NewQuery("Item").Filter("UserId =", UserId).Limit(10)
  var items []Item
  keys, err := query.GetAll(aectx, &items)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }

  // It's a item request, so handle the form submission.
  // [START exist_item]
  for inx, key := range keys {
    items[inx].Id = key.IntID()
    items[inx].OrderId = OrdId
  }
  // [END exist_item]

  // [START add_item]
  fmt.Println(items)
  _, err = datastore.PutMulti(aectx, keys, items)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END add_item]
  ginctx.JSON(http.StatusOK, gin.H{"Message": "Successfully Updated"})
}

func DeleteItemV1(ginctx *gin.Context) {
  // [START new_context]
  aectx := appengine.NewContext(ginctx.Request)
  // [END new_context]

  // [START get_Id]
  //path := req.URL.Path
  //Id := strings.TrimPrefix(path, "/deleteitem/")
  itmId := ginctx.Param("ItmId")
  ItmId, err := strconv.ParseInt(itmId, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END get_Id]

  // [START new_key]
  key := datastore.NewKey(aectx, "Item", "", ItmId, nil)
  // [END new_key]
  // [START add_item]
  err = datastore.Delete(aectx, key)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END add_item]
  ginctx.JSON(http.StatusOK, gin.H{"Message": "Successfully deleted"})
}

