package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Scan()
	var name string
	var x string
	fmt.Scan(&name)
	fmt.Println(name)

	//fmt.Scanf()
	fmt.Scanf("1:%s", name)
	fmt.Println(name)

	//fmt.Scanf()
	reader := strings.NewReader("1:xxx")
	fmt.Fscanf(reader, "1:%s", &x)
	fmt.Printf("name=%s", x)

}
