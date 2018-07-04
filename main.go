package main

import (
	"github.com/shawpo/lagouWordCloud/spider"
	"net/url"
	"time"
	"math/rand"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
	"io/ioutil"
	"bytes"
)
func main() {
	i := 1
	sum := 1
	kd, first := "java", "true"
	positionsMap := make(map[int]string)
	for {
		if i > 1 {
			first = "false"
			timeSleep(1, 2)
		}
		values:= url.Values{
			"first": {first},
			"pn" : {strconv.Itoa(i)},
			"kd": {kd},
		}
		request, err := spider.KdPositionsRequest(values)
		if err != nil {
			log.Print(err)
			continue
		}
		resp, err := spider.GetClient().Do(request)
		if err != nil {
			log.Print(err)
			continue
		}
		positions ,err := spider.GetPositions(resp.Body)

		if err != nil{
			log.Print(err)
			continue
		}

		if len(positions) == 0 {
			break
		}

		for _, position := range positions {
			log.Printf("No.%d   %d: %s",sum, position.PositionID, position.PositionName)
			positionsMap[position.PositionID] = position.PositionName
			sum++
		}
		resp.Body.Close()
		i++
	}

	filePath := kd + ".txt"

	file, err := os.Create(filePath)

	if err != nil {
		 log.Fatalf("can't create file %s to save data", filePath)
	}

	if len(positionsMap) == 0 {
		log.Fatal("get none positions")
	}
	i = 1
	for positionId := range positionsMap {
		timeSleep(1,2)
		request, err := spider.DetailRequest(positionId)
		if err != nil {
			log.Print(err)
			continue
		}
		client := spider.GetClient()
		resp, err := client.Do(request)

		html, err := ioutil.ReadAll(resp.Body)

		if err !=nil {
			log.Print(err)
			continue
		}
		if resp.StatusCode != 200{
			log.Print("request error")
		}
		detail ,err := spider.GetDetail(bytes.NewReader(html))

		if detail == "" {

			log.Printf("positionId %d: get detail none", positionId)
			log.Fatal(string(html))

		}
		if err != nil{
			log.Print(err)
			continue
		}
		log.Printf("write to file: %d of %d", i, len(positionsMap))
		_, err = file.WriteString(detail)
		if err != nil {
			log.Print(err)
			continue
		}


		resp.Body.Close()
		i++
	}

}

func timeSleep(base int, random int)  {
	rand.Seed(time.Now().UnixNano())
	sleep := base + rand.Intn(random)
	time.Sleep(time.Second * time.Duration(sleep))
}
