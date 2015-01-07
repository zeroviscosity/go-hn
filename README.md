# Go Package for the Hacker News API

See the API docs for details on the JSON returned: https://github.com/HackerNews/API  

## Example

```
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
