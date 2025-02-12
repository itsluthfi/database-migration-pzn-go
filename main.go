package main

import (
	"net/http"
	"rest-api-pzn-go/app"
	"rest-api-pzn-go/controller"
	"rest-api-pzn-go/helper"
	"rest-api-pzn-go/middleware"
	"rest-api-pzn-go/repository"
	"rest-api-pzn-go/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
