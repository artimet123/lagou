package analysis

import (
	"testing"
	"github.com/shawpo/sego"
)



func TestSeg(t *testing.T)  {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("THUOCL_it.txt")
	t.Log(sego.SegmentsToString(segmenter.Segment([]byte(`
		岗位职责：
		1、负责企业及应用的系统架构设计与开发；
		2、负责实现具体开发编码工作，并对系统质量负主要责任；
		3、负责windows或Linux平台下后台程序开发；
		4、负责数据库的设计与开发；
		5、完成上级交办的其他工作。
		6、对主要的模块进行优化和设计。
		任职要求：
		1、计算机或相关专业本科以上学历，具有扎实的计算机基础理论知识；
		2、有1年以上后台开发工作经验，能独立完成项目，有完整项目经验者优先；
		3、熟悉Go语言，具有go开发经验
		4、熟悉使用Nginx，Redis，MySQL等
		5、熟悉Docker
		6、熟悉通用通信协议：http，rpc 等
		 
		互联网中心福利：
		1、法定假期：按国家统一规定员工享有法定节假日。
		2、带薪年假：员工可享受最高15天的带薪年假。
		3、五险一金：为员工提供社会保险及公积金福利。
		4、幸福基金：为员工提供低息借款，解决购房、购车、结婚等大额支出困难。
		5、心基金：为困难员工提供爱心资助，帮助员工及其家人度过难关。
		6、节日福利：春节、三八、端午、中秋节、教师节等，公司发放节日礼品。
		7、年度体检：为员工提供年度免费体检。
		8、定期团建：烧烤、唱K、聚餐等，增强团队凝聚力。
		9、定期省内外旅游：如海南、云南、西藏等。
		10、员工活动：全天候无限量零食供应，每周水果下午茶，每月不同主题生日PARTY，轻松欢快的工作氛围。
		11、工作环境：坐标广州天河市CBD，交通便捷，办公环境怡人。
		12、课程优惠：员工子女报读课程可享受一定优惠，低至5折优惠。
		13、完善的晋升机制，多通道全方位发展。
		 
		您是否愿意加入我们，成为12000余名人才中令人骄傲与艳羡的一员？`)), false))
}

func TestIsValid(t *testing.T)  {
	testdata := []struct{
		word string
		valid bool
	} {
		{
			"1000",
			false,
		},
		{
			"c++",
			true,
		},
		{
			"数据结构",
			true,
		},
	}

	for _, test := range testdata{
		if valid := isValid(test.word); valid != test.valid {
			t.Fatalf("%s's valid, want %v, get %v", test.word, test.valid, valid)
		}
	}
}


func TestAnalysisCount(t *testing.T) {
	testdata := []struct{
		dicPath string
		contentPath string
		wordsCount []struct{
			word string
			count int
		}
	} {
		{
			"testdata/dic.txt",
			"testdata/content.txt",
			[]struct{
				word string
				count int
			} {
				{
					"后端",
					3,
				},
				{
					"服务",
					3,
				},
				{
					"数据结构",
					1,
				},
				{
					"数据",
					1,
				},
			},
		},
	}

	for _, test := range testdata {
		var segmenter sego.Segmenter
		segmenter.LoadDictionary(test.dicPath)
		err, words := Analysis(test.contentPath,  segmenter)

		if err != nil {
			t.Fatal(err)
		}
		t.Log(words)
		for _, wordCount := range test.wordsCount {
			if words[wordCount.word] != wordCount.count {
				t.Fatalf("%s's count: want %d, get %d",
					wordCount.word, wordCount.count, words[wordCount.word])

			}
		}
	}

}

func TestSynonymMap(t *testing.T) {
	testdata := []struct{
		dicPath string
		count int
		Synonyms []struct{
			word string
			synonym string
		}
	} {
		{
			"testdata/synonym.txt",
			2,
			[]struct{
				word string
				synonym string
			} {
				{
					"go语言",
					"go",
				},
				{
					"golang",
					"go",
				},
			},
		},
	}



	for _, test := range testdata {
		synonymMap := SynonymMap(test.dicPath)
		t.Log(synonymMap)
		if count := len(synonymMap); count != test.count {
			t.Fatalf("synonys count need %d, but get %d", test.count, count)
		}

		for _, synonym := range test.Synonyms{
			if synonym.synonym != synonymMap[synonym.word] {
				t.Fatalf("%s's synonym: need %s, but get %s", synonym.word, synonym.synonym, synonymMap[synonym.word])
			}
		}
	}
}

func TestFilterMap(t *testing.T) {
	testdata := []struct{
		dicPath string
		count int
		filters []struct{
			word string
		}
	} {
		{
			"testdata/filter.txt",
			3,
			[]struct{
				word string
			} {
				{
					"com",
				},
				{
					"www",
				},
				{
					"the",
				},
			},
		},
	}



	for _, test := range testdata {
		filtersMap := FilterMap(test.dicPath)
		t.Log(filtersMap)
		if count := len(filtersMap); count != test.count {
			t.Fatalf("filters count need %d, but get %d", test.count, count)
		}

		for _, filter := range test.filters{
			if _, exist := filtersMap[filter.word]; !exist  {
				t.Fatalf("need %s, but get don't have", filter.word)
			}
		}
	}
}

