package main

import (
	"bufio"
	"os"
	"strings"
)

func GetWordsMap(path string) (wordsMap map[string] int , err error){

	wordsMap = make(map[string]int)
	file, err := os.Open(path)

	if err != nil {
		return
	}
	bufReader := bufio.NewReader(file)
	for {

		var line []byte
		line, _, err =bufReader.ReadLine()

		if err != nil {
			err = nil
			break
		}

		word :=  strings.Split(string(line), "\t")[0]
		word = strings.Trim(word, " ")
		wordsMap[word] = 0
	}
	return
}