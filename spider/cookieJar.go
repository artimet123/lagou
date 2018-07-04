package spider

import (
	"net/url"
	"net/http/cookiejar"
	"net/http"
)

var CookieJar *cookiejar.Jar

func init()  {
	var err error
	CookieJar, err = cookiejar.New(&cookiejar.Options{})
	if err != nil{
		return
	}
	cookies := []*http.Cookie {
		{
			Name: "Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6",
			Value: "1530676663",
		},
		{
			Name: "Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6",
			Value: "1528943294,1530098405,1530534396",
		},
		{
			Name: "JSESSIONID",
			Value: "ABAAABAACEFAACG48CB23BE6AEE57048631756283BF2611",
		},
		{
			Name: "JSESSIONID",
			Value: "ABAAABAACBHABBIBA280B2651A93B66B87FB3433980E4E9",
		},
		{
			Name: "LGRID",
			Value: "20180704115742-66c1e1dc-7f3e-11e8-bdf6-525400f775ce",
		},
		{
			Name: "LGUID",
			Value: "20180704115713-555285b5-7f3e-11e8-bdf6-525400f775ce",
		},
		{
			Name: "LG_LOGIN_USER_ID",
			Value: "dd1f278f34f12c092a909a88914e684fc8343ffb49507844",
		},
		{
			Name: "SEARCH_ID",
			Value: "256b46e5cc2a40c09023a3da0247d5cc",
		},
		{
			Name: "TG-TRACK-CODE",
			Value: "index_search",
		},
		{
			Name: "X_HTTP_TOKEN",
			Value: "b4fa97ace355eb0a05b77f3a1c0c9d0d",
		},
		{
			Name: "_ga",
			Value: "GA1.2.1068078087.1528943294",
		},
		{
			Name: "_gid",
			Value: "GA1.2.1984759381.1530534395",
		},
		{
			Name: "_putrc",
			Value: "943AEA2F6E967FB3",
		},
		{
			Name: "gate_login_token",
			Value: "f0394d035dcb2d4699ecd37b3e9b48973cbf43b66e75f573",
		},
		{
			Name: "hasDeliver",
			Value: "0",
		},
		{
			Name: "index_location_city",
			Value: "%E5%85%A8%E5%9B%BD",
		},
		{
			Name: "login",
			Value: "true",
		},
		{
			Name: "showExpriedCompanyHome",
			Value: "1",
		},
		{
			Name: "showExpriedIndex",
			Value: "1",
		},
		{
			Name: "showExpriedMyPublish",
			Value: "1",
		},
		{
			Name: "unick",
			Value: "%E6%AD%A6%E5%B0%8F%E5%8D%9A",
		},
		{
			Name: "user_trace_token",
			Value: "20180614102812-9acd5884-da1b-4a29-880a-ea08ca43f3e4",
		},
		{
			Name: "witkey_login_authToken",
			Value: "R7naTX31+N8ObPCrjslS6CanTnVXv0ERtVkEyiiNN9VByElUeUhWguMn6TA76OTFuU/GEfdudxzLTGFu2rT+wPrmHTAAkR4HU6+ae8IecM79XRkLi/82f8uPQHq2T49p6rgVBEeVpvs8Fk5yaDef2WgZ/3+e5dESM/jIRJA6C554rucJXOpldXhUiavxhcCELWDotJ+bmNVwmAvQCptcy5e7czUcjiQC32Lco44BMYXrQ+AIOfEccJKHpj0vJ+ngq/27aqj1hWq8tEPFFjdnxMSfKgAnjbIEAX3F9CIW8BSiMHYmPBt7FDDY0CCVFICHr2dp5gQVGvhfbqg7VzvNsw==",
		},
	}
	lagou , err:= url.Parse("https://www.lagou.com")
	if err != nil {
		return
	}
	CookieJar.SetCookies(lagou, cookies)
}


