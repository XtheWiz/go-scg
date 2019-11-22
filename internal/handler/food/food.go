package food

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"

	"go-scg/internal/model/place"

	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
)

//https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=13.8035134,100.5373821&radius=2000&type=food&keyword=ส้มตำ&key={{Place API Key}}
func FoodHandler(c *gin.Context) {
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
	queryString.Set("key", apiKey)
	destURL.RawQuery = queryString.Encode()

	fmt.Println("URL = " + destURL.String())

	resp, err := http.Get(destURL.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	placeResp := new(PlacesSearchResponse)
	json.NewDecoder(resp.Body).Decode(placeResp)

	pretty.Println(placeResp)
	fmt.Println("Found total: ", len(placeResp.Results))
	fmt.Println("Status: ", placeResp.Status)

	if placeResp.Status == "OK" {
		placeResultReturn := []ReturnPlaceResult{}
		for _, place := range placeResp.Results {
			p := ReturnPlaceResult{}
			p.Name = place.Name
			if len(place.Photos) > 0 {
				p.PhotoRef = place.Photos[0].PhotoReference
			}
			p.Lat = place.Geometry.Location.Lat
			p.Lng = place.Geometry.Location.Lng
			p.Vicinity = place.Vicinity
			p.Distance = p.findDistance()

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

//https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=13.8035134,100.5373821&radius=2000&type=food&keyword=ส้มตำ&key={{Place API Key}}
func FoodHandler(c *gin.Context) {
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
	queryString.Set("key", apiKey)
	destURL.RawQuery = queryString.Encode()

	fmt.Println("URL = " + destURL.String())

	resp, err := http.Get(destURL.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	placeResp := new(PlacesSearchResponse)
	json.NewDecoder(resp.Body).Decode(placeResp)

	pretty.Println(placeResp)
	fmt.Println("Found total: ", len(placeResp.Results))
	fmt.Println("Status: ", placeResp.Status)

	if placeResp.Status == "OK" {
		placeResultReturn := []ReturnPlaceResult{}
		for _, place := range placeResp.Results {
			p := ReturnPlaceResult{}
			p.Name = place.Name
			if len(place.Photos) > 0 {
				p.PhotoRef = place.Photos[0].PhotoReference
			}
			p.Lat = place.Geometry.Location.Lat
			p.Lng = place.Geometry.Location.Lng
			p.Vicinity = place.Vicinity
			p.Distance = p.findDistance()

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
