package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"time"
)

type book struct {
	ID       string `json:"id" bson:"_id, omitempty"`
	Title    string `json:"title" bson:"title"`
	Author   string `json:"author" bson:"author"`
	Quantity int    `json:"quantity" bson:"quantity"`
}

var booksCollection *mongo.Collection

// connecting to mongoDB
func connectoMongoDb() (*mongo.Client, error) {
	clientoptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.NewClient(clientoptions)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Println("connected to mongodb")
	return client, nil

}

func getBooks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := booksCollection.Find(ctx, bson.M{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "cannot retrive the data"})
		return
	}
	var books []book
	if err = cursor.All(ctx, &books); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "cannot retrive the data"})
		return
	}
	c.IndentedJSON(http.StatusOK, books)

}
func main() {
	client, err := connectoMongoDb()
	if err != nil {
		panic(err)
	}
	booksCollection = client.Database("Library").Collection("books")
	router := gin.Default()
	router.GET("/", getBooks)
	router.Run("localhost:8082")
}
