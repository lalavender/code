package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	//变量：程序运行期间，可以改变的量， 变量声明需要var
	//常量：程序运行期间，不可以改变的量，常量声明需要const

	const a int = 10
	//a = 20 //err, 常量不允许修改
	fmt.Println("a = ", a)

	const b = 11.2 //没有使用:=
	fmt.Printf("b type is %T\n", b)
	fmt.Println("b = ", b)

	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
