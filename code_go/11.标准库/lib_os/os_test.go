package lib_os

import (
	"fmt"
	"os"
	"testing"
)

func TestCreat(t *testing.T) {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func TestMkdir(t *testing.T) {
	err := os.Mkdir("test", 0666)
	if err != nil {
		panic(err)
	}
}
func TestMkdirAll(t *testing.T) {
	err := os.MkdirAll("test/1/2/3", 0666)
	if err != nil {
		panic(err)
	}
}
func TestRemove(t *testing.T) {
	err := os.Remove("test/1/2/3")
	if err != nil {
		panic(err)
	}
}
func TestRemoveAll(t *testing.T) {
	err := os.RemoveAll("test")
	if err != nil {
		panic(err)
	}
}

func TestGetwd(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
}

func TestChDir(t *testing.T) {
	err := os.Chdir("D:/tmp")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(os.Getwd())
}
func TestTempDir(t *testing.T) {
	s := os.TempDir()
	fmt.Println(s)
}

func TestRename(t *testing.T) {
	err := os.Rename("test.txt", "test1.txt")
	if err != nil {
		fmt.Println(err)
	}
}

//os.Chmod()
//os.Chown()

//读取文件
//os.Stat()
//os.Read()
//os.ReadAt()
//os.ReadDir()
//os.Seek()
