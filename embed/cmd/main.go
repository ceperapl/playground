package main

import (
	"embed"
	"fmt"
	"example/bundle"
	"net/http"
	"log"
)

//go:embed files/hello.txt

var s string

//go:embed files/file.xml
var b []byte

//go:embed files/*
var staticFiles embed.FS


func main() {
	fmt.Println(s)

	fmt.Println(string(b))

	data, _ := staticFiles.ReadFile("files/example.json")
	fmt.Println(string(data))
	data, _ = staticFiles.ReadFile("go.mod")
	fmt.Println(string(data))

	bundle := bundle.GetBundle()
	fmt.Println(bundle)

	http.Handle("/", http.FileServer(http.FS(staticFiles)))
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
