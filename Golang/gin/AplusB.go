package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AAndB struct {
	A int
	B int
}

func main() {
	router := gin.Default()
	router.POST("/add", func(c *gin.Context) {
		a := c.Query("a")
		b := c.Query("b")
		if a == "" || b == "" {
			c.JSON(http.StatusOK, gin.H{
				"errorMessage": "please input a or b",
			})
		} else {
			aInt, _ := strconv.Atoi(a)
			bInt, _ := strconv.Atoi(b)
			c.JSON(http.StatusOK, gin.H{
				"result": strconv.Itoa(aInt + bInt),
			})
		}
	})
	router.Run()
}
