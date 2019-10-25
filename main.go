package main

import (
	"fmt"
	"os"
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

	store := persistence.NewInMemoryStore(time.Second)

	api := r.Group("/scg")
	{
		api.GET("/puzzle", cache.CachePage(store, time.Minute, PuzzleHandler))
		api.GET("/food", cache.CachePage(store, time.Minute, FoodHandler))
	}

	r.Run(getPort())
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No port in Heroku")
	}

	return ":" + port
}
