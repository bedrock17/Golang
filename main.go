package main

import (
	"./router"
	"fmt"
	"net/http"
)

func main() {

	str := "Hello World"
	fmt.Print()
	fmt.Println(str)
	fmt.Println("Server Start")

	mainRouter := &router.Router{make(map[string]map[string]http.HandlerFunc)}

	mainRouter.HandleFunc("GET", "/", index)

	mainRouter.HandleFunc("GET", "/frontpage", frontpage)

	http.ListenAndServe(":8080", mainRouter)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome!")
	fmt.Println("Hello World end")
}

func frontpage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "It's Front Page!")
}
