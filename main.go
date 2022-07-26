package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type entity struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Email string `json:"email"`
}

var entities = []entity {
	{Id: "1", Name: "Nuntapong Siripanyawong", Gender: "Male", Email: "nuntapong@odds.team"},
	{Id: "2", Name: "Kanyanat Nomsian", Gender: "Female", Email: "bow@odds.team"},

}

func main() {

	router := gin.Default()

	router.GET("/entities", getEntities)
	router.GET("/entities/:id", getEntitiesById)
	router.POST("/entities", postEntities)
	router.DELETE("/entities/:id", deleteEntities)

	router.Run()
}

func getEntities(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, entities)
}

func getEntitiesById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range entities {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "entities not found"})
}

func postEntities(c *gin.Context) {
	var newEntity entity

	if err := c.BindJSON(&newEntity); err != nil {
		return
	}

	entities = append(entities, newEntity)
	c.IndentedJSON(http.StatusOK, newEntity)
}

func deleteEntities(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i <= len(entities)-1; i++ {
		if entities[i].Id == id {
			entities = append(entities[:i], entities[i+1:]... )
			c.IndentedJSON(http.StatusOK, gin.H{"message": "delete entities successfully"})
			return
		}
	}
}