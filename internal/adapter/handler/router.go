package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/config"
)

type Router struct {
	*gin.Engine
}

func InitRouter(
	userHandler UserHandler,
	authHandler AuthHandler,
	empHandler EmployeeHandler,
) *Router {
	r := gin.New()

	// public route
	pb := r.Group("/api/v1")

	pb.POST("/register", userHandler.Register)
	pb.POST("/login", authHandler.Login)

	// user routes

	pb.GET("/employees", empHandler.GetEmployees)

	// private route (admin only)
	ad := pb.Group("/admin")

	ad.POST("/employees", empHandler.CreateEmployee)

	return &Router{r}
}

func (r *Router) Start(conf *config.HTTP) error {
	uri := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	if err := r.Run(uri); err != nil {
		return err
	}

	return nil
}
