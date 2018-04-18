package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "测试name"
	fmt.Println(os.Args[0])
	fmt.Println(name)

	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		<body>
		<h1>` + name +
		`</h1>
		</body>
		</head>
		</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("创建文件失败", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

}
