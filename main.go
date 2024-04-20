package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//		controllers.Start()
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file is found")
	}
	uri := os.Getenv("MONGODB_URL")
	if uri == "" {
		log.Fatal("Set the MONGODB_URL in the env file")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err)
		return
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Panic(err)
			return
		}
	}()
	collection := client.Database("tronics").Collection("products")
	// res, err := collection.InsertOne(context.Background(), models.Galaxy)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
	res, err := collection.InsertOne(context.Background(), bson.M{
		"name":    "pradeep",
		"surname": "kumar",
		"hobbies": bson.A{"sports", "learning"},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
}
