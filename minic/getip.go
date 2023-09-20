package minic

import (
	"io"
	"net/http"
)
//Если ip не отдаётся то скорее всего проблеммы с сайтом https://api.ipify.org.
func Getip() string {
	resp, _ := http.Get("https://api.ipify.org")

	body, _ := io.ReadAll(resp.Body)

	return string(body)
}