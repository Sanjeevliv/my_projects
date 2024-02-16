package main

import{
	"fmt"
	"log"
	"net/http"
}
//srksrk serk
func formHandler(w http.ResponseWriter, r *http.Request){
	if err := rParseForm(); err !=nil {
		fmt.Fprintf(w. "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name :=r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name =%s/n", neme)
	fmt.Fprintf(w, "Address = %s/n", address)
}

funchelloHandler(w http.ResponceWriter, r *http.Request){
if r.URLPath !="/hello" {
	http.Error(w, "404 not found", http.StatusNotFound)
	return
}
if r.Method !="GET"{
	http.Error(w, "method is not supported", http.StatusNotFound)
	return
}
fmt.Fprintf(w, "hello!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HanleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Prinf("Starting server at port 8000/n")
	if err := http.ListenAndServer(":8080", nil); err !=nil {
		log.Fatal(err)
	}
}