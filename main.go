package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request)  {
	if r.URL.Path!="/hello"{
		http.Error(w,"404 path not found",http.StatusNotFound)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"Method not applicable, use GET",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello world")
}


func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	// http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("server starting on port 8080")
	if err:= http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal(err)
	}
}
