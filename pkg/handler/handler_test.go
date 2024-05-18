package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Aadil101/ayah-backend/pkg/internal"

	"github.com/stretchr/testify/assert"
)

var handler = NewHandler()

func TestPingHandler(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/ping", nil)
	handler.ServeHTTP(writer, request)

	var pingResponse struct {
		Message string `json:"message"`
	}
	json.Unmarshal(writer.Body.Bytes(), &pingResponse)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "pong", pingResponse.Message)
}

func TestGetRandomVerseHandler(t *testing.T) {
	writer := httptest.NewRecorder()
	var verse internal.Verse

	request, _ := http.NewRequest("GET", "/verse/random", nil)
	handler.ServeHTTP(writer, request)
	json.Unmarshal(writer.Body.Bytes(), &verse)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEqual(t, internal.Verse{}, verse)
}

func TestGetVerseHandler(t *testing.T) {
	var verse internal.Verse

	request, _ := http.NewRequest("GET", "/verse/by_key/invalid/invalid", nil)
	writer := httptest.NewRecorder()
	handler.ServeHTTP(writer, request)
	json.Unmarshal(writer.Body.Bytes(), &verse)
	assert.Equal(t, http.StatusInternalServerError, writer.Code)

	request, _ = http.NewRequest("GET", "/verse/by_key/1/1", nil)
	writer = httptest.NewRecorder()
	handler.ServeHTTP(writer, request)
	json.Unmarshal(writer.Body.Bytes(), &verse)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, internal.Verse{
		Key:         "1:1",
		Text:        "بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ",
		Translation: "In the Name of Allah—the Most Compassionate, Most Merciful.",
	}, verse)

	request, _ = http.NewRequest("GET", "/verse/by_key/1/1?textType=invalid&translation=invalid", nil)
	writer = httptest.NewRecorder()
	handler.ServeHTTP(writer, request)
	json.Unmarshal(writer.Body.Bytes(), &verse)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, internal.Verse{
		Key:         "1:1",
		Text:        "بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ",
		Translation: "In the Name of Allah—the Most Compassionate, Most Merciful.",
	}, verse)

	request, _ = http.NewRequest("GET", "/verse/by_key/1/1?textType=Indopak&translation=YusufAli", nil)
	writer = httptest.NewRecorder()
	handler.ServeHTTP(writer, request)
	json.Unmarshal(writer.Body.Bytes(), &verse)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, internal.Verse{
		Key:         "1:1",
		Text:        "بِسۡمِ اللهِ الرَّحۡمٰنِ الرَّحِيۡمِ",
		Translation: "In the name of Allah, Most Gracious, Most Merciful.",
	}, verse)
}

func TestGetTextTypes(t *testing.T) {
	var ttr internal.TextTypesResponse

	request, _ := http.NewRequest("GET", "/resources/textTypes", nil)
	writer := httptest.NewRecorder()
	handler.ServeHTTP(writer, request)
	json.Unmarshal(writer.Body.Bytes(), &ttr)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.ElementsMatch(t, []internal.TextTypeResponse{
		{ID: "Imlaei", Name: "Imlaei"},
		{ID: "Indopak", Name: "Indopak"},
		{ID: "Uthmani", Name: "Uthmani"},
	}, ttr.TextTypeResponses)
}

func TestGetTranslations(t *testing.T) {
	var tr internal.TranslationsResponse

	request, _ := http.NewRequest("GET", "/resources/translations", nil)
	writer := httptest.NewRecorder()
	handler.ServeHTTP(writer, request)
	json.Unmarshal(writer.Body.Bytes(), &tr)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.ElementsMatch(t, []internal.TranslationResponse{
		{ID: "MustafaKhattab", Name: "Mustafa Khattab"},
		{ID: "YusufAli", Name: "Abdullah Yusuf Ali"},
		{ID: "AbdulHaleem", Name: "M.A.S. Abdel Haleem"},
	}, tr.TranslationResponses)
}
