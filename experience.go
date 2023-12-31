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

type Skills struct {
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Title string `json:"title"`
}

type Language struct{
	Language string `json:"language"`
	Proficiency string `json:"proficiency"`
}

type Motivation struct{
	What string `json:"what"`
	Icon string `json:"icon"`
	Explanation string `json:"explanation"`
	Video string `json:"video"`
}

type Faq struct{
	Languages []Language
	Motivations []Motivation
	Exploring []Skills
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

func handleSkills(resp http.ResponseWriter, req *http.Request) {

	component := getComponent("components/skills.html") // The file you want to read
	var skills []Skills

	getSkills("data/skills.json", &skills)

	var buf bytes.Buffer

	t, err := template.New("skills").Parse(component)
	if err != nil {
		fmt.Println(err)
	}
	err = t.ExecuteTemplate(&buf, "skills", skills)
	if err != nil {
		fmt.Println(err)
	}

	if req.Method == "GET" {
		fmt.Fprintf(resp, "%v", &buf)
	}
}


func handleFAQ(resp http.ResponseWriter, req *http.Request) {
	component := getComponent("components/faq.html") // The file you want to read
	
	var languages []Language
	getLanguages("data/languages.json", &languages)
	var motivations []Motivation
	getMotivations("data/motivation.json", &motivations)
	var exploring []Skills
	getSkills("data/exploring.json", &exploring)
	faqs := Faq{
		Languages: languages,
		Motivations: motivations,
		Exploring: exploring,
	}

	var buf bytes.Buffer
	
	t, err := template.New("faq").Parse(component)
	if err != nil {
		fmt.Println(err)
	}
	err = t.ExecuteTemplate(&buf, "faq", faqs)
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
func getSkills(path string, skills *[]Skills) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&skills)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}
func getLanguages(path string, skills *[]Language) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&skills)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}
func getMotivations(path string, skills *[]Motivation) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&skills)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

