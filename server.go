package main


import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World  %s!", request.URL.Path[1:])
}

func handlerTest(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "This is test %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", handlerTest)
	http.ListenAndServe(":8080", nil)
}