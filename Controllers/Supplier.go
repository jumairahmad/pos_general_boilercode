package Controllers

import (
	"POS/ApiHelpers"
	Models "POS/Models"

	"github.com/gin-gonic/gin"
)

func ListSupplier(c *gin.Context) {
	var supplier []Models.Supplier
	err := Models.GetAllSupplier(&supplier)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, supplier)
	} else {
		ApiHelpers.RespondJSON(c, 200, supplier)
	}
}

func AddNewSupplier(c *gin.Context) {
	var supplier Models.Supplier
	c.BindJSON(&supplier)
	err := Models.AddNewSupplier(&supplier)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, supplier)
	} else {
		ApiHelpers.RespondJSON(c, 200, supplier)
	}
}

func GetOneSupplier(c *gin.Context) {
	id := c.Params.ByName("id")
	var supplier Models.Supplier
	err := Models.GetOneSupplier(&supplier, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, supplier)
	} else {
		ApiHelpers.RespondJSON(c, 200, supplier)
	}
}

func PutOneSupplier(c *gin.Context) {
	var supplier Models.Supplier
	id := c.Params.ByName("id")
	err := Models.GetOneSupplier(&supplier, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, supplier)
	}
	c.BindJSON(&supplier)
	err = Models.PutOneSupplier(&supplier, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, supplier)
	} else {
		ApiHelpers.RespondJSON(c, 200, supplier)
	}
}

func DeleteSupplier(c *gin.Context) {
	var supplier Models.Supplier
	id := c.Params.ByName("id")
	err := Models.DeleteSupplier(&supplier, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, supplier)
	} else {
		ApiHelpers.RespondJSON(c, 200, supplier)
	}
}
