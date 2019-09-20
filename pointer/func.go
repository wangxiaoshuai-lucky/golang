package main

import "fmt"

/*
 函数类型
*/

type Op func(int, int) int

func do(op Op, a int, b int) {
	fmt.Println(op(a, b))
}

func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Printf("%T\n", add)
	do(add, 10, 10)
}
