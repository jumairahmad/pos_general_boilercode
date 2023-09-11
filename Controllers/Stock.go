package Controllers

import (
	"POS/ApiHelpers"
	Models "POS/Models"

	"github.com/gin-gonic/gin"
)

func ListStock(c *gin.Context) {
	var stock []Models.Stock
	err := Models.GetAllStock(&stock)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, stock)
	} else {
		ApiHelpers.RespondJSON(c, 200, stock)
	}
}

func AddNewStock(c *gin.Context) {
	var stock Models.Stock
	c.BindJSON(&stock)
	err := Models.AddNewStock(&stock)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, stock)
	} else {
		ApiHelpers.RespondJSON(c, 200, stock)
	}
}

func GetOneStock(c *gin.Context) {
	id := c.Params.ByName("id")
	var stock Models.Stock
	err := Models.GetOneStock(&stock, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, stock)
	} else {
		ApiHelpers.RespondJSON(c, 200, stock)
	}
}

func PutOneStock(c *gin.Context) {
	var stock Models.Stock
	id := c.Params.ByName("id")
	err := Models.GetOneStock(&stock, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, stock)
	}
	c.BindJSON(&stock)
	err = Models.PutOneStock(&stock, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, stock)
	} else {
		ApiHelpers.RespondJSON(c, 200, stock)
	}
}

func DeleteStock(c *gin.Context) {
	var stock Models.Stock
	id := c.Params.ByName("id")
	err := Models.DeleteStock(&stock, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, stock)
	} else {
		ApiHelpers.RespondJSON(c, 200, stock)
	}
}
