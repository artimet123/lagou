package main

import "testing"

func TestLoadDictionary(t *testing.T) {
	testData := []struct{
		path string
		count int
		words []string
	} {
		{
			"testdata/words.txt",
			6,
			[]string{"字符串", "初始化", "数组", "加载", "客户端", "JAVA"},
		},
	}


	for _, test := range testData {

		var wordsMap = make(WordsMap)
		err := LoadDictionary(test.path, &wordsMap)
		t.Log(wordsMap)

		if err != nil {
			t.Fatal(err)
		}

		if count := len(wordsMap); count != test.count {
			t.Fatalf("want %d, get %d", test.count, count)
		}

		for _, testWord := range test.words{
			if _, ok := (wordsMap)[testWord]; !ok {
				t.Fatalf("want %s, but don't contains", testWord)
			}

		}
	}
}

func TestAddFileEnWords(t *testing.T) {
	testData := []struct{
		path string
		count int
		words []string
	} {
		{
			"testdata/file.txt",
			7,
			[]string{"LINUX", "TCP", "IP", "HTTP", "GO", "C", "PYTHON"},
		},
	}

	for _, test := range testData {

		var wordsMap = make(WordsMap)
		err := AddFileEnWords(test.path, &wordsMap)
		t.Log(wordsMap)

		if err != nil {
			t.Fatal(err)
		}

		if count := len(wordsMap); count != test.count {
			t.Fatalf("want %d, get %d", test.count, count)
		}

		for _, testWord := range test.words{
			if _, ok := (wordsMap)[testWord]; !ok {
				t.Fatalf("want %s, but don't contains", testWord)
			}

		}
	}
}

