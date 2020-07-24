package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	apiUrl := "http://127.0.0.1:8099/"

	data := url.Values{}
	data.Add("name", "stt")
	data.Add("age", "18")
	data.Add("addr", "everywhere")

	urls, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println("url is error")
	}

	urls.RawQuery = data.Encode()

	fmt.Println(urls.String())

	resp, err := http.Get(urls.String())
	if err != nil {
		fmt.Println("resp is error")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body is error")
	}

	fmt.Println(string(body))

}
