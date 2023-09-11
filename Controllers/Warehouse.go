package Controllers

import (
	"POS/ApiHelpers"
	Models "POS/Models"

	"github.com/gin-gonic/gin"
)

func ListWarehouse(c *gin.Context) {
	var warehouse []Models.Warehouse
	err := Models.GetAllWarehouse(&warehouse)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, warehouse)
	} else {
		ApiHelpers.RespondJSON(c, 200, warehouse)
	}
}

func AddNewWarehouse(c *gin.Context) {
	var warehouse Models.Warehouse
	c.BindJSON(&warehouse)
	err := Models.AddNewWarehouse(&warehouse)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, warehouse)
	} else {
		ApiHelpers.RespondJSON(c, 200, warehouse)
	}
}

func GetOneWarehouse(c *gin.Context) {
	id := c.Params.ByName("id")
	var warehouse Models.Warehouse
	err := Models.GetOneWarehouse(&warehouse, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, warehouse)
	} else {
		ApiHelpers.RespondJSON(c, 200, warehouse)
	}
}

func PutOneWarehouse(c *gin.Context) {
	var warehouse Models.Warehouse
	id := c.Params.ByName("id")
	err := Models.GetOneWarehouse(&warehouse, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, warehouse)
	}
	c.BindJSON(&warehouse)
	err = Models.PutOneWarehouse(&warehouse, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, warehouse)
	} else {
		ApiHelpers.RespondJSON(c, 200, warehouse)
	}
}

func DeleteWarehouse(c *gin.Context) {
	var warehouse Models.Warehouse
	id := c.Params.ByName("id")
	err := Models.DeleteWarehouse(&warehouse, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, warehouse)
	} else {
		ApiHelpers.RespondJSON(c, 200, warehouse)
	}
}
