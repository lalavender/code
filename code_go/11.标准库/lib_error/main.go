package main

import (
	"errors"
	"fmt"
)

func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空！")
		return "", err
	} else {
		return s, nil
	}

}

func main() {
	s, err := check("123")
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	} else {
		fmt.Println("s:", s)
	}
}
