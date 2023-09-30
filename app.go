package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	BASEURL = "http://localhost:8080"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleDefault).Methods("GET")
	r.HandleFunc("/header", handleHeader).Methods("GET")
	r.HandleFunc("/footer", handleFooter).Methods("GET")
	r.HandleFunc("/icons/{icon}", handleIcons).Methods("GET")
	r.HandleFunc("/content", handleMasterContent).Methods("GET")
	r.HandleFunc("/header-contacts", handleHeaderContacts).Methods("GET")
	r.HandleFunc("/skills", handleSkills).Methods("GET")
	r.HandleFunc("/leftpane", handleLeftPane).Methods("GET")

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":8080", r)
}

func handleDefault(resp http.ResponseWriter, req *http.Request) {
	filePath := "components/index.html" // The file you want to read

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
func handleHeader(resp http.ResponseWriter, req *http.Request) {
	filePath := "components/header.html" // The file you want to read

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
func handleHeaderContacts(resp http.ResponseWriter, req *http.Request) {
	filePath := "components/contacts.html" // The file you want to read

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
func handleSkills(resp http.ResponseWriter, req *http.Request) {
	filePath := "components/skills.html" // The file you want to read

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
func handleMasterContent(resp http.ResponseWriter, req *http.Request) {
	filePath := "components/master-content.html" // The file you want to read

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
func handleFooter(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		fmt.Fprintf(resp, `<li>footer</li>`)
	}
}
func handleLeftPane(resp http.ResponseWriter, req *http.Request) {
	filePath := "components/master-left-pane.html" // The file you want to read

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

func handleLinuxIcon(resp http.ResponseWriter, req *http.Request) {
	filePath := "icons/linux-icon.svg" // The file you want to read

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
