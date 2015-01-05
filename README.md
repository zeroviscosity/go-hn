# Go Package for the Hacker News API

See the API docs for details: https://github.com/HackerNews/API  

## Usage

### Get the ids of the top stories

```
topIds, err := hn.GetTopIds()
```

This returns an array of integer item ids.

### Get a single item

```
item, err := hn.GetItem(id) // `id` is an int
```

This returns an Item:

```
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
```

### Get a single user

```
user, err := hn.GetUser(id) // `id` is a string
```

This retuns a User:

```
type User struct {
	Id        string
	Delay     int
	Created   int
	Karma     int
	About     string
	Submitted []int
}
```
