package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const queryStringPrefix = "var"

func getValueFromPosition(c *gin.Context, position string) int {
	value, err := strconv.Atoi(c.Query(queryStringPrefix + position))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSequence": "n",
			"message":    "1st number is not digit",
		})
	}
	return value
}

func isValidSequence(var1, var2, var3, var4 int) bool {
	return (var2-var1 != 4) && (var3-var2 != 6) && (var4-var3 != 8)
}

func returnCorrectVariables(var1, var4 int) (x, y, z int) {
	x = var1 - 2
	y = var4 + 10
	z = y + 12

	return
}

// PuzzleHandler will receive query string var1, var2, var3 and var4
// if there is no value or parse to int error, the handler will return error
// if the sequence is invalid, the handler will return error
// Query String:
//		var1: first input number in the sequence
//		var2: second input number in the sequence
//		var3: third input number in the sequence
//		var4: fourth input number in the sequence
// Return:
//		isSequence: flag to tell that this number sequence is valid or not
//		message: additional message
func PuzzleHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var1 := getValueFromPosition(c, "1")
	var2 := getValueFromPosition(c, "2")
	var3 := getValueFromPosition(c, "3")
	var4 := getValueFromPosition(c, "4")

	if isValidSequence(var1, var2, var3, var4) {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSequence": "n",
			"message":    "This is not desire sequence",
		})
	} else {
		x, y, z := returnCorrectVariables(var1, var4)
		c.JSON(http.StatusOK, gin.H{
			"isSequence": "y",
			"x":          x,
			"y":          y,
			"z":          z,
			"message":    fmt.Sprintf("X = %d, Y = %d, Z = %d", x, y, z),
		})
	}
}
