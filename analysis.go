package main

import (
	"bufio"
	"os"
)

func GetWords(path string) (wordsMap map[string] int , err error){

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

		wordsMap[string(line)] = 0
	}
	return
}