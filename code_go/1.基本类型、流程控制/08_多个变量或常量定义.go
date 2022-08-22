package main //必须有一个main包

import "fmt"

func main() {
	//可以自动推导类型
	var (
		a = 1
		b = 2.0
	)

	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//可以自动推导类型
	const (
		i = 10
		j = 3.14
	)

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

}
