package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handleIcons(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	icon := vars["icon"]
	filePath := fmt.Sprintf("icons/%v-icon.svg", icon) // The file you want to read

	// Read the contents of the file
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
	}
	content := string(contentBytes)

	if req.Method == "GET" {
		fmt.Fprintf(resp, "%v", content)
	}
}
func handlePic(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	pic := vars["pic"]
	filePath := fmt.Sprintf("components/pic%v.html", pic) // The file you want to read

	// Read the contents of the file
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
	}
	content := string(contentBytes)

	if req.Method == "GET" {
		fmt.Fprintf(resp, "%v", content)
	}
}
