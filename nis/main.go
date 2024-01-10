package main

import (
	"context"
	"fmt"
	"github.com/rbojan2000/nis/config"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/rbojan2000/nis/http"
	"github.com/rbojan2000/nis/repository"
)

func main() {
	config := config.NewConfig()

	db_connection := fmt.Sprintf("mongodb://%s:%s/", config.NisLibraryDBHost, config.NisLibraryDBPort)

	println(db_connection)

	client, err := mongo.NewClient(options.Client().ApplyURI(db_connection))

	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	// create a repository
	repository := repository.NewRepository(client.Database("borrows"))

	// create an http server
	server := http.NewServer(repository)

	// create a gin router
	router := gin.Default()
	{
		router.GET("/borrows/:id", server.GetBorrow)
		router.POST("/borrows", server.CreateBorrow)
		router.DELETE("/borrows/:id", server.DeleteBorrow)
	}

	// start the router
	router.Run(fmt.Sprintf(":%s", config.Port))
}
