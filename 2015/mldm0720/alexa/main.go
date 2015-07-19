package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func fetch(link string) ([]string, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	parser, err := regexp.Compile(`(?s)<a href="/siteinfo/.*?">(.*?)</a>`)
	if err != nil {
		return nil, err
	}
	content := parser.FindAllStringSubmatch(string(body), -1)
	var results = make([]string, 0)
	for _, c := range content {
		results = append(results, c[1])
	}
	return results, nil

}

func main() {
	url := "http://www.alexa.com/topsites/global;"
	var links = make([]string, 0)
	for i := 0; i < 20; i++ {
		link := fmt.Sprintf("%s%d", url, i)
		r, _ := fetch(link)
		links = append(links, r...)
	}
	for _, c := range links {
		fmt.Println(c)
	}
}
