package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

// TODO: fazer um componente para dar handle desta lista
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
type EntriesTemplatable struct {
	De          string   
	Ate         string   
	JobTitle    string   
	Company     string   
	Where       string   
	WhereLink   string   
	Description string 
	Stack       string   
}

func handleExperienceTemplate(resp http.ResponseWriter, req *http.Request) {

	component := getComponent("components/eduexperience-template.html") // The file you want to read
	var entries []Entries
	var templatableEntries []EntriesTemplatable
	
	getEntries("data/experience.json", &entries)

	getTemplatable(entries, &templatableEntries)
	

	var buf bytes.Buffer

	t, err := template.New("experience").Parse(component)
	if err != nil {
		fmt.Println(err)
	}
	

	for _, entry := range templatableEntries {
		
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
	var templatableEntries []EntriesTemplatable
	
	getEntries("data/education.json", &entries)

	getTemplatable(entries, &templatableEntries)
	

	var buf bytes.Buffer

	t, err := template.New("education").Parse(component)
	if err != nil {
		fmt.Println(err)
	}
	

	for _, entry := range templatableEntries {
		
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

func getTemplatable(entries []Entries, to *[]EntriesTemplatable)  {
	description := getComponent("components/description-template.html") // The file you want to read
	t2, err := template.New("description").Parse(description)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range entries {
		var buf2 bytes.Buffer
		for _, val := range entry.Description{
			err := t2.ExecuteTemplate(&buf2, "description", val)
			if err != nil{
				fmt.Println(err)
			}
		}
		*(to) = append(*(to), EntriesTemplatable{
			De: entry.De,
			Ate: entry.Ate,
			JobTitle: entry.JobTitle,
			Company: entry.Company,
			Where: entry.Where,
			WhereLink: entry.WhereLink,
			Stack: entry.Stack,
			Description: buf2.String(),
		})
	}
}