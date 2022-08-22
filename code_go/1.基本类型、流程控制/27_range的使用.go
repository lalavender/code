package main //必须有一个main包

import "fmt"

func main() {

	str := "abc"

	//通过for打印每个字符
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	//迭代打印每个元素，默认返回2个值: 一个是元素的位置，一个是元素本身
	for i, data := range str {
		fmt.Printf("str[%d]=%c\n", i, data)
	}

	for i := range str { //第2个返回值，默认丢弃，返回元素的位置(下标)
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	for i, _ := range str { //第2个返回值，默认丢弃，返回元素的位置(下标)
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
