package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Entries struct {
	De          string   `json:"De"`
	Ate         string   `json:"Ate"`
	JobTitle    string   `json:"JobTitle"`
	Company     string   `json:"Company"`
	Where       string   `json:"Where"`
	WhereLink   string   `json:"WhereLink"`
	Description []string `json:"Description"`
	Stack       string   `json:"Stack"`
}

func handleExperienceTemplate(resp http.ResponseWriter, req *http.Request) {

	component := getComponent("components/eduexperience-template.html") // The file you want to read
	var entries []Entries

	getEntries("data/experience.json", &entries)

	var buf bytes.Buffer

	t, err := template.New("experience").Parse(component)
	if err != nil {
		fmt.Println(err)
	}

	for _, entry := range entries {

		err = t.ExecuteTemplate(&buf, "experience", entry)
		if err != nil {
			fmt.Println(err)
		}
	}

	if req.Method == "GET" {
		fmt.Fprintf(resp, "%v", &buf)
	}

}
func handleEducationTemplate(resp http.ResponseWriter, req *http.Request) {
	component := getComponent("components/eduexperience-template.html") // The file you want to read
	var entries []Entries

	getEntries("data/education.json", &entries)

	var buf bytes.Buffer

	t, err := template.New("education").Parse(component)
	if err != nil {
		fmt.Println(err)
	}

	for _, entry := range entries {

		err = t.ExecuteTemplate(&buf, "education", entry)
		if err != nil {
			fmt.Println(err)
		}
	}

	if req.Method == "GET" {
		fmt.Fprintf(resp, "%v", &buf)
	}

}

func getComponent(path string) string {
	experienceTemplate := path // The file you want to read

	contentBytes, err := os.ReadFile(experienceTemplate)
	if err != nil {
		fmt.Println("Error reading the file:", err)
	}
	component := string(contentBytes)
	return component
}

func handleTraits(resp http.ResponseWriter, req *http.Request) {
	component := getComponent("components/traits.html") // The file you want to read
	var traits []string

	getTraits("data/traits.json", &traits)

	var buf bytes.Buffer

	t, err := template.New("traits").Parse(component)
	if err != nil {
		fmt.Println(err)
	}
	err = t.ExecuteTemplate(&buf, "traits", traits)
	if err != nil {
		fmt.Println(err)
	}

	if req.Method == "GET" {
		fmt.Fprintf(resp, "%v", &buf)
	}
}

func getEntries(path string, entries *[]Entries) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&entries)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}
func getTraits(path string, traits *[]string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&traits)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}
