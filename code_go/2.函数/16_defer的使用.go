package main //必须

import "fmt"

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main() {
	//defer延迟调用，main函数结束前调用
	defer fmt.Println("bbbbbbbbbbbbb")
	fmt.Println("aaaaaaaaaaaaaaa")
	f(3)
}
