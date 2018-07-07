package utils

import (
	"sort"
	"strconv"
)

// 用于将map排序为slice
// github.com/gizak/termui 的表格组件需要 [][]string类型的数据

type WordCountList [][]string
func (p WordCountList) Len() int { return len(p) }
func (p WordCountList) Less(i, j int) bool {
	valuei, _ :=  strconv.Atoi(p[i][1])
	valuej, _ :=  strconv.Atoi(p[j][1])
	return valuei < valuej
}
func (p WordCountList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }
func RankByWordCount(wordFrequencies map[string]int) WordCountList{
	pl := make(WordCountList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = []string{k, strconv.Itoa(v)}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}