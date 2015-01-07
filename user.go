package hn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	Id        string `json:"id"`
	Delay     int    `json:"delay"`
	Created   int    `json:"created"`
	Karma     int    `json:"karma"`
	About     string `json:"about"`
	Submitted []int  `json:"submitted"`
}

func GetUser(id string) (*User, error) {
	url := "https://hacker-news.firebaseio.com/v0/user/" + id + ".json"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
