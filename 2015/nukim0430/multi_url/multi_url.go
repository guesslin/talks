package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// fetch_url 抓取給定的網址，並將結果回傳到指定的channel內
func fetch_url(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- ""
	}
	ch <- string(body)
}

func main() {
	counter := 0 // 宣告計數器
	urls := []string{"http://www.google.com/",
		"http://tw.yahoo.com/",
		"http://www.im.nuk.edu.tw/"}
	ch := make(chan string) // 宣告 ch 作為與goroutine 溝通之管道，傳遞string型別的資料
	for i := range urls {
		go fetch_url(urls[i], ch)
		counter++
	} // 分派任務給goroutine
	for {
		select {
		case res := <-ch:
			fmt.Println(res)
			counter--
			if counter == 0 {
				close(ch)
				return
			}
		case <-time.After(time.Second * 1):
			fmt.Println("time out second 1")
		}
	} // 等待所有工作完成
}
