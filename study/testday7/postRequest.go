package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	resp, err := http.PostForm("http://127.0.0.1:8099/psot",
		url.Values{"name": {"qqa"}, "age": {"1998"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

func main1() {
	apiUrl := "http://127.0.0.1:8099/psot"
	data := `{"name":"qqa","age":"1998"}`
	contentType := "application/json"

	resp, err := http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("error from NewReader")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error from ReadAll")
	}

	fmt.Println(string(body))

}
