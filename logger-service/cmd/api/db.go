package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongo()(*mongo.Client,error){
	//create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
})

//connect
c,err := mongo.Connect(context.TODO(),clientOptions)
if err != nil{
	log.Println("Error connecting:", err)
	return nil, err
}

log.Println("Connected to mongo!")
return c, nil

}