package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Node[T any] struct {
	value T
	next  map[string]*Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		value: value,
		next:  make(map[string]*Node[T]),
	}
}

type Trie[T any] struct {
	root *Node[T]
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{
		root: &Node[T]{
			next: make(map[string]*Node[T]),
		},
	}
}

func (t *Trie[T]) Insert(keys []string, value T) *Node[T] {
	current := t.root

	for _, key := range keys {
		if len(key) > 1 && key[0] == '{' && key[len(key)-1] == '}' {
			key = "*"
		}

		if _, exist := current.next[key]; !exist {
			current.next[key] = NewNode[T](value)
		}
		current = current.next[key]
	}

	return current
}

func (t *Trie[T]) Search(keys []string) (T, error) {
	current := t.root

	for _, key := range keys {
		if next, exist := current.next[key]; exist {
			current = next
		} else if next, exist := current.next["*"]; exist {
			current = next
		} else {
			var zeroValue T
			return zeroValue, errors.New("key does not exist")
		}
	}

	return current.value, nil
}

type handler func()

func main() {
	trie := NewTrie[handler]()

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
		}, {
			key: "dree/{*}",
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
