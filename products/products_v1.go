// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Polmorphism is a piece of code or a program behave in different ways according to the different datas.
// Sample datastore demonstrates use of the cloud.google.com/go/datastore package from App Engine flexible.
package mainproducts

import (
    "strconv"
	  "net/http"
	  //"encoding/json"
    "github.com/gin-gonic/gin"
	  "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
)


func ProductsOrdSqnc(ginctx *gin.Context) {
  aectx := appengine.NewContext(ginctx.Request)
  var err error
  ordSqnce := ginctx.Param("OrdSqnce")
  OrdSqnce := OrderSquence(ginctx, ordSqnce)
  if OrdSqnce == "" {
    ginctx.JSON(http.StatusOK, gin.H{"Error": "Order Squence not found"})
    return
  }
  query := datastore.NewQuery("Product").Order(OrdSqnce).Limit(10)
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

func OrderSquence(ginctx *gin.Context, ordSqnce string) (string) {
  var OrderSquences = map[string]string {"PriceAsc": "Price", "PriceDesc": "-Price", "DiscountAsc": "Discount", "DiscountDesc": "-Discount", "Popularity": "-OrderCount"}
  OrdSqnce := OrderSquences[ordSqnce]
  return OrdSqnce
}

func GetProductV1(ginctx *gin.Context) {
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
  prdKey := datastore.NewKey(aectx, "Product", "", prdId, nil)
  // [END exist_key]

  var product Product
  err = datastore.Get(aectx, prdKey, &product)
  if err != nil {
    ginctx.JSON(http.StatusOK, gin.H{"Error": err.Error()})
    return
  }

  product.Id = prdKey.IntID()
  ginctx.JSON(http.StatusOK, gin.H{"Product": product})
}