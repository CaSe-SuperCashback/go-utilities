package language

import "github.com/goccy/go-json"

type Multilingual struct {
	English    string `json:"en,omitempty"`
	Vietnamese string `json:"vi,omitempty"` // Vietnamese
}

func (m Multilingual) GetLocalized(lang string) Multilingual {
	// English by default
	result := Multilingual{
		English: m.English,
	}

	switch lang {
	case Vietnamese.String():
		result.Vietnamese = m.Vietnamese
	}

	return result
}

func (m Multilingual) IsEmpty() bool {
	return m.English == "" && m.Vietnamese == ""
}

func (m Multilingual) ToString() (string, error) {
	if data, err := json.Marshal(m); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}
