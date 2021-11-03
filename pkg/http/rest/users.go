package rest

import (
	"net/http"
	"strconv"

	users "github.com/HongXiangZuniga/CrudGoExample/pkg/Users"
	"github.com/gin-gonic/gin"
)

type UsersHandlers interface {
	GetUserById(*gin.Context)
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
	user, err := port.UsersServices.GetUser(idInt)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"user": nil, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user, "error": nil})
	return
}
