package utils

import (
	"math/rand"
	"time"
	"path/filepath"
	"os"
	"fmt"
	"strings"
	"log"
)

func RandTimeSleep(base int, random int)  {
	rand.Seed(time.Now().UnixNano())
	sleep := base + rand.Intn(random)
	time.Sleep(time.Second * time.Duration(sleep))
}

func ExistPositions(path string, ext string) (positions []string, err error ){
	files, err := filepath.Glob(fmt.Sprintf("%s/*%s", path, ext))
	if err != nil {
		return
	}

	for _, file := range files {

		fileInfo, err := os.Stat(fmt.Sprintf("%s", file))

		if err != nil {
			log.Print(err)
			continue
		}

		if !(fileInfo.Size() >0) {
			continue
		}

		positions = append(positions, strings.Replace(fileInfo.Name(), ext, "", 1))
	}

	return
}