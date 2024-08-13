package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Release struct {
	AU string `json:"au"`
	EU string `json:"eu"`
	JP string `json:"jp"`
	NA string `json:"na"`
}

type Amiibo struct {
	AmiiboSeries string  `json:"amiiboSeries"`
	Character    string  `json:"character"`
	GameSeries   string  `json:"gameSeries"`
	Head         string  `json:"head"`
	Image        string  `json:"image"`
	Name         string  `json:"name"`
	Release      Release `json:"release"`
	Tail         string  `json:"tail"`
	Type         string  `json:"type"`
}

type Response struct {
	Amiibo []Amiibo `json:"amiibo"`
}

func GetData() []Amiibo {
	url := "https://www.amiiboapi.com/api/amiibo/"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("HTTP GET request failed: %s\n", err)
		return []Amiibo{}
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %s\n", err)
		return []Amiibo{}
	}

	var data Response
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		fmt.Printf("Failed to unmarshal response body: %s\n", err)
		return []Amiibo{}
	}

	// Filter out only the amiibo figures, not cards
	var amiiboFigures []Amiibo
	for _, item := range data.Amiibo {
		if item.Type == "Figure" {
			amiiboFigures = append(amiiboFigures, item)
		}
	}

	return amiiboFigures
}
