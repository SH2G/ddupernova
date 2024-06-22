package main

import "fmt"

func main() {
    var s1 string
    fmt.Scan(&s1)
    for _, s := range s1 {
        if s < 97 {
			fmt.Print(string(s + 32))
	    } else {
			fmt.Print(string(s - 32))
		}
	}
}