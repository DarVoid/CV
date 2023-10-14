package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleDefault).Methods("GET")
	r.HandleFunc("/header", handleHeader).Methods("GET")
	r.HandleFunc("/header-contacts", handleHeaderContacts).Methods("GET")

	r.HandleFunc("/leftpane", handleLeftPane).Methods("GET")
	r.HandleFunc("/skills", handleSkills).Methods("GET")
	r.HandleFunc("/traits", handleTraits).Methods("GET")
	r.HandleFunc("/faq", handleFAQ).Methods("GET")

	r.HandleFunc("/experience-template", handleExperienceTemplate).Methods("GET")
	r.HandleFunc("/education-template", handleEducationTemplate).Methods("GET")

	r.HandleFunc("/footer", handleFooter).Methods("GET")

	r.HandleFunc("/icons/{icon}", handleIcons).Methods("GET")
	r.HandleFunc("/components/{component}", handleComponent).Methods("GET")
	r.HandleFunc("/pic/{pic}", handlePic).Methods("GET")
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/"))))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}

func handleDefault(resp http.ResponseWriter, req *http.Request) {
	filePath := "pages/index.html" // The file you want to read

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


func handleFooter(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		fmt.Fprintf(resp, `<li>footer</li>`)
	}
}
func handleLeftPane(resp http.ResponseWriter, req *http.Request) {
	filePath := "components/leftpane.html" // The file you want to read

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
