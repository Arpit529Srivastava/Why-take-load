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

type Student struct {
	ID      string `json:"name" bson:"_id, omitempty"`
	Name    string `json:"s_name" bson:"Name"`
	Age     int    `json:"age" bson:"Age"`
	Class   string `json:"class" bson:"Class"`
	Section string `json:"section" bson:"Section"`
}

var StudentCollection *mongo.Collection

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

func GetAllStudentsNames(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := StudentCollection.Find(ctx, bson.M{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "cannot retrive the data"})
		return
	}
	var student []Student
	if err = cursor.All(ctx, &student); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "cannot retrive the data"})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func main() {
	client, err := connectoMongoDb()
	if err != nil {
		panic(err)
	}
	StudentCollection = client.Database("school").Collection("students")
	router := gin.Default()
	router.GET("/", GetAllStudentsNames)
	router.Run("localhost:8083")
}
