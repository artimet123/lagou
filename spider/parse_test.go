package spider

import (
	"testing"
	"os"
	"strings"
)

func TestGetDetail(t *testing.T) {

	var testData = []struct{
		path string
		contains string
	} {
		{
			"testdata/TestGetDetail.html",
			"GO",
		},
	}

	for _, test := range testData{


		file, err := os.Open(test.path)

		if err != nil {
			t.Fatal(err)
		}

		detail, err := GetDetail(file)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(detail, test.contains) {
			t.Fatalf("want contains %s", test.contains)
		}
	}
}

func TestGetPositions(t *testing.T) {
	var testData = []struct{
		path  string
		ids   []int
		count int
	}{
		{
			"testdata/TestGetPositions.json",
			[]int{ 2962839, 4195931, 4054423, 4492533, 4215065},
			15,
		},
	}

	for _, test := range testData{

		file, err := os.Open(test.path)

		if err != nil {
			t.Fatal(err)
		}
		positions, err := GetPositions(file)

		if err != nil {
			t.Fatal(err)
		}

		if len(positions) != test.count {
			t.Fatalf("want positions count %d, get %d",test.count, len(positions))
		}

		for i := 0; i < len(test.ids); i++ {
			if positions[i].PositionID != test.ids[i] {
				t.Fatalf("want position id %d, get %d", test.ids[i], positions[i].PositionID )
			}
		}
	}

}