# WGF
Wrap in a Go File.

Creates Go source files by wrapping the contents of files in Go files and exposing them as constants or variable

## Purpose
This program was created to be used by the "go generate" command so that its contents can be included in the compiled source code. This helps encapsulate all functionality in a compiled go executable

## Install

```
go get -u github.com/SimonChong/wgf
```


## Example

wgf is meant to be used with the ```go generate``` command as follows. This example can be found in the ```example``` directory of this repo.
```go
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


```

To generate the file sometext.go run ```go genereate``` in the main package directory.


## Usage
Command line usage:
```
-c string
    Constant: The name of the constant
-i string
    INPUT: The file to convert to go source
-o string
    OUTPUT: The go output file
-p string
    Package: The package for the go file
```
