package appcontent

import (
	"github.com/CaSe-SuperCashback/go-utilities/language"
	"github.com/goccy/go-json"
)

type ActionType struct {
	Type  string                `json:"type"`
	Value string                `json:"value"`
	Text  language.Multilingual `json:"text"`
}

func (at ActionType) ToString() (string, error) {
	if data, err := json.Marshal(at); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}
