package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spoonbuoy/lcdb/api"
	"github.com/spoonbuoy/lcdb/clients"
)

func main() {

	initServices()
	r := gin.Default()

	api.InitRESTRoutes(r)

	r.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	fmt.Println("no bullshit programmer...")
	fmt.Println("lcdb running")
	fmt.Println("lcdb running")
	fmt.Println("lcdb running")
	fmt.Println("lcdb running")
	fmt.Println("lcdb running")
	fmt.Println("lcdb running")

	r.Run(":8081")
	return
}

// interfaces
// -------------------
// cacherouter
// cacheservice
// cacherepo

// dbrouter
// dbservice
// dbrepo

// API
// -------------------
// -cache
// init
// Get
// set
// DeleteCacheKey

// -db
// init

func initServices() {
	clients.Init()
	// handle dep injec here
	return
}
