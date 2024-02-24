package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spoonbuoy/lcdb/interfaces/icache"
	"github.com/spoonbuoy/lcdb/interfaces/idatabase"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/", handler.GetCacheKey)
}

type httprouter struct {
	cacheService icache.ICacheService
	dbService    idatabase.IDatabaseService
}

func NewHttpRouter(cache icache.ICacheService, db idatabase.IDatabaseService) *httprouter {
	return &httprouter{
		cacheService: cache,
		dbService:    db,
	}
}

func (hr *httprouter) GetCacheKey(c *gin.Context) {
	key := ""
	_, _ = hr.cacheService.GetCacheKey(c, key)
	c.JSON(http.StatusOK, gin.H{})
}
