package main

import (
	"fmt"
	"net/http"
)

func httpHeaders (w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	fmt.Println(data.Get("addr"))

	answers := `{"status":"ok"}`

	w.Write([]byte(answers))
}

func main ()  {
	http.HandleFunc("/", httpHeaders)
	http.ListenAndServe(":8099", nil)
}