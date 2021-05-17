package main

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Total         int    `json:"total"`
	Message       string `json:"msg"`
	Code          int    `json:"code"`
	GauthUserInfo `json:"data"`
}

type GauthUserInfo struct {
	Username string `json:"loginName"`
	Password string `json:"password"`
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	result := &Result{
		Total:   1,
		Message: "操作成功",
		Code:    200,
		GauthUserInfo: GauthUserInfo{
			Username: "test",
			Password: "123456",
		},
	}
	json, _ := json.Marshal(result)
	w.Write(json)
}

func main() {

	server := http.Server{
		Addr: "localhost:8888",
	}
	http.HandleFunc("/common/getUserBykey", jsonExample)
	server.ListenAndServe()

}
