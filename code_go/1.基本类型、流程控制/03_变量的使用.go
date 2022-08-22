package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {
	//变量，程序运行期间，可以改变的量

	var a = "initial" //声明格式 var 变量名 类型, 变量声明了，必须要使用
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e) //只是声明没有初始化的变量，默认值为0

	f := "short"
	fmt.Println(f)

	var g int = 10 //初始化，声明变量时，同时赋值
	g = 20         //赋值， 先声明，后赋值
	fmt.Println("g = ", g)

	//自动推导类型，必须初始化，通过初始化的值确定类型(常用)
	h := 30
	//%T打印变量所属的类型
	fmt.Printf("h type is %T\n", h)

	h = 30 //重新赋值
	fmt.Println("h = ", h)

	//赋值，赋值前，必须先声明变量
	var i int
	i = 10 //赋值，可以使用n次
	i = 20
	i = 30
	fmt.Println("i = ", i)

}
