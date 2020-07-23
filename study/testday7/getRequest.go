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

	urls, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Sprintf("error from ParseRequestURI")
	}

	urls.RawQuery = data.Encode()

	resp, err :=  http.Get(string(urls.String()))
	if err != nil {
		fmt.Println("resp error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error from ReadAll")
	}

	fmt.Println(string(body))

}