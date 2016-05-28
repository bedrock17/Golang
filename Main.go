package main

import(
	"fmt"
	"net/http"


)




func main() {



	fmt.Printf("Server Start")

	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"Welcome!")
	})
	http.HandleFunc("/about",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"about")
	})

	http.ListenAndServe(":8080",nil);
}
