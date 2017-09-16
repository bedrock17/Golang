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

	//bind handlers
	mainRouter.HandleFunc("GET", "/", index)

	mainRouter.HandleFunc("GET", "/frontpage", frontpage)

	mainRouter.HandleFunc("GET", "/menu/:item", menu)

	http.ListenAndServe(":8080", mainRouter)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome! It's main")
	fmt.Println("Hello World end")
}

func frontpage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "It's Front Page!")
}

func menu(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, req.URL.Path)
	fmt.Println("MENU", req.URL.Path)
}
