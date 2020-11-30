package main

import "fmt"

// source: https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-function-call
// 传值会被拷贝，原变量不受影响；
// 传指针也会发生拷贝，原变量受影响；

type MyStruct struct {
	i int
}

func myFunction(a MyStruct, b *MyStruct) {
	a.i = 31
	b.i = 41
	fmt.Printf("in_myFunction  -a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
}

func main() {
	a := MyStruct{i: 30}
	b := &MyStruct{i: 40}
	fmt.Printf("before calling -a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
	myFunction(a, b)
	fmt.Printf("after calling  -a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
}
