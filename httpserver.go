// httpserver.go
package main

import (
	//"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() //http 요청 멀티플렉서 생성

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!")) //웹 브라우저에 응답
	}) // "/" 경로에 접속했을 때 실행할 함수 설정

	http.ListenAndServe(":80", mux) //80번 포트에서 웹 서버 실행
}
