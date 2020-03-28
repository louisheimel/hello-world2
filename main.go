package main

import "fmt"

func main() {
fmt.Println("LONG BACK")
fmt.Println(Subtract(2))
}

func Subtract(val int) int {
	return val * val 
}
