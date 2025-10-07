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
) *Router {
	r := gin.New()

	// public route
	pb := r.Group("/api/v1")

	pb.POST("/register", userHandler.Register)

	return &Router{r}
}

func (r *Router) Start(conf *config.HTTP) error {
	uri := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	if err := r.Run(uri); err != nil {
		return err
	}

	return nil
}
