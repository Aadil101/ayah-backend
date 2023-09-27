package internal

// TODO: Pick which text types to support
type TextType int64

const (
	Imlaei TextType = iota
	Indopak
	Uthmani
)

type backendTextType struct {
	ID         string
	Name       string
	ResourceID string
}

var backendTextTypes = []backendTextType{
	{ID: "Imlaei", Name: "Imlaei", ResourceID: "text_imlaei"},
	{ID: "Indopak", Name: "Indopak", ResourceID: "text_indopak"},
	{ID: "Uthmani", Name: "Uthmani", ResourceID: "text_uthmani"},
}

func NewTextType(id string) TextType {
	for idx, btt := range backendTextTypes {
		if btt.ID == id {
			return TextType(idx)
		}
	}
	return Imlaei
}

func (t TextType) getBackendTextType() backendTextType {
	idx := int(t)
	if idx < 0 || idx >= len(backendTextTypes) {
		idx = 0
	}
	return backendTextTypes[idx]
}

func (t TextType) GetID() string {
	return t.getBackendTextType().ID
}

func (t TextType) GetName() string {
	return t.getBackendTextType().Name
}

func (t TextType) GetResourceID() string {
	return t.getBackendTextType().ResourceID
}

func GetTextTypes() []TextType {
	var result []TextType
	for i := 0; i < len(backendTextTypes); i++ {
		result = append(result, TextType(i))
	}
	return result
}

type TextTypeResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TextTypesResponse struct {
	TextTypeResponses []TextTypeResponse `json:"textTypes"`
}

// TODO: Pick which translations to support
type Translation int64

const (
	MustafaKhattab Translation = iota
	YusufAli
	AbdulHaleem
)

type backendTranslation struct {
	ID         string
	Name       string
	ResourceID int
}

var backendTranslations = []backendTranslation{
	{ID: "MustafaKhattab", Name: "Mustafa Khattab", ResourceID: 131},
	{ID: "YusufAli", Name: "Abdullah Yusuf Ali", ResourceID: 22},
	{ID: "AbdulHaleem", Name: "M.A.S. Abdel Haleem", ResourceID: 85},
}

func NewTranslation(id string) Translation {
	for idx, bt := range backendTranslations {
		if bt.ID == id {
			return Translation(idx)
		}
	}
	return MustafaKhattab
}

func (t Translation) getBackendTranslation() backendTranslation {
	idx := int(t)
	if idx < 0 || idx >= len(backendTranslations) {
		idx = 0
	}
	return backendTranslations[idx]
}

func (t Translation) GetID() string {
	return t.getBackendTranslation().ID
}

func (t Translation) GetName() string {
	return t.getBackendTranslation().Name
}

func (t Translation) GetResourceID() int {
	return t.getBackendTranslation().ResourceID
}

func GetTranslations() []Translation {
	var result []Translation
	for i := 0; i < len(backendTranslations); i++ {
		result = append(result, Translation(i))
	}
	return result
}

type TranslationResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TranslationsResponse struct {
	TranslationResponses []TranslationResponse `json:"translations"`
}

type Verse struct {
	Key         string `json:"key"`
	Text        string `json:"text"`
	Translation string `json:"translation"`
}

type backendVerse struct {
	ID                 int    `json:"id"`
	VerseNumber        int    `json:"verse_number"`
	VerseKey           string `json:"verse_key"`
	TextUthmani        string `json:"text_uthmani"`
	TextUthmaniSimple  string `json:"text_uthmani_simple"`
	TextImlaei         string `json:"text_imlaei"`
	TextImlaeiSimple   string `json:"text_imlaei_simple"`
	TextIndopak        string `json:"text_indopak"`
	TextUthmaniTajweed string `json:"text_uthmani_tajweed"`
	Translations       []struct {
		ID         int    `json:"id"`
		ResourceID int    `json:"resource_id"`
		Text       string `json:"text"`
	} `json:"translations"`
}

// TODO: Make unit tests
func (v backendVerse) getText(textType TextType) string {
	propertyMap := map[string]string{
		"text_uthmani": v.TextUthmani,
		// "text_uthmani_simple":  v.TextUthmaniSimple,
		"text_imlaei": v.TextImlaei,
		// "text_imlaei_simple":   v.TextImlaeiSimple,
		"text_indopak": v.TextIndopak,
		// "text_uthmani_tajweed": v.TextUthmaniTajweed,
	}

	property, ok := propertyMap[textType.GetResourceID()]
	if !ok {
		// Handle the case where textType is not in the map
		return ""
	}

	return property
}

// TODO: Make unit tests
func (v backendVerse) getTranslation(translation Translation) string {
	for _, t := range v.Translations {
		if t.ResourceID == translation.GetResourceID() {
			return t.Text
		}
	}
	return ""
}
