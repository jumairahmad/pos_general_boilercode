package Controllers

import (
	"POS/ApiHelpers"
	Models "POS/Models"

	"github.com/gin-gonic/gin"
)

func ListRoles(c *gin.Context) {
	var roles []Models.Roles
	err := Models.GetAllRoles(&roles)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, roles)
	} else {
		ApiHelpers.RespondJSON(c, 200, roles)
	}
}

func AddNewRoles(c *gin.Context) {
	var roles Models.Roles
	c.BindJSON(&roles)
	err := Models.AddNewRoles(&roles)
	if err != nil {
		ApiHelpers.RespondJSON(c, 501, roles)
	} else {
		ApiHelpers.RespondJSON(c, 200, roles)
	}
}

func GetOneRoles(c *gin.Context) {
	id := c.Params.ByName("id")
	var roles Models.Roles
	err := Models.GetOneRoles(&roles, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, roles)
	} else {
		ApiHelpers.RespondJSON(c, 200, roles)
	}
}

func PutOneRoles(c *gin.Context) {
	var roles Models.Roles
	id := c.Params.ByName("id")
	err := Models.GetOneRoles(&roles, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, roles)
	}
	c.BindJSON(&roles)
	err = Models.PutOneRoles(&roles, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, roles)
	} else {
		ApiHelpers.RespondJSON(c, 200, roles)
	}
}

func DeleteRoles(c *gin.Context) {
	var roles Models.Roles
	id := c.Params.ByName("id")
	err := Models.DeleteRoles(&roles, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, roles)
	} else {
		ApiHelpers.RespondJSON(c, 200, roles)
	}
}
