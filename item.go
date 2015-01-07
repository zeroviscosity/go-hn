package hn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Item struct {
	Id      int    `json:"id"`
	Deleted bool   `json:"deleted"`
	Type    string `json:"type"`
	By      string `json:"by"`
	Time    int    `json:"time"`
	Text    string `json:"text"`
	Dead    bool   `json:"dead"`
	Parent  int    `json:"parent"`
	Kids    []int  `json:"kids"`
	Url     string `json:"url"`
	Score   int    `json:"score"`
	Title   string `json:"title`
	Parts   []int  `json:"parts"`
}

func GetItem(id int) (*Item, error) {
	url := "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(id) + ".json"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var item Item
	err = json.Unmarshal(body, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
