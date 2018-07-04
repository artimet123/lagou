package spider

import (
	"encoding/json"
	"io/ioutil"
	"io"
	"github.com/PuerkitoBio/goquery"
)

func GetPositions(reader io.Reader) (positions []Position, err error) {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}

	positionResp := PositionResp{}
	err = json.Unmarshal(body, &positionResp)
	if err != nil {
		return
	}

	positions = positionResp.Content.PositionResult.Result
	return
}

func GetDetail(body io.Reader) (detail string, err error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return
	}

	detail = doc.Find(".job_bt div").Text()
	return
}