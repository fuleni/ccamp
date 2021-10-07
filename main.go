package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func getHeader(res http.ResponseWriter, req *http.Request) {
	if req.Header == nil {
		return
	}
	// 将响应头设置和请求头一样
	for key, val := range req.Header {
		res.Header().Set(key, strings.Join(val, ";"))
	}
	version := os.Getenv("VERSION")
	if version == "" {
		log.Println("version not exists or version is \"\"")
	}
	// 设置version
	res.Header().Set("version", version)
	fmt.Fprintln(res, res.Header())
	log.Println(req.RemoteAddr, "200")
}


func healthz(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, 200)
}

func setEnv() {
	os.Setenv("version", "1.0")
}

func main() {
	setEnv()
	server := http.Server{
		Addr: "127.0.0.1:8090",
	}
	http.HandleFunc("/getHeader", getHeader)
	http.HandleFunc("/healthz", healthz)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
