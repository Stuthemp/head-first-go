package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func reportPanic() {
	p := recover()
	if p == nil {
		fmt.Println("No panic happened")
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func checkDir(dirName string, tabs string) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println(tabs+"Directory: ", file.Name())
			checkDir(dirName+"/"+file.Name(), tabs+"\t")
		} else {
			fmt.Println(tabs+"File:", file.Name())
		}
	}
}

func main() {
	dirName := strings.TrimSpace(os.Args[1])
	defer reportPanic()
	checkDir(dirName, "")
}
