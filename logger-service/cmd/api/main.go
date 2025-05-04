package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)


const(
	webPort = "8083"
	rpcPort = "5001"
	mongoURL = "mongodb://mongo:27017"
	grpcPort ="50001"
)

var client *mongo.Client

type Config struct {
Models data.Models
}

func main(){
	//connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	//create a context inorder to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//close connection
	defer func(){
		if err = client.Disconnect(ctx);err != nil{
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}
	// go app.serve()
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil{
		log.Panic()
	}
	
}

func(app *Config) serve(){
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil{
		log.Panic()
	}
}