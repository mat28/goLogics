package users

import(
	"github.com/gin-gonic/gin"
	"Go-Rubix/controllers/users"
)



func SetUserRoutes(group *gin.RouterGroup) {
	group.GET("/users", users.GetUsers())
	group.GET("/users/:id", users.GetUsers())
	group.POST("/users", users.GetUsers())
	group.PUT("/users/:id", users.GetUsers())
	group.DELETE("/users/:id", users.GetUsers())

	return group
}