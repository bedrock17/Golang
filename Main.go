package main

import("fmt")

func F(value int) *int {

	p:=new(int)
	*p=10
	return p
}

func main() {

	fmt.Println("Hello World")

	test := F(10)

	fmt.Println(*test)



}
