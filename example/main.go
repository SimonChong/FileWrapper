package main

import (
	"fmt"

	"github.com/simonchong/wgf/example/resources"
)

//go:generate wgf -i=./resources/sometext.txt -o=./resources/sometext.go -p=resources -c=SomeText
//go:generate wgf -i=./resources/sometext1.txt -o=./resources/sometext1.go -p=resources -c=SomeText1

func main() {

	fmt.Println(resources.SomeText)
	fmt.Println(resources.SomeText1)
}
