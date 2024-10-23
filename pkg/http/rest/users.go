package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	users "github.com/HongXiangZuniga/CrudGoExample/pkg/Users"
	"github.com/HongXiangZuniga/CrudGoExample/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UsersHandlers interface {
	GetUserById(*gin.Context)
	GetUsersByField(*gin.Context)
	GetAllUsers(*gin.Context)
	DeleteUser(*gin.Context)
	CreateUser(*gin.Context)
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
		ctx.JSON(http.StatusNotFound, gin.H{"user": nil, "error": "Id not valid"})
		return
	}
	user, err := port.UsersServices.GetUserById(idInt)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"user": nil, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (port *UsersPort) GetUsersByField(ctx *gin.Context) {
	field := ctx.Params.ByName("field")
	if field == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"users": nil, "error": "field not found"})
		return
	}
	value := ctx.Params.ByName("value")
	if value == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"users": nil, "error": "value not found"})
		return
	}
	page, isExist := ctx.GetQuery("page")
	if !isExist {
		page = "0"
	}
	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"users": nil, "error": utils.PageNotValid().Error()})
		return
	}
	users, err := port.UsersServices.GetUsersByField(field, value, intPage)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"users": nil, "error": utils.PageNotValid().Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (port *UsersPort) GetAllUsers(ctx *gin.Context) {
	page, isExist := ctx.GetQuery("page")
	if !isExist {
		page = "0"
	}
	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"users": nil, "error": "page no valid"})
		return
	}
	users, err := port.UsersServices.GetAllUser(intPage)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"users": nil, "error": err.Error()})
		return
	}
	nextUrl := ctx.Request.URL.String()
	prevUrl := ctx.Request.URL.String()
	if port.UsersServices.NextPageExist(intPage) == nil {
		nextUrl = ctx.Request.URL.Path + "?page=" + strconv.Itoa(intPage+1)
	}
	if intPage != 1 {
		prevUrl = ctx.Request.URL.Path + "?page=" + strconv.Itoa(intPage-1)
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users, "next": nextUrl, "prev": prevUrl})
}

func (port *UsersPort) DeleteUser(ctx *gin.Context) {
	parameterid := ctx.Params.ByName("id")
	if parameterid == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}
	id, err := strconv.Atoi(parameterid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}
	err = port.UsersServices.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (port *UsersPort) CreateUser(ctx *gin.Context) {
	var newUser users.User
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	err = json.Unmarshal([]byte(bodyString), &newUser)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	if newUser.Email == "" {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "Email is empty"})
		return
	}
	if newUser.Age == 0 {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "Age is empty or is 0"})
		return
	}
	if newUser.Name == "" {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "Name is empty"})
		return
	}
	if newUser.Country == "" {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "Country is empty"})
		return
	}
	err = port.UsersServices.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
