package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Nikola Obradovic"
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello Go!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)

	f, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer f.Close()

	io.Copy(f, strings.NewReader(str))
}

// output directly to file
