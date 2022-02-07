package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var customers = []customer{
	{ID: "1", Name: "mirror"},
	{ID: "2", Name: "blackmirror"},
}

func main() {
	fmt.Println("Hello, Service !")
	c := NewCustomer("1", "Yoohoo")
	fmt.Print(c.Name)

	router := gin.Default()
	router.GET("/getCustomer/:id", GetCustomer)
	router.GET("/", Homepage)
	router.Run("localhost:8080")

}

func Homepage(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Welcome to GO server"})
}

func GetCustomer(c *gin.Context) {

	id := c.Param("id")

	for _, customer := range customers {
		if customer.ID == id {

			c.IndentedJSON(http.StatusOK, customer)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No customer found with given id"})

}
