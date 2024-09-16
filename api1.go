package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)
type Movies struct {
	ID         string `json:"id" bson:"_id, omitempty"`
	Name       string `json:"name" bson:"name"`
	Release    int    `json:"date" bson: "Date"`
	Collection string `json:"money" bson: "collection"`
}
var moviesCollection *mongo.Collection

//connect to mongodb

func connectToMongo()(*mongo.Client, error){
	clientoptions:= options.Client().ApplyURI("")
}

var movies = []Movies{

	
}

func GetMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}
func main() {
	router := gin.Default()
	router.GET("/", GetMovies)
	router.Run("localhost:8081")
}
