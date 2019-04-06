package main

import (
	"fmt"
	"net/http"

	"./config"
	"./router"
)

func main() {

	//서버 설정 로딩
	var config config.Config
	err := config.Load("serverconfig.ini")
	if err != nil {
		fmt.Println(err)
	}
	var serverPort string
	serverPort = fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Server Start [%s] \n", serverPort)

	mainRouter := &router.Router{make(map[string]map[string]http.HandlerFunc)}

	//bind handlers
	mainRouter.HandleFunc("GET", "/", index)

	mainRouter.HandleFunc("GET", "/frontpage", frontpage)

	mainRouter.HandleFunc("GET", "/menu/:item", menu)

	http.ListenAndServe(serverPort, mainRouter)
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
