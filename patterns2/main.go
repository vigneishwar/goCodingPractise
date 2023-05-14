package main

import "fmt"

func pattern2(x int) {

	for i := 0; i < x; i++ {
		for j := 0; j <=i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	x:=5
	pattern2(x)
}
