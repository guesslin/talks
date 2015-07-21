package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetch_url(url string) (string, error) {
	resp, err := http.Get(url)
	// 檢查收到的GET回應是否有錯誤產生
	if err != nil {
		return "", err
	}
	// 稍後將用完的 http 回應關閉
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// 檢查讀取回應內容時是否有錯誤產生
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	var url string
	fmt.Println("input any valid url: ")
	fmt.Scanln(&url)
	body, _ := fetch_url(url)
	fmt.Printf("%s\n", body)
}
