package Controllers

import (
	"POS/ApiHelpers"
	Models "POS/Models"

	"github.com/gin-gonic/gin"
)

func ListQuantity(c *gin.Context) {
	var quantity []Models.Quantity
	err := Models.GetAllQuantity(&quantity)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, quantity)
	} else {
		ApiHelpers.RespondJSON(c, 200, quantity)
	}
}

func AddNewQuantity(c *gin.Context) {
	var quantity Models.Quantity
	c.BindJSON(&quantity)
	err := Models.AddNewQuantity(&quantity)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, quantity)
	} else {
		ApiHelpers.RespondJSON(c, 200, quantity)
	}
}

func GetOneQuantity(c *gin.Context) {
	id := c.Params.ByName("id")
	var quantity Models.Quantity
	err := Models.GetOneQuantity(&quantity, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, quantity)
	} else {
		ApiHelpers.RespondJSON(c, 200, quantity)
	}
}

func PutOneQuantity(c *gin.Context) {
	var quantity Models.Quantity
	id := c.Params.ByName("id")
	err := Models.GetOneQuantity(&quantity, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, quantity)
	}
	c.BindJSON(&quantity)
	err = Models.PutOneQuantity(&quantity, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, quantity)
	} else {
		ApiHelpers.RespondJSON(c, 200, quantity)
	}
}

func DeleteQuantity(c *gin.Context) {
	var quantity Models.Quantity
	id := c.Params.ByName("id")
	err := Models.DeleteQuantity(&quantity, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, quantity)
	} else {
		ApiHelpers.RespondJSON(c, 200, quantity)
	}
}
