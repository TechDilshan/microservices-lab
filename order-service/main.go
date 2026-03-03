package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Order struct {
	ID        int    `json:"id"`
	Item      string `json:"item"`
	Quantity  int    `json:"quantity"`
	Customer  string `json:"customerId"`
	Status    string `json:"status"`
}

var orders []Order
var idCounter = 1

func main() {
	r := gin.Default()

	r.GET("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, orders)
	})

	r.POST("/orders", func(c *gin.Context) {
		var newOrder Order
		if err := c.BindJSON(&newOrder); err != nil {
			return
		}
		newOrder.ID = idCounter
		newOrder.Status = "PENDING"
		idCounter++
		orders = append(orders, newOrder)
		c.JSON(http.StatusCreated, newOrder)
	})

	r.GET("/orders/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for _, o := range orders {
			if o.ID == id {
				c.JSON(http.StatusOK, o)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	r.Run(":8082")
}