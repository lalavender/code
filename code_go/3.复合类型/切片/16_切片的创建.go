package main //必须有个main包

import "fmt"
func main() {

	//自动推导类型， 同时初始化
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1 = ", s1)

	//借助make函数, 格式 make(切片类型, 长度, 容量)
	s2 := make([]int, 5, 10)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))

	//没有指定容量，容量和长度一样
	s3 := make([]int, 5)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	//切片和数组的区别
	//数组[]里面的长度时固定的一个常量， 数组不能修改长度， len和cap永远都是5
	a := [5]int{}
	fmt.Printf("len = %d, cap = %d\n", len(a), cap(a))

	//切片， []里面为空，或者为...，切片的长度或容易可以不固定
	s := []int{}
	fmt.Printf("1: len = %d, cap = %d\n", len(s), cap(s))

	s = append(s, 11) //给切片末尾追加一个成员
	fmt.Printf("append: len = %d, cap = %d\n", len(s), cap(s))


    s4 := make([]string, 3)
    fmt.Println("emp:", s)

    s4[0] = "a"
    s4[1] = "b"
    s4[2] = "c"
    fmt.Println("set:", s4)
    fmt.Println("get:", s4[2])

    fmt.Println("len:", len(s4))

    s4 = append(s4, "d")
    s4 = append(s4, "e", "f")
	fmt.Println("apd:", s4)
	//如果超过原来的容量，通常以2倍容量扩容
	s5 := make([]int, 0, 1) //容量为1
	oldCap := cap(s5)
	for i := 0; i < 20; i++ {
		s = append(s5, i)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("cap: %d ===> %d\n", oldCap, newCap)

			oldCap = newCap
		}
	}

    c := make([]string, len(s))
    copy(c, s4)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
