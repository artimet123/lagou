package utils

import (
	"testing"
)

func TestExistPositions(t *testing.T) {
	exists, err := ExistPositions("./testdata", ".txt")

	if err != nil{
		t.Fatal(err)
	}

	if i := len(exists); i != 2 {
		t.Fatalf("want 2, get %d", i)
	}

	t.Log(exists)
}

func TestRankByWordCount(t *testing.T) {
	wordsMap := make(map[string]int)

	wordsMap["a"] = 12
	wordsMap["b"] = 22
	wordsMap["c"] = 32

	result := [][]string {
		[]string {"c", "32"},
		[]string {"b", "22"},
		[]string {"a", "12"},
	}

	ranks := RankByWordCount(wordsMap)
	t.Log(ranks)
	for i, words := range ranks {
		if words[0] != result[i][0] || words[1] != result[i][1] {
			t.Fatal("rank error!")
		}
	}
}