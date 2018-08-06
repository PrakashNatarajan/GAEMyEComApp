// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Polmorphism is a piece of code or a program behave in different ways according to the different datas.
// Sample datastore demonstrates use of the cloud.google.com/go/datastore package from App Engine flexible.
package mainproducts

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


func ProductsList(ginctx *gin.Context) {
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

func CreateProduct(ginctx *gin.Context) {
  // [START new_context]
  aectx := appengine.NewContext(ginctx.Request)
  // [END new_context]

  prc := ginctx.PostForm("Price")
  Prc, err := strconv.ParseFloat(prc, 32)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  Price := float32(Prc)

  discnt := ginctx.PostForm("Discount")
  Discnt, err := strconv.ParseFloat(discnt, 32)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  Discount := float32(Discnt)

  // It's a product request, so handle the form submission.
  // [START new_product]
  product := Product{
    Title:  ginctx.PostForm("Title"),
    Description: ginctx.PostForm("Description"),
    Price: Price,
    Discount: Discount,
    ImageUrl: ginctx.PostForm("ImageUrl"),
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
  // [END new_product]

  if product.Title == "" {
    ginctx.JSON(http.StatusOK, gin.H{"Error": "Title must be there"})
    return
  }
  // [START new_key]
  key := datastore.NewIncompleteKey(aectx, "Product", nil)
  // [END new_key]
  // [START add_product]
  prdKey, err := datastore.Put(aectx, key, &product)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END add_product]
  product.Id = prdKey.IntID()
  ginctx.JSON(http.StatusOK, gin.H{"Product": product})
}

func UpdateProduct(ginctx *gin.Context) {
  // [START new_context]
  aectx := appengine.NewContext(ginctx.Request)
  // [END new_context]

  // [START get_Id]
  //path := req.URL.Path
  //Id := strings.TrimPrefix(path, "/UpdateProduct/")
  Id := ginctx.Param("Id")
  prdId, err := strconv.ParseInt(Id, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END get_Id]

  // [START exist_key]
  key := datastore.NewKey(aectx, "Product", "", prdId, nil)
  // [END exist_key]

  var product Product
  err = datastore.Get(aectx, key, &product)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // It's a product request, so handle the form submission.
  // [START exist_product]
  product.Id = prdId
  product.Title = ginctx.PostForm("title")
  product.Description = ginctx.PostForm("Description")
  product.UpdatedAt = time.Now()
  // [END exist_product]

  if product.Title == "" {
    ginctx.JSON(http.StatusOK, gin.H{"Error": "Title must be there"})
    return
  }

  // [START add_product]
  fmt.Println(product)
  prdKey, err := datastore.Put(aectx, key, &product)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END add_product]
  product.Id = prdKey.IntID()
  ginctx.JSON(http.StatusOK, gin.H{"Product": product})
}

func DeleteProduct(ginctx *gin.Context) {
  // [START new_context]
  aectx := appengine.NewContext(ginctx.Request)
  // [END new_context]

  // [START get_Id]
  //path := req.URL.Path
  //Id := strings.TrimPrefix(path, "/DeleteProduct/")
  Id := ginctx.Param("Id")
  prdId, err := strconv.ParseInt(Id, 10, 64)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END get_Id]

  // [START new_key]
  key := datastore.NewKey(aectx, "Product", "", prdId, nil)
  // [END new_key]
  // [START add_product]
  err = datastore.Delete(aectx, key)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }
  // [END add_product]
  ginctx.JSON(http.StatusOK, gin.H{"Message": "Successfully deleted"})
}


type Item struct {
  ProdId  int64 `json:"ProdId"`
  Quantity  int `json:"Quantity"`
  CreatedAt time.Time `json:"CreatedAt"`
}

func UpdateProductOrdCnt(ginctx *gin.Context) {
  aectx := appengine.NewContext(ginctx.Request)

  dateTime := time.Now()
  limit := 100

  query := datastore.NewQuery("Item").Filter("CreatedAt <", dateTime)
  itemCount, err := query.Count(aectx)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }

  pagesCount := (itemCount/limit) + 1

  for page := 0; page < pagesCount; page++ {
    err := ProductOrdersCount(ginctx, dateTime, limit, page)
    if err != nil {
      ginctx.JSON(503, gin.H{"Error": err.Error()})
      return
      //break
    }
  }

  ginctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated products orders count"})
}

func ProductOrdersCount(ginctx *gin.Context, dateTime time.Time, limit, page int) (err error) {
  aectx := appengine.NewContext(ginctx.Request)
  query := datastore.NewQuery("Item").Project("ProdId").Filter("CreatedAt <", dateTime).Offset(page).Limit(limit)
  var prodIds []Item
  prodKeys, err := query.GetAll(aectx, &prodIds)
  if err != nil {
    return err
  }
  fmt.Println(prodKeys)
  var items []Item
  var product Product
  for inx, _ := range prodKeys {
    ProdId := prodIds[inx].ProdId
    query := datastore.NewQuery("Item").Project("ProdId").Project("Quantity").Filter("CreatedAt <", dateTime).Filter("ProdId =", ProdId).Limit(1000)
    _, err := query.GetAll(aectx, &items)
    if err != nil {
      return err
    }

    OrderCount := 0
    for inx, _ := range items {
      OrderCount += items[inx].Quantity
    }

    key := datastore.NewKey(aectx, "Product", "", ProdId, nil)
    // [END exist_key]

    err = datastore.Get(aectx, key, &product)
    if err != nil {
      return err
    }
    product.OrderCount = int32(OrderCount)
    _, err = datastore.Put(aectx, key, &product)
    if err != nil {
      return err
    }
  }
  return err
}