package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/thinkerou/favicon"

	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()

	return fmt.Sprintf("%02d-%02d-%02d", day, month, year)
}

func main() {
	r := gin.Default()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	r.Use(favicon.New("./assets/underscore.png"))

	r.LoadHTMLFiles("./templates/index.tmpl")

	r.Static("/static", "./assets")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"now": formatAsDate(time.Now()),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
