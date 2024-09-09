package appcontent

import "github.com/goccy/go-json"

type ActionType struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Text  string `json:"text"`
}

func (at ActionType) ToString() (string, error) {
	if data, err := json.Marshal(at); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}
