package main


import (
	"github.com/gin-gonic/gin"
	"Go-Rubix/routes"
)

func main(){
	r := gin.Default();
	r = routes.InitRouteV1(r)
	r.Run(":8182")
}
