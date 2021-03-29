package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Total    int    `json:"total"`
	Message  string `json:"message"`
	Code     int    `json:"code"`
	UserInfo `json:"data"`
}

type UserInfo struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	url := "http://localhost:8888/getUser"
	res, _ := http.Get(url)
	if res.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		var result Result
		json.Unmarshal([]byte(body), &result)
		fmt.Println("json:", json.Unmarshal([]byte(body), &result))
		fmt.Printf("打印Result: %v\n", result)
	}

}
