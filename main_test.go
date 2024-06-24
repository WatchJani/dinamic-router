package main

import (
	"strings"
	"testing"
)

func BenchmarkSpeedTest(b *testing.B) {
	b.StopTimer()

	trie := NewTrie[handler]()

	allEndpoints := []struct {
		key   string
		value handler
	}{
		{
			key: "dree/{*}",
			value: func() {
				// fmt.Println("NEAH")
			},
		},
	}

	for _, endpoint := range allEndpoints {
		key := strings.Split(endpoint.key, "/")
		trie.Insert(key, endpoint.value)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		value, err := trie.Search([]string{"dree", "janko"})
		if err != nil {
			panic(err)
		}

		value()
	}
}
