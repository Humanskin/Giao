package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func httpHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//data := r.URL.Query()
	//fmt.Println(data.Get("name"))
	//fmt.Println(data.Get("age"))

	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"))
	fmt.Println(r.PostForm.Get("age"))

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error from ReadAll")
	}

	fmt.Println(string(body))

	answers := `{"status":"ok"}`
	w.Write([]byte(answers))
}

func main() {
	http.HandleFunc("/", httpHandle)
	http.ListenAndServe(":8099", nil)

}
