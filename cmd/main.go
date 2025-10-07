package main

import (
	"context"
	"fmt"

	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/handler"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/service"
)

func main() {
	// get .env configs
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ .env configs imported successfully")

	ctx := context.Background()
	
	// init db
	db, err := postgres.InitDB(ctx, conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ postgres db initialized successfully")

	// migrate dbs
	if err := db.MigrateDB(&domain.User{}); err != nil {
		panic(err)
	}
	fmt.Println("✅ dbs migrated successfully")

	// dependency injection
	userRepo := repository.InitUserRepository(db)
	userSvc := service.InitUserService(userRepo)
	userHandler := handler.InitUserHandler(userSvc)

	// routing
	r := handler.InitRouter(*userHandler)

	// run server
	if err := r.Start(conf.HTTP); err != nil {
		panic(err)
	}
}
