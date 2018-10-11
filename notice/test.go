package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func main() {
	http.HandleFunc("/notice", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Code: http.StatusOK,
			Data: "一个Web 服务器",
			Msg:  "",
		}
		rep, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(rep)
		//w.Write([]byte("一个Web 服务器"))
	})
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err)
	}
}
