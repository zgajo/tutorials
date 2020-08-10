package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	name := "Darko Pranjić"
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` +
		name +
		`</h1>
		</body>
		</html>
	`)

	nf, err := os.Create("003_string-to-html/02_file/index.html")

	if err != nil {
		println("Error", err)
		os.Exit(1)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

}
