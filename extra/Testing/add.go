package main

import (
	"fmt"
)
func main() {
 anInt := 42
anotherInt := 56
fmt.Println("Sum of anInt and anotherInt is:", add(anInt, anotherInt))
}
func add(a int, b int) int {
	return a + b
}