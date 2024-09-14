package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movies struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Release int `json:"date"`
	Collection string `json:"money"`
}

var movies = []Movies{
	
		{ID: "123", Name:"Intersteller", Release: 2014, Collection: "1.23 Crores"},
		{ID: "673", Name:"Shutter Island", Release: 2010, Collection: "186.4 Crores"},
		{ID: "221", Name:"Tenet", Release: 2020, Collection: "20.4 Crores"},
		{ID: "829", Name:"The Prestige", Release: 2006, Collection: "185 Millions"},
		{ID: "234", Name:"Memento", Release: 2000, Collection: "9 Millions"},
		{ID: "454", Name:"Dunkirk", Release: 2017, Collection: "530 Millions"},
	
}

func GetMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}
func main() {
	router:= gin.Default()
	router.GET("/", GetMovies)
	router.Run("localhost:8081")
}
