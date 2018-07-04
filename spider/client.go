package spider

import (
	"net/http"
)


func GetClient() *http.Client{
	client := &http.Client{
		Jar: CookieJar,
	}
	return client
}