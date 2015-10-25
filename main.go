package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

var (
	inputFile  string
	outputFile string
	pack       string
)

type configType struct {
	InputFile    string
	OutputFile   string
	PackageName  string
	ConstantName string
	InputFileTxt string
}

var config = configType{}

const tpl = `package {{.PackageName}}

const {{.ConstantName}} = ` + "`{{.InputFileTxt}}`" + `

`

func init() {
	flag.StringVar(&config.InputFile, "i", "", "INPUT: The file to convert to go source")
	flag.StringVar(&config.OutputFile, "o", "", "OUTPUT: The go output file")
	flag.StringVar(&config.PackageName, "p", "", "Package: The package for the go file")
	flag.StringVar(&config.ConstantName, "c", "", "Constant: The name of the constant")
	flag.Parse()

	if len(os.Args) < 4 {
		fmt.Println("Please enter the following flags")
		flag.Usage()
		os.Exit(1)
	}
}

func main() {

	log.Println(config.InputFile)
	log.Println(config.OutputFile)
	t := template.Must(template.New("tpl").Parse(tpl))

	in, err := ioutil.ReadFile(config.InputFile)
	config.InputFileTxt = strings.Replace(string(in), "`", "` + \"`", -1)
	checkErr(err)

	f, err := os.Create(config.OutputFile)
	checkErr(err)
	defer f.Close()

	t.Execute(f, config)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
