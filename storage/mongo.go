package mongo

//TODO: implement mongo storage perhaps?

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/mongodb/mongo-go-driver/bson"
// 	"github.com/mongodb/mongo-go-driver/bson/objectid"
// 	"github.com/mongodb/mongo-go-driver/mongo"
// )

// type Set struct {
// 	Id   objectid.ObjectID `bson:"_id,omitempty"`
// 	Name string
// }

// func init() {
// 	client, err := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db := client.Database("dominionshuffle")
// 	setCollection := db.Collection("set")

// 	_, err = setCollection.DeleteMany(
// 		context.Background(),
// 		bson.NewDocument(),
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {
// 	// connect to MongoDB
// 	client, err := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db := client.Database("dominionshuffle")
// 	setCollection := db.Collection("set")

// 	// write document
// 	itemWrite := Set{Name: "Dominion"}
// 	itemWrite.Id = objectid.New()
// 	fmt.Printf("itemWrite = %v\n", itemWrite)

// 	result, err := setCollection.InsertOne(context.Background(), itemWrite)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("result = %#v\n", result)

// 	// read documents
// 	cursor, err := setCollection.Find(
// 		context.Background(),
// 		// bson.NewDocument(bson.EC.String("name", "Dominion")),
// 		bson.NewDocument(),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cursor.Close(context.Background())

// 	itemRead := Set{}
// 	for cursor.Next(context.Background()) {
// 		err := cursor.Decode(&itemRead)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("itemRead = %v\n", itemRead)
// 	}

// 	fmt.Println("Hello world!")
// }
