package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Student struct {
	ID      string `json:"name"`
	Name    string `json:"s_name"`
	Age     int    `json:"age"`
	Class   string `json:"class"`
	Section string `json:"section"`
}

var students = []Student{
	{ID: "1ds23is023", Name: "Arpit_Srivastava", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is021", Name: "Anurag_Jain", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is022", Name: "Ananaya_Gowda", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is020", Name: "Khushi_Agrawal", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is024", Name: "Harsh_Gupta", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is025", Name: "Aaskar_Rana", Age: 20, Class: "3", Section: "A"},
}

func GetAllStudentsNames(c *gin.Context){
	c.IndentedJSON(http.StatusOK,students)
}

func main(){
	router:= gin.Default()
	router.GET("/",GetAllStudentsNames)
	router.Run("localhost:8083")
}

