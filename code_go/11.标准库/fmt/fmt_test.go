package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPrint1(t *testing.T) {
	fmt.Print("this is a test!")
	fmt.Println("this is a test!")
	fmt.Printf("this is a test info:%s", "aaa")
}

type User struct {
	Id int64
}

func TestPrint2(t *testing.T) {
	user := &User{Id: 1}
	// struct
	fmt.Printf("%v\n", user)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%T\n", user) // 类型
	fmt.Printf("%%\n")
	//bool
	fmt.Printf("%t\n", true)

	//int
	fmt.Printf("%b\n", 123) //二进制
	fmt.Printf("%c\n", 123) // unicode码值
	fmt.Printf("%d\n", 123)
	fmt.Printf("%o\n", 123)
	fmt.Printf("%x\n", 123)
	fmt.Printf("%X\n", 123)
	fmt.Printf("%U\n", 123)
	a := 96
	fmt.Printf("%q\n", a)
	fmt.Printf("%q\n", 0x4E2D)

	//float
	f := 18.54
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%10f\n", f) //宽度
	fmt.Printf("%.2f\n", f) // 精度
	fmt.Printf("%10.2f\n", f)
	fmt.Printf("%F\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)

	// str and []byte
	s := "我是字符串"
	b := []byte{65, 66, 67}
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", b)
	fmt.Printf("%q\n", s)
	fmt.Printf("%q\n", b)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)

	//指针
	fmt.Printf("%p\n", &s)

	// other
	fmt.Printf("%    d\n", 10)
	fmt.Printf("%10s\n", s)
	fmt.Printf("%-10s\n", s)
	fmt.Printf("%-10.2f\n", 10.56)
	fmt.Printf("%010s\n", s)

}

func TestFPrint(t *testing.T) {
	fmt.Fprint(os.Stdout, "this is a demo!")
	fmt.Fprintln(os.Stdout, "this is a demo!")
	fmt.Fprintf(os.Stdout, "this is a demo!")
}
func TestFPrint1(t *testing.T) {
	s := os.FileMode(0777).String()
	fmt.Println(s)
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(file, "this is a demo!")
	fmt.Fprintln(file, "this is a demo!")
	fmt.Fprintf(file, "this is a demo!")
}

func TestSPrint1(t *testing.T) {
	host := "localhost"
	port := 6379
	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Println(addr)

}
func TestErrorF(t *testing.T) {
	host := "localhost"
	port := 6379
	err := fmt.Errorf("ERROR:%s:%d", host, port)
	if err != nil {
		panic(err)
	}
}
