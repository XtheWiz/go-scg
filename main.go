package main

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	store := persistence.NewInMemoryStore(time.Hour)

	api := r.Group("/scg")
	{
		api.GET("/puzzle", cache.CachePage(store, time.Hour, PuzzleHandler))
		api.GET("/food", cache.CachePage(store, time.Hour, FoodHandler))
	}

	r.Run(":3111")
}
