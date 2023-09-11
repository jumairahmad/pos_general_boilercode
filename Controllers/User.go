package Controllers

import (
	"POS/ApiHelpers"
	Models "POS/Models"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUser(&user)

	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func AddNewUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.AddNewUser(&user)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, err)
	}
}

func GetOneUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetOneUser(&user, id)

	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func PutOneUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetOneUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	}
	c.BindJSON(&user)
	err = Models.PutOneUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}
