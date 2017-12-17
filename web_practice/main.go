package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintln(writer, "hello world, ", request.URL.Path)
}


func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}