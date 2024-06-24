package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	t "github.com/WatchJani/dinamic-router.git/trie"
)

type handler func()

func main() {
	trie := t.NewTrie[handler]()

	allEndpoints := []struct {
		key   string
		value handler
	}{
		{
			key: "https://chatgpt.com/c/2c9dc1cd-5b43-4f2c-bcfb-22972fb7f52a",
			value: func() {
				fmt.Println("yessss")
			},
		}, {
			key: "https://www.youtube.com/watch?v=-Ung58PHgus",
			value: func() {
				fmt.Println("#")
			},
		}, {
			key: "https://www.youtube.com/watch?v=TFkOzuUz6Wo&t=145s",
			value: func() {
				fmt.Println("mneeeeeeeeeee")
			},
		}, {
			key: "janko/kondic",
			value: func() {
				fmt.Println("yes")
			},
		}, {
			key: "janko/first",
			value: func() {
				fmt.Println("neah")
			},
		},
		{
			key: "dree/{user}/pre",
			value: func() {
				fmt.Println("NEAH")
			},
		},
	}

	for _, endpoint := range allEndpoints {
		key := strings.Split(endpoint.key, "/")
		trie.Insert(key, endpoint.value)
	}

	for {
		value, err := trie.Search(strings.Split(allEndpoints[rand.Intn(5)].key, "/"))
		if err != nil {
			panic(err)
		}

		value()

		time.Sleep(time.Second)
	}
}
