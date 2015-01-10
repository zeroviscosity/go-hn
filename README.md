# Go Package for the Hacker News API v0

## Usage

Add `"github.com/zeroviscosity/go-hn"` to your imports. You can then access the following three methods:

#### Get the Top 100 Stories

```go
topIds, err := hn.GetTopIds()
```

This returns an array of the integer ids of the top 100 stories as described in the [https://github.com/HackerNews/API#top-stories](API docs).

#### Get an Item

```go
item, err := hn.GetItem(id)
```

This takes an interger id and returns an `Item` struct corresponding to the JSON object returned by the API. As described in the [https://github.com/HackerNews/API#items](docs), its `type` field will be one of the following: 

* "job"
* "story"
* "comment"
* "poll"
* "pollopt"

Depending on the `type`, some fields will be empty.

```go
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
```

#### Get a User

```go
user, err := hn.GetUser(item.By)
```

This takes a string id (ie: username) and returns a `User` struct with fields matching those provided by the API. See the [https://github.com/HackerNews/API#users](docs) for more details.

```go
type User struct {
	Id        string `json:"id"`
	Delay     int    `json:"delay"`
	Created   int    `json:"created"`
	Karma     int    `json:"karma"`
	About     string `json:"about"`
	Submitted []int  `json:"submitted"`
}
```

## Example

```go
package main

import (
	"fmt"

	"github.com/zeroviscosity/go-hn"
)

func displayItem(rank int, id int) {
	// Get a single item by its integer id
	item, err := hn.GetItem(id)
	if err != nil {
		fmt.Printf("#%v ERROR: %v\n\n", rank, err)
		return
	}

	// Get a single user by his/her string id
	user, err := hn.GetUser(item.By)
	if err != nil {
		fmt.Printf("#%v ERROR: %v\n\n", rank, err)
		return
	}

	fmt.Printf("#%v \"%v\" by %v (karma: %v)\n\n", rank, item.Title, item.By, user.Karma)
}

func main() {
	// Get the top 100 item integer ids
	topIds, err := hn.GetTopIds()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the top ten
	for i := 0; i < 10; i++ {
		displayItem(i+1, topIds[i])
	}
}
```
