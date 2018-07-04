package main

import "testing"

func TestGetWordsMap(t *testing.T) {
	testData := []struct{
		path string
		count int
		words []string
	} {
		{
			"testdata/words.txt",
			5,
			[]string{"字符串", "初始化", "数组", "加载", "客户端"},
		},
	}


	for _, test := range testData {
		wordsMap, err := GetWordsMap(test.path)
		t.Log(wordsMap)

		if err != nil {
			t.Fatal(err)
		}

		if count := len(wordsMap); count != test.count {
			t.Fatalf("want %d, get %d", count, test.count)
		}

		for _, testWord := range test.words{
			if _, ok := wordsMap[testWord]; !ok {
				t.Fatalf("want %s, bug don't contains", testWord)
			}

		}
	}


}