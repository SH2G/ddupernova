package main

import "fmt"

func main() {
	var a int
    fmt.Scanf("%d", &a)
	r := a * (a - 1)
        
    fmt.Printf("%d\n", r)
}