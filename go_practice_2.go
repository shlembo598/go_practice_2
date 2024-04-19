package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type PageData struct {
	Title   string
	Content string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	fff := fmt.Sprintf("%s", strings.Join(argsWithoutProg, " "))
	pageData := PageData{Title: "Home Page", Content: fff}

	http.HandleFunc("/homepage", func(writer http.ResponseWriter, request *http.Request) {
		tmpl, err := template.ParseFiles("html_template.html")
		checkError(err)

		err = tmpl.Execute(writer, pageData)
		checkError(err)

	})

	error := http.ListenAndServe(":8080", nil)
	checkError(error)
}
