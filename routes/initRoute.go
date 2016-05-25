package routes

import "github.com/gin-gonic/gin"
import "Go-Rubix/routes/users"

func InitRouteV1(r *gin.Engine) (*gin.Engine) {
	v1 := r.Group("v1")
	v1 = users.SetUserRoutes(v1)
	return r
}
