package client

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type DataPost struct {
	Post []struct {
		ID    int      `json:"id"`
		Title string   `json:"title"`
		Body  string   `json:"body"`
		Tags  []string `json:"tags"`
		Views int      `json:"views"`
	} `json:"posts"`
}

func GetPost() (*DataPost, error) {
	resp, err := http.Get(os.Getenv("DATA_URI"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	var data DataPost
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &data, nil
}
