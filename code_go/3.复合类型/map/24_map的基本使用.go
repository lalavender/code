package main //必须有个main包

import "fmt"

func main() {
	//定义一个变量， 类型为map[int]string
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	//对于map只有len，没有cap
	fmt.Println("len = ", len(m1))

	//可以通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))

	//可以通过make创建，可以指定长度，只是指定了容量，但是里面却是一个数据也没有
	m3 := make(map[int]string, 2)
	m3[1] = "mike" //元素的操作
	m3[2] = "go"
	m3[3] = "c++"

	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))

	//键值是唯一的
	m4 := map[int]string{1: "mike", 2: "go", 3: "c++"}
	fmt.Println("m4 = ", m4)

	//赋值，如果已经存在的key值，修改内容
	m4[1] = "c++"
	m4[3] = "go" //追加，map底层自动扩容，和append类似
	fmt.Println("m4 = ", m4)

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2") //删除键值对
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

}
