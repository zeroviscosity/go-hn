package hn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetTopIds() ([]int, error) {
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var topIds []int
	err = json.Unmarshal(body, &topIds)
	if err != nil {
		return nil, err
	}

	return topIds, nil
}
