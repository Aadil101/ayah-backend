package main

import (
	"github.com/Aadil101/ayah-backend/pkg/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", handler.Ping)
	router.GET("/verse/random", handler.GetRandomVerse)
	router.GET("/verse/by_key/:chapterID/:verseID", handler.GetVerse)
	router.GET("/resources/textTypes", handler.GetTextTypes)
	router.GET("/resources/translations", handler.GetTranslations)
	router.Run() // listen and serve on 0.0.0.0:8080
}
