// main.go
package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	router := gin.Default()

	// Serve HTML page with the spinwheel
	router.GET("/", func(c *gin.Context) {
		renderTemplate(c, "index.html", nil)
	})

	// Start the server
	router.Run(":8080")
}

func renderTemplate(c *gin.Context, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(c.Writer, tmpl, data)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
