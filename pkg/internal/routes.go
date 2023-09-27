package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sort"
)

const (
	quranAPIBaseURL = "https://api.quran.com/api/v4"
	nVerses         = 6236
	nChapters       = 114
)

func getChapterAndVerseIDs(k int) (int, int) {
	if k >= nVerses {
		log.Fatal("k is too large: ", k)
	}

	chapterID := sort.Search(len(verseIDCumSum), func(i int) bool { return verseIDCumSum[i] > k })

	if chapterID == 0 {
		return chapterID + 1, k + 1
	}

	return chapterID + 1, k - verseIDCumSum[chapterID-1] + 1
}

func GetRandomChapterAndVerseIDs(rng *rand.Rand) (int, int) {
	return getChapterAndVerseIDs(rng.Intn(nVerses))
}

func GetVerse(chapterID int, verseID int, textType TextType, translation Translation) (Verse, error) {
	url := fmt.Sprintf("%s/verses/by_key/%d:%d?fields=%s&translations=%d", quranAPIBaseURL, chapterID, verseID, textType.GetResourceID(), translation.GetResourceID())
	resp, err := http.Get(url)
	if err != nil {
		return Verse{}, fmt.Errorf("error while retrieving verse from Quran.com API: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return Verse{}, fmt.Errorf("error while retrieving verse from Quran.com API: %s", string(b))
	}
	defer resp.Body.Close()

	var temp struct {
		Verse backendVerse `json:"verse"`
	}
	err = json.NewDecoder(resp.Body).Decode(&temp)

	result := Verse{
		Key:         temp.Verse.VerseKey,
		Text:        temp.Verse.getText(textType),
		Translation: temp.Verse.getTranslation(translation),
	}

	return result, err
}
