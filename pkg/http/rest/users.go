package rest

import (
	"net/http"
	"strconv"

	users "github.com/HongXiangZuniga/CrudGoExample/pkg/Users"
	"github.com/gin-gonic/gin"
)

type UsersHandlers interface {
	GetUserById(*gin.Context)
	GetUserByCountry(*gin.Context)
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

func (port *UsersPort) GetUserByCountry(ctx *gin.Context) {
	country := ctx.Params.ByName("country")
	if country == "" {
		ctx.JSON(http.StatusOK, gin.H{"users": nil, "error": "country not found"})
		return
	}
	users, err := port.UsersServices.GetUsersByCountrys(country)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"users": nil, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users, "error": nil})
	return
}
