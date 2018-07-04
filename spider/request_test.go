package spider

import (
	"testing"
	"io/ioutil"
	"net/url"
)

func TestDetailRequest(t *testing.T) {
	positionId := 2735719
	request, err := DetailRequest(positionId)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := GetClient().Do(request)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatal(string(body))
	}
}

func TestKdPositionsRequest(t *testing.T) {
	values:= url.Values{
		"first": {"true"},
		"pn" : {"1"},
		"kd": {"go"},
	}
	request, err := KdPositionsRequest(values)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := GetClient().Do(request)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatal(string(body))
	}
}