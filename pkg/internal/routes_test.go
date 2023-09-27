package internal

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rng = rand.New(rand.NewSource(42))

func TestGetChapterAndVerseIDs(t *testing.T) {
	tests := []struct {
		name            string
		k               int
		expectedChapter int
		expectedVerse   int
	}{
		{
			name:            "Chapter 1, Verse 1",
			k:               0,
			expectedChapter: 1,
			expectedVerse:   1,
		},
		{
			name:            "Chapter 1, Verse 2",
			k:               1,
			expectedChapter: 1,
			expectedVerse:   2,
		},
		{
			name:            "Chapter 1, Verse chapter_1_len",
			k:               6,
			expectedChapter: 1,
			expectedVerse:   7,
		},
		{
			name:            "Chapter 2, Verse 1",
			k:               7,
			expectedChapter: 2,
			expectedVerse:   1,
		},
		{
			name:            "Chapter 2, Verse 2",
			k:               8,
			expectedChapter: 2,
			expectedVerse:   2,
		},
		{
			name:            "Chapter n_chapters, Verse (chapter_n_len-1)",
			k:               6234,
			expectedChapter: 114,
			expectedVerse:   5,
		},
		{
			name:            "Chapter n_chapters, Verse chapter_n_len",
			k:               6235,
			expectedChapter: 114,
			expectedVerse:   6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			chapterID, verseID := getChapterAndVerseIDs(test.k)
			assert.Equal(t, test.expectedChapter, chapterID)
			assert.Equal(t, test.expectedVerse, verseID)
		})
	}
}

func TestGetRandomChapterAndVerseIDs(t *testing.T) {
	chapterID, verseID := GetRandomChapterAndVerseIDs(rng)
	assert.True(t, 0 < chapterID && chapterID <= nChapters)
	assert.True(t, 0 < verseID && verseID <= nVerses)
}

func TestGetVerse(t *testing.T) {
	tests := []struct {
		name                string
		chapter             int
		verse               int
		textType            TextType
		translation         Translation
		expectedText        string
		expectedTranslation string
		expectedError       bool
	}{
		{
			name:          "Invalid chapter",
			chapter:       0,
			verse:         1,
			expectedError: true,
		},
		{
			name:          "Invalid verse",
			chapter:       1,
			verse:         0,
			expectedError: true,
		},
		{
			name:          "Invalid chapter (>n_chapters)",
			chapter:       115,
			verse:         1,
			expectedError: true,
		},
		{
			name:          "Invalid verse (>chapter_1_len)",
			chapter:       1,
			verse:         8,
			expectedError: true,
		},
		{
			name:                "Invalid text type",
			chapter:             1,
			verse:               1,
			textType:            NewTextType("invalid"),
			expectedText:        "بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ",
			expectedTranslation: "In the Name of Allah—the Most Compassionate, Most Merciful.",
		},
		{
			name:                "Invalid translation",
			chapter:             1,
			verse:               1,
			translation:         NewTranslation("invalid"),
			expectedText:        "بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ",
			expectedTranslation: "In the Name of Allah—the Most Compassionate, Most Merciful.",
		},
		{
			name:                "Chapter 1, Verse 1",
			chapter:             1,
			verse:               1,
			expectedText:        "بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ",
			expectedTranslation: "In the Name of Allah—the Most Compassionate, Most Merciful.",
		},
		{
			name:                "Chapter 114, Verse 6",
			chapter:             114,
			verse:               6,
			expectedText:        "مِنَ الْجِنَّةِ وَالنَّاسِ",
			expectedTranslation: "from among jinn and humankind.”",
		},
		{
			name:                "Chapter 1, Verse 1 (non-default text type)",
			chapter:             1,
			verse:               1,
			textType:            Indopak,
			expectedText:        "بِسۡمِ اللهِ الرَّحۡمٰنِ الرَّحِيۡمِ",
			expectedTranslation: "In the Name of Allah—the Most Compassionate, Most Merciful.",
		},
		{
			name:                "Chapter 1, Verse 1 (non-default translation)",
			chapter:             1,
			verse:               1,
			translation:         YusufAli,
			expectedText:        "بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ",
			expectedTranslation: "In the name of Allah, Most Gracious, Most Merciful.",
		},
		{
			name:                "Chapter 1, Verse 1 (non-default text type and translation)",
			chapter:             1,
			verse:               1,
			textType:            Indopak,
			translation:         YusufAli,
			expectedText:        "بِسۡمِ اللهِ الرَّحۡمٰنِ الرَّحِيۡمِ",
			expectedTranslation: "In the name of Allah, Most Gracious, Most Merciful.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			verse, err := GetVerse(test.chapter, test.verse, test.textType, test.translation)
			if test.expectedError {
				assert.Error(t, err)
				assert.Equal(t, Verse{}, verse)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedText, verse.Text)
				assert.Equal(t, test.expectedTranslation, verse.Translation)
			}
		})
	}
}
