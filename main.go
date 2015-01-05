package main

import (
	"fmt"
	"hn"
)

func displayItem(rank int, id int) {
	item, err := hn.GetItem(id)
	if err != nil {
		fmt.Printf("#%v ERROR: %v", rank, err)
		return
	}

	user, err := hn.GetUser(item.By)
	if err != nil {
		fmt.Printf("#%v ERROR: %v", rank, err)
		return
	}

	fmt.Printf("#%v \"%v\" by %v (karma: %v)\n\n", rank, item.Title, item.By, user.Karma)
}

func main() {
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
