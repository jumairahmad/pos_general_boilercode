package Controllers

import (
	"POS/ApiHelpers"
	Models "POS/Models"

	"github.com/gin-gonic/gin"
)

func ListOrders(c *gin.Context) {
	var orders []Models.Orders
	err := Models.GetAllOrders(&orders)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, orders)
	} else {
		ApiHelpers.RespondJSON(c, 200, orders)
	}
}

func AddNewOrders(c *gin.Context) {
	var orders Models.Orders
	c.BindJSON(&orders)
	err := Models.AddNewOrders(&orders)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, orders)
	} else {
		ApiHelpers.RespondJSON(c, 200, orders)
	}
}

func GetOneOrders(c *gin.Context) {
	id := c.Params.ByName("id")
	var orders Models.Orders
	err := Models.GetOneOrders(&orders, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, orders)
	} else {
		ApiHelpers.RespondJSON(c, 200, orders)
	}
}

func PutOneOrders(c *gin.Context) {
	var orders Models.Orders
	id := c.Params.ByName("id")
	err := Models.GetOneOrders(&orders, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, orders)
	}
	c.BindJSON(&orders)
	err = Models.PutOneOrders(&orders, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, orders)
	} else {
		ApiHelpers.RespondJSON(c, 200, orders)
	}
}

func DeleteOrders(c *gin.Context) {
	var orders Models.Orders
	id := c.Params.ByName("id")
	err := Models.DeleteOrders(&orders, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, orders)
	} else {
		ApiHelpers.RespondJSON(c, 200, orders)
	}
}
