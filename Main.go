package main

import("fmt")

func main() {

	arr:=[10][10]int{}

	for i:=0 ; i<10;i++{
		for j:=0;j<10;j++{
			arr[i][j]=i*j
		}

	}

	for i:=0 ; i<10;i++{
		for j:=0;j<10;j++{
			fmt.Printf("%d ",arr[i][j])
		}
		fmt.Println()
	}
	fmt.Printf("Hello World")


}
