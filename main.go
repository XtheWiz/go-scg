package main

import (
	"fmt"
	"os"
	"time"

	"go-scg/internal/handler"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//For test, allow all origin
	r.Use(cors.Default())
	store := persistence.NewInMemoryStore(time.Second)

	api := r.Group("/scg")
	{
		api.GET("/puzzle", cache.CachePage(store, time.Minute, handler.PuzzleHandler))
		api.GET("/food", cache.CachePage(store, time.Minute, handler.FoodHandler))
	}

	r.Run(getPort())
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No port in environment variable")
	}

	return ":" + port
}
