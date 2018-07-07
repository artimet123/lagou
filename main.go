package main

import (
	"github.com/shawpo/lagouWordCloud/spider"
	"net/url"
	"os"
	"strconv"
	"io/ioutil"
	"bytes"
	"github.com/shawpo/sego"
	"log"
	"github.com/shawpo/lagouWordCloud/analysis"
	"github.com/manifoldco/promptui"
	"fmt"
	"github.com/shawpo/lagouWordCloud/utils"
	ui "github.com/gizak/termui"
)

// 岗位数据后缀
const POSITIONSEXT = ".txt"
const TABLEPAGECOUNT  = 15
var kd string

func main() {
	exists, err :=  utils.ExistPositions(".", POSITIONSEXT)
	if len(exists) > 0 && err == nil {
		selects := append(exists, "no")
		prompt := promptui.Select{
			Label: "--已有以下岗位数据，请选择其中一个进行分析，选择no表示获取新的岗位数据",
			Items: selects,
		}

		_, kd, err = prompt.Run()

		if err != nil {
			log.Fatalf("获取选择失败 %v\n", err)
		}
	}

	if kd != "no" && kd != "" {
		analysisKd(kd)
		return
	}

	prompt := promptui.Prompt{
		Label: "输入岗位名称",
	}

	kd, err = prompt.Run()

	if err != nil {
		log.Fatalf("获取岗位名称失败：%v\n", err)
	}
	getData(kd)
	analysisKd(kd)
}

func analysisKd(kd string)  {
	fmt.Println("--开始词频分析：")
	var segment sego.Segmenter
	segment.LoadDictionary(analysis.ITDIC)
	err, wordsMap := analysis.Analysis(kd + POSITIONSEXT, segment)
	var sum = len(wordsMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("--词频分析已完成！总计%d个关键词！\n--开始进行排序：\n", sum)
	wordRands := utils.RankByWordCount(wordsMap)
	fmt.Println("--排序已完成！")

	displayRank(wordRands)
}

func displayRank(ranks utils.WordCountList)  {
	header := [][]string{
		[]string{"关键词", "出现频次"},
	}
	ranks = ([][]string)(ranks)
	var sum = len(ranks)
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()
	var i, j = 0, 0+TABLEPAGECOUNT
	// 标题
	p := ui.NewPar(fmt.Sprintf("%s职位关键词(总数%d):q退出，w向上翻页，s向下翻页", kd, sum))
	p.Height = 3
	p.Width = 60
	p.TextFgColor = ui.ColorRed
	p.BorderLabel = fmt.Sprintf("排名为%d~%d的关键词", i+1, j)
	p.BorderFg = ui.ColorWhite

	table := ui.NewTable()
	table.Rows = append(header, ranks[i:j]...)
	table.FgColor = ui.ColorYellow
	table.Block.BorderFg = ui.ColorWhite
	table.Separator = false
	table.Width = 60
	table.Height = j - i + 3
	table.TextAlign = ui.AlignCenterHorizontal
	//ui.Render(p, table)
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(8, 2, p)),
		ui.NewRow(
			ui.NewCol(8, 2, table)))

	// calculate layout
	ui.Body.Align()

	ui.Render(ui.Body)

	// handle key q pressing
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})
	render := func(i, j int) {
		table.Rows = append(header, ranks[i:j]...)
		p.BorderLabel = fmt.Sprintf("排名为%d~%d的关键词", i+1, j)
		ui.Body.Rows[0].Cols[0] = ui.NewCol(8, 2, p)
		ui.Body.Rows[1].Cols[0] = ui.NewCol(8, 2, table)
		//table.Height = j - i + 3
		ui.Render(ui.Body)
	}

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Handle("/sys/kbd/s", func(e ui.Event) {
		if i+TABLEPAGECOUNT > sum{
			render(i, j)
			return
		}

		if j+TABLEPAGECOUNT > sum {
			j = sum
			i = i+TABLEPAGECOUNT
		} else {
			i = i+TABLEPAGECOUNT
			j = j+TABLEPAGECOUNT
		}
		render(i, j)

	})
	ui.Handle("/sys/kbd/w", func(e ui.Event) {
		if j-TABLEPAGECOUNT < 0 || i-TABLEPAGECOUNT < 0{
			render(i, j)
			return
		}
		j = i
		i = i-TABLEPAGECOUNT
		render(i, j)
	})
	ui.Loop()
}


func getData(kd string)  {
	i := 1
	sum := 0
	first := "true"
	positionsMap := make(map[int]string)
	fmt.Println("--开始获取职位列表：")
	for {
		if i > 1 {
			first = "false"
			utils.RandTimeSleep(1, 2)
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
			fmt.Printf("----No.%d   %d: %s\n",sum, position.PositionID, position.PositionName)
			positionsMap[position.PositionID] = position.PositionName
			sum++
		}
		resp.Body.Close()
		i++
	}

	filePath := kd + POSITIONSEXT

	file, err := os.Create(filePath)

	if err != nil {
		log.Fatalf("can't create file %s to save data", filePath)
	}

	if len(positionsMap) == 0 {
		log.Fatal("get none positions")
	}

	i = 1
	fmt.Println("--开始获取职位详细数据：")
	for positionId := range positionsMap {
		utils.RandTimeSleep(1,2)
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
			log.Printf("----PositionId %d: request error",
				positionId)
		}
		detail ,err := spider.GetDetail(bytes.NewReader(html))

		if detail == "" {

			filePath := fmt.Sprintf("error-%s-%d.html", kd, positionId)
			file, _ := os.Create(filePath)
			file.Write(html)
			log.Fatalf("PositionId %d: can't get detail, response has write in %s, maybe should update cookies",
				positionId, filePath)
		}
		if err != nil{
			log.Print(err)
			continue
		}
		fmt.Printf("----Write position %d to file: %d of %d\n", positionId, i, len(positionsMap))
		_, err = file.WriteString(detail)
		if err != nil {
			log.Print(err)
			continue
		}


		resp.Body.Close()
		i++
	}
}
