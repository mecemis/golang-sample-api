package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang-sample-api/app"
	"golang-sample-api/configs"
	"golang-sample-api/repositories"
	"golang-sample-api/services"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDb()
	dbClient := configs.GetCollection(configs.Db, "todos")

	TodoRepositoryDb := repositories.NewTodoRepositoryDb(dbClient)

	td := app.TodoHandler{TodoService: services.NewTodoService(TodoRepositoryDb)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todo", td.GetAllTodo)
	appRoute.Delete("api/todo/:id", td.DeleteTodo)
	err := appRoute.Listen(":8080")
	if err != nil {
		fmt.Println(err)
	}

}

/*
go get -u github.com/gofiber/fiber/v2
go get go.mongodb.org/mongo-driver/x/mongo/driver/ocsp@v1.9.1
go get golang.org/x/sync/errgroup
go get  github.com/golang/mock/mockgen/model
go install github.com/golang/mock/mockgen@v1.6.0
go get github.com/stretchr/testify/assert
go get github.com/joho/godotenv
go.mongodb.org/mongo-driver/mongo/options
*/
