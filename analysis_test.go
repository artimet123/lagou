package main

import "testing"

func TestGetWords(t *testing.T) {
	testData := []struct{
		path string
		count int
		words []string
	} {
		{
			"testdata/words.txt",
			5,
			[]string{},
		},
	}

	for _, test := range testData {
		wordsMap, err := GetWords(test.path)

		if err != nil {
			t.Fatal(err)
		}

		if count := len(wordsMap); count != test.count {
			t.Fatalf("want %d, get %d", count, test.count)
		}
	}


}