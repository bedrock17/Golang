package main

import(
	"fmt"
	"net/http"
	"./router"
)

func main() {

	str := "Hello World"
	fmt.Print()
	fmt.Println(str)
	fmt.Println("Server Start")

	mainRouter := &router.Router{make(map[string]map[string]http.HandlerFunc)}

	mainRouter.HandleFunc("GET", "/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Welcome!")
		fmt.Println("Hello World end");
	})


	http.ListenAndServe(":8080", mainRouter);
}
