package main

import (
	"os"
	"io"
	"log"
	"html/template"
	"fmt"
	"net/http"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
)

func fileHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Fatal("error getting a file for the provided form key: ", err)
		return
	}
	defer file.Close()
	out, pathError := os.Create("tmp/uploadedFile.png")
	if pathError != nil {
		log.Fatal("error creating a file for writing: ", pathError)
		return
	}
	defer out.Close()
	_, copyFileError := io.Copy(out, file)
	if copyFileError != nil {
		log.Fatal("error occurred while file copy: ", copyFileError)
		return
	}
	fmt.Fprintf(w, "File uploaded successfully: " + header.Filename)
}

func index(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/upload-file.html")
	parsedTemplate.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", fileHandler)

	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, nil)
	if err != nil {
		log.Fatal("Error starting http server: ", err)
		return
	}
}