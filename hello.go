package main

import "fmt"

func main() {
	fmt.Println(Hello("André"))
}

func Hello(name string) string {
	return "Hello, " + name
}
