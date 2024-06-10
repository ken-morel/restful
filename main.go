package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Time struct {
	Pm     bool `json: "pm"`
	Hour   int  `json: "hour"`
	Minute int  `json: "minute"`
	Second int  `json: "second"`
}

type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Time  Time   `json:"time"`
}

var notes = []Note{
	{
		ID:    1,
		Title: "goto bed",
		Time: Time{
			Pm:     true,
			Hour:   9,
			Minute: 0,
			Second: 0,
		},
	},
}

func main() {
	router := gin.Default()
	router.GET("/", getNotes)
	router.POST("/", createNotes)

	router.Run("localhost:80")
}
func getNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, notes)
}

func createNote(c *gin.Context) {
	var newNote Note

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newNote); err != nil {
		return
	}

	// Add the new album to the slice.
	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}

func getNoteByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
