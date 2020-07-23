package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main () {
	resp, err := http.Get("http://127.0.0.1:8099/get")
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