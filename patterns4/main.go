package main

import "fmt"

func pattern4(x int) {
	for i := 1; i < x; i++ {
		for j := 1; j <=i; j++ {
			fmt.Print(i," ")
		}
		fmt.Println()
	}
}

func main() {
	x:=6
	pattern4(x)
}
