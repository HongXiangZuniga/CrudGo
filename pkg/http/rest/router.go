package rest

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewHandler(userHandler UsersHandlers) *gin.Engine {
	r := gin.Default()
	corsCfg := cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"Accept",
			"Origin",
			"Cache-Control",
			"X-Requested-With",
		},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}
	r.Use(cors.New(corsCfg))
	_ = r.Group("/")
	{
		users := r.Group("user")
		{
			users.Use(TokenAuthMiddleware())
			users.GET("/", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserById)
			users.GET("/search/:field/:value", userHandler.GetUsersByField)
			users.DELETE("/:id", userHandler.DeleteUser)
			users.POST("/", userHandler.CreateUser)
		}

	}
	return r
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")
	if requiredToken == "" {
		panic("Please set API_TOKEN environment variable")
	}
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			var response Response
			response.Message = "API token required"
			c.AbortWithStatusJSON(401, response)
			return
		}

		if token != requiredToken {
			var response Response
			response.Message = "Invalid API token"
			c.AbortWithStatusJSON(401, response)
			return
		}
		c.Next()
	}
}
