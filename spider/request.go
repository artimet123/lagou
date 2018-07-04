package spider

import (
	"net/http"
	"net/url"
	"strings"
	"fmt"
)

func KdPositionsRequest(values url.Values) (request *http.Request, err error) {
	request, err = http.NewRequest("POST",
		"https://www.lagou.com/jobs/positionAjax.json?needAddtionalResult=false",
			strings.NewReader(values.Encode()))
	if err != nil {
		return
	}
	request.Header = http.Header{
		"content-type": {"application/x-www-form-urlencoded; charset=UTF-8"},
		"Accept-Encoding": {"gzip, deflate"},
		"Host": {"www.lagou.com"},
		"Origin": {"http://www.lagou.com"},
		"User-Agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36"},
		"X-Requested-With": {"XMLHttpRequest"},
		"Referer": {"https://www.lagou.com/jobs/list_"+ values.Get("kd") +"?labelWords=&fromSearch=true&suginput="},
		"Proxy-Connection": {"keep-alive"},
		"X-Anit-Forge-Code": {"0"},
		"X-Anit-Forge-Token": {"None"},

	}
	return
}

func DetailRequest(positionId int) (request *http.Request, err error) {
	detailUrl := fmt.Sprintf("https://www.lagou.com/jobs/%d.html", positionId)
	request, err = http.NewRequest("GET", detailUrl, nil)
	if err != nil {
		return
	}
	//request.Header = http.Header{
	//	"Accept": {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"},
	//	"Accept-Encoding": {"gzip, deflate, br"},
	//	"Host": {"www.lagou.com"},
	//	"Origin": {"http://www.lagou.com"},
	//	"User-Agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36"},
	//	"Referer": {detailUrl},
	//	"Cookie": {`user_trace_token=20180614102812-9acd5884-da1b-4a29-880a-ea08ca43f3e4; LGUID=20180614102814-96dbf9a6-6f7a-11e8-a3af-525400f775ce; _ga=GA1.2.1068078087.1528943294; X_HTTP_TOKEN=b4fa97ace355eb0a05b77f3a1c0c9d0d; LG_LOGIN_USER_ID=dd1f278f34f12c092a909a88914e684fc8343ffb49507844; _putrc=943AEA2F6E967FB3; login=true; unick=%E6%AD%A6%E5%B0%8F%E5%8D%9A; witkey_login_authToken="R7naTX31+N8ObPCrjslS6CanTnVXv0ERtVkEyiiNN9VByElUeUhWguMn6TA76OTFuU/GEfdudxzLTGFu2rT+wPrmHTAAkR4HU6+ae8IecM79XRkLi/82f8uPQHq2T49p6rgVBEeVpvs8Fk5yaDef2WgZ/3+e5dESM/jIRJA6C554rucJXOpldXhUiavxhcCELWDotJ+bmNVwmAvQCptcy5e7czUcjiQC32Lco44BMYXrQ+AIOfEccJKHpj0vJ+ngq/27aqj1hWq8tEPFFjdnxMSfKgAnjbIEAX3F9CIW8BSiMHYmPBt7FDDY0CCVFICHr2dp5gQVGvhfbqg7VzvNsw=="; _gid=GA1.2.1984759381.1530534395; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1528943294,1530098405,1530534396; JSESSIONID=ABAAABAACBHABBIBA280B2651A93B66B87FB3433980E4E9; index_location_city=%E5%85%A8%E5%9B%BD; showExpriedIndex=1; showExpriedCompanyHome=1; showExpriedMyPublish=1; hasDeliver=0; TG-TRACK-CODE=search_code; gate_login_token=f0394d035dcb2d4699ecd37b3e9b48973cbf43b66e75f573; SEARCH_ID=256b46e5cc2a40c09023a3da0247d5cc; _gat=1; LGSID=20180704115713-555285b5-7f3e-11e8-bdf6-525400f775ce; PRE_UTM=; PRE_HOST=; PRE_SITE=; PRE_LAND=https%3A%2F%2Fwww.lagou.com%2Fjobs%2F2962839.html; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1530676659; LGRID=20180704115740-65e13c56-7f3e-11e8-98e3-5254005c3644`},
	//}
	return
}
