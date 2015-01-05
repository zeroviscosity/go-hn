package hn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Item struct {
	Id      int
	Deleted bool
	Type    string
	By      string
	Time    int
	Text    string
	Dead    bool
	Parent  int
	Kids    []int
	Url     string
	Score   int
	Title   string
	Parts   []int
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
