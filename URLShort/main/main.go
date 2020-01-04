package main

import (
	"flag"
	"fmt"
	"github.com/divbhasin/gophercises/URLShort"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	var yamlFile string
	flag.StringVar(&yamlFile, "file", "redirects.yml",
		"Name of file that contains the redirects")
	flag.Parse()

	redirectAbs, _ := filepath.Abs("URLShort/" + yamlFile)
	yamlFile = redirectAbs

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := URLShort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Fatal("Could not read redirects file: ", err)
	}

	yamlHandler, err := URLShort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
