package main

import (
	"bufio"
	"os"
	"strings"
	"io/ioutil"
	"regexp"
)

type WordsMap map[string] int

func LoadDictionary(path string, wordsMap *WordsMap) (err error){
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
		word = strings.ToUpper(word)
		(*wordsMap)[word] = 0
	}
	return
}

func AddFileEnWords(path string, wordsMap *WordsMap) (err error) {
	file, err := os.Open(path)

	if err != nil {
		return
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	re := regexp.MustCompile(`\b[a-zA-Z]+\b`)

	enWords := re.FindAllString(string(contents), -1)
	for _, enWord := range enWords{
		enWord = strings.ToUpper(enWord)
		(*wordsMap)[enWord] = 0
	}
	return
}