package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Aadil101/ayah-backend/pkg/internal"
)

func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", ping)
	mux.HandleFunc("GET /verse/random", getRandomVerse)
	mux.HandleFunc("GET /verse/by_key/{chapterID}/{verseID}", getVerse)
	mux.HandleFunc("GET /resources/textTypes", getTextTypes)
	mux.HandleFunc("GET /resources/translations", getTranslations)
	return mux
}

func ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "pong",
	})
}

func getChapterVerse(w http.ResponseWriter, r *http.Request, chapterID int, verseID int) {
	textType := internal.NewTextType(r.URL.Query().Get("textType"))
	translation := internal.NewTranslation(r.URL.Query().Get("translation"))
	verse, err := internal.GetVerse(chapterID, verseID, textType, translation)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
	} else {
		json.NewEncoder(w).Encode(verse)
	}
}

func getRandomVerse(w http.ResponseWriter, r *http.Request) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	chapterID, verseID := internal.GetRandomChapterAndVerseIDs(rng)
	getChapterVerse(w, r, chapterID, verseID)
}

func getVerse(w http.ResponseWriter, r *http.Request) {
	chapterID, _ := strconv.Atoi(r.PathValue("chapterID"))
	verseID, _ := strconv.Atoi(r.PathValue("verseID"))
	getChapterVerse(w, r, chapterID, verseID)
}

func getTextTypes(w http.ResponseWriter, r *http.Request) {
	var ttr internal.TextTypesResponse
	for _, tt := range internal.GetTextTypes() {
		ttr.TextTypeResponses = append(ttr.TextTypeResponses, internal.TextTypeResponse{ID: tt.GetID(), Name: tt.GetName()})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ttr)
}

func getTranslations(w http.ResponseWriter, r *http.Request) {
	var tr internal.TranslationsResponse
	for _, tt := range internal.GetTranslations() {
		tr.TranslationResponses = append(tr.TranslationResponses, internal.TranslationResponse{ID: tt.GetID(), Name: tt.GetName()})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tr)
}
