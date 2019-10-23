package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const imgURL = "https://maps.googleapis.com/maps/api/place/photo?maxwidth=350&photoreference="
const apiKey = "YOUR_API_KEY"

func PuzzleHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var1, err := strconv.Atoi(c.Query("var1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSequence": "n",
			"message":    "1st number is not digit",
		})
		return
	}

	var2, err := strconv.Atoi(c.Query("var2"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSequence": "n",
			"message":    "2nd number is not digit",
		})
		return
	} else if var1 >= var2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSequence": "n",
			"message":    "var2 should not equal or less than var1",
		})
		return
	}

	var3, err := strconv.Atoi(c.Query("var3"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSequence": "n",
			"message":    "3rd number is not digit",
		})
		return
	} else if var2 >= var3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSequence": "n",
			"message":    "var3 should not equal or less than var2",
		})
		return
	}

	var4, err := strconv.Atoi(c.Query("var4"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSequence": "n",
			"message":    "4th number is not digit",
		})
		return
	} else if var3 >= var4 {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSequence": "n",
			"message":    "var4 should not equal or less than var3",
		})
		return
	}

	if var2-var1 != 4 ||
		var3-var2 != 6 ||
		var4-var3 != 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSequence": "n",
			"message":    "This is not desire sequence",
		})
	} else {
		x := var1 - 2
		y := var4 + 10
		z := y + 12
		c.JSON(http.StatusOK, gin.H{
			"isSequence": "y",
			"x":          x,
			"y":          y,
			"z":          z,
			"message":    fmt.Sprintf("X = %d, Y = %d, Z = %d", x, y, z),
		})
	}
}

//https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=13.8035134,100.5373821&radius=2000&type=food&keyword=ส้มตำ&key={{Place API Key}}
func FoodHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	foodType := c.Query("foodType")
	foodType = strings.TrimSpace(foodType)

	baseURL := "https://maps.googleapis.com"
	relativeURL := "/maps/api/place/nearbysearch/json"
	url, _ := url.Parse(baseURL)
	url.Path = path.Join(url.Path, relativeURL, "/")

	queryString := url.Query()
	queryString.Set("location", "13.8035134,100.5373821")
	queryString.Set("radius", "5000")
	queryString.Set("type", "food")
	queryString.Set("language", "th")
	if len(foodType) > 0 {
		queryString.Set("keyword", foodType)
	}
	queryString.Set("key", apiKey)
	url.RawQuery = queryString.Encode()

	fmt.Println("URL = " + url.String())

	resp, err := http.Get(url.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	placeResp := new(PlacesSearchResponse)
	json.NewDecoder(resp.Body).Decode(placeResp)
	// pretty.Println(placeResp)
	// fmt.Println("Found total: ", len(placeResp.Results))

	if placeResp.Status != "OK" ||
		len(placeResp.Results) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":   placeResp.Status,
			"foodlist": placeResp.Results,
		})
	} else {
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
	}
}
