package main

import(
	"fmt"
	"net/http"

)

func main() {

	str:="Hello World"
	fmt.Print()
	fmt.Print(str)
	fmt.Printf("Server Start")

	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"Welcome!")
	})
	http.HandleFunc("/about",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"about")
	})

	http.ListenAndServe(":8080",nil);
}
