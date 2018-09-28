package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println(request.Form)
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	for key, value := range request.Form {
		fmt.Println("key:", key)
		fmt.Println("val:", strings.Join(value, ""))
	}
	fmt.Fprintf(writer, "Hello World")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
