package main

import (
	"context"
	"fmt"
	"github.com/rbojan2000/central-library/config"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/rbojan2000/central-library/http"
	"github.com/rbojan2000/central-library/repository"
)

func main() {
	config := config.NewConfig()

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/", config.CentralLibraryDBHost, config.CentralLibraryDBPort)))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	// create a repository
	repository := repository.NewRepository(client.Database("users"))

	// create an http server
	server := http.NewServer(repository)

	// create a gin router
	router := gin.Default()
	{
		router.GET("/users/:id", server.GetUser)
		router.POST("/users", server.CreateUser)
		router.PUT("/users/:id", server.UpdateUserNumOfBooksRented)
		router.DELETE("/users/:id", server.DeleteUser)
	}

	// start the router
	router.Run(fmt.Sprintf(":%s", config.Port))
}
