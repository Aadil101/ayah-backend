package handler

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Aadil101/ayah-backend/pkg/internal"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getVerse(c *gin.Context, chapterID int, verseID int) {
	textType := internal.NewTextType(c.DefaultQuery("textType", ""))
	translation := internal.NewTranslation(c.DefaultQuery("translation", ""))

	verse, err := internal.GetVerse(chapterID, verseID, textType, translation)
	if err != nil {
		//log.Fatal("Error while retrieving verse from Quran.com API: ", err)
		return
	}

	c.IndentedJSON(http.StatusOK, verse)
}

func GetRandomVerse(c *gin.Context) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	chapterID, verseID := internal.GetRandomChapterAndVerseIDs(rng)
	getVerse(c, chapterID, verseID)
}

func GetVerse(c *gin.Context) {
	chapterID, _ := strconv.Atoi(c.Param("chapterID"))
	verseID, _ := strconv.Atoi(c.Param("verseID"))
	getVerse(c, chapterID, verseID)
}

func GetTextTypes(c *gin.Context) {
	var ttr internal.TextTypesResponse
	for _, tt := range internal.GetTextTypes() {
		ttr.TextTypeResponses = append(ttr.TextTypeResponses, internal.TextTypeResponse{ID: tt.GetID(), Name: tt.GetName()})
	}
	c.IndentedJSON(http.StatusOK, ttr)
}

func GetTranslations(c *gin.Context) {
	var tr internal.TranslationsResponse
	for _, tt := range internal.GetTranslations() {
		tr.TranslationResponses = append(tr.TranslationResponses, internal.TranslationResponse{ID: tt.GetID(), Name: tt.GetName()})
	}
	c.IndentedJSON(http.StatusOK, tr)
}
