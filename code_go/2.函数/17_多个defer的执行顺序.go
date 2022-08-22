package main //必须

import "fmt"

func test1(x int) {
	result := 100 / x

	fmt.Println("result = ", result)
	defer fmt.Println("test1")
}

func main() {

	defer fmt.Println("aaaaaaaaaaaaaa")
	defer fmt.Println("bbbbbbbbbbbbbb")

	//调用一个函数，导致内存出问题
	defer test1(4)
	defer fmt.Println("ccccccccccccccc")
}
