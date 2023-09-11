package Controllers

import (
	"POS/ApiHelpers"
	Models "POS/Models"

	"github.com/gin-gonic/gin"
)

func ListProduct(c *gin.Context) {
	var product []Models.Product
	err := Models.GetAllProduct(&product)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, product)
	} else {
		ApiHelpers.RespondJSON(c, 200, product)
	}
}

func AddNewProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.AddNewProduct(&product)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, product)
	} else {
		ApiHelpers.RespondJSON(c, 200, product)
	}
}

func GetOneProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetOneProduct(&product, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, product)
	} else {
		ApiHelpers.RespondJSON(c, 200, product)
	}
}

func PutOneProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetOneProduct(&product, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, product)
	}
	c.BindJSON(&product)
	err = Models.PutOneProduct(&product, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, product)
	} else {
		ApiHelpers.RespondJSON(c, 200, product)
	}
}

func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, product)
	} else {
		ApiHelpers.RespondJSON(c, 200, product)
	}
}
