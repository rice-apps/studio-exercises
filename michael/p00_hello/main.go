package main

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hi %s!", name)
}

func main() {
	fmt.Println(Greet("Michael"))
}
