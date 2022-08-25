package main

import (
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	log.Println("this is a log ")
	log.Panicln("this is a log ")
}
