package rest

import (
	"net/http"
	"strconv"

	users "github.com/HongXiangZuniga/CrudGoExample/pkg/Users"
	"github.com/gin-gonic/gin"
)

type UsersHandlers interface {
	GetUserById(*gin.Context)
	GetUsersByField(*gin.Context)
}

type UsersPort struct {
	UsersServices users.UserServices
}

func NewUsersHandler(usersServices users.UserServices) UsersHandlers {
	return &UsersPort{usersServices}
}

func (port *UsersPort) GetUserById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"user": nil, "error": "Id not valid"})
		return
	}
	user, err := port.UsersServices.GetUserById(idInt)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"user": nil, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user, "error": nil})
	return
}

func (port *UsersPort) GetUsersByField(ctx *gin.Context) {
	field := ctx.Params.ByName("field")
	if field == "" {
		ctx.JSON(http.StatusOK, gin.H{"users": nil, "error": "field not found"})
		return
	}
	value := ctx.Params.ByName("value")
	if value == "" {
		ctx.JSON(http.StatusOK, gin.H{"users": nil, "error": "value not found"})
		return
	}
	users, err := port.UsersServices.GetUsersByField(field, value)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"users": nil, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users, "error": nil})
	return
}
