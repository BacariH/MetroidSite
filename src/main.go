package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseFiles("./templates/index.html"))
}

// start here and see if you can get the backend connected to this and sen setup htmx frontend
func main() {
	fmt.Println(specificSamusAbility())

}

func specificSamusAbility() string {
	resp, err := http.Get("") // env

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	errorHandler := templ.Execute(os.Stdout, bodyString)
	if errorHandler != nil {
		log.Fatal(errorHandler)
	}

	return bodyString
	// fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Header)
	// fmt.Println(resp.Header["Content-Type"])
	// fmt.Println(resp.Header["Content-Type"][0])
	// fmt.Println(bodyString)
}
