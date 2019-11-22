package handler

import (
	"encoding/json"
	"fmt"
	"go-scg/internal/config"
	"go-scg/internal/model"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
)

// FoodHandler will receive query string foodType as optional
// Query String:
//		foodType (optional): Specify type of food for place api
// Return:
//		status: 	status return from place api
//		foodlist: list is restaurant return from place api
func FoodHandler(c *gin.Context) {
	config := config.LoadConfig("configs/config.yml")
	c.Header("Content-Type", "application/json")

	foodType := c.Query("foodType")
	foodType = strings.TrimSpace(foodType)

	fmt.Println("Food Type: " + foodType)

	baseURL := "https://maps.googleapis.com"
	relativeURL := "/maps/api/place/nearbysearch/json"
	destURL, _ := url.Parse(baseURL)
	destURL.Path = path.Join(destURL.Path, relativeURL, "/")

	queryString := destURL.Query()
	queryString.Set("location", "13.8035134,100.5373821")
	queryString.Set("radius", "5000")
	queryString.Set("type", "food")
	queryString.Set("language", "th")
	if len(foodType) > 0 {
		decodedValue, err := url.QueryUnescape(foodType)
		if err == nil {
			queryString.Add("keyword", decodedValue)
		}
	}
	queryString.Set("key", config.PlaceApiKey)
	destURL.RawQuery = queryString.Encode()

	fmt.Println("URL = " + destURL.String())

	resp, err := http.Get(destURL.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	placeResp := new(model.PlacesSearchResponse)
	json.NewDecoder(resp.Body).Decode(placeResp)

	pretty.Println(placeResp)
	fmt.Println("Found total: ", len(placeResp.Results))
	fmt.Println("Status: ", placeResp.Status)

	if placeResp.Status == "OK" {
		placeResultReturn := []model.ReturnPlaceResult{}
		for _, place := range placeResp.Results {
			p := model.ReturnPlaceResult{}
			p.Name = place.Name
			if len(place.Photos) > 0 {
				p.PhotoRef = place.Photos[0].PhotoReference
			}
			p.Lat = place.Geometry.Location.Lat
			p.Lng = place.Geometry.Location.Lng
			p.Vicinity = place.Vicinity
			p.Distance = p.FindDistance(config.BangsueLat, config.BangsueLng)

			placeResultReturn = append(placeResultReturn, p)
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   placeResp.Status,
			"foodlist": placeResultReturn,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":   placeResp.Status,
			"foodlist": placeResp.Results,
		})
	}
}
