package data

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Product defines the structure for an API product
type product struct {
	ID          string     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// GetProducts responds with a list of all products
func GetProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, productList)
}

// postProducts adds a product from JSON received
// in the request body.
func PostProducts(c *gin.Context) {
	var newProduct product

	// Call BindJSON to bind the received JSON
	// to newProduct.
	if err := c.BindJSON(&newProduct); err != nil {
		fmt.Print(err)
		return
	}
	//  Add the new product to the slice.
	productList = append(productList, newProduct)
}

// GetProductByID returns a product from a
// given ID if a product with that ID was found
func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of products, looking for
	// a product with the given ID.
	for _, p := range productList {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}

var productList = []product{
	{
		ID:          "1",
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          "2",
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
