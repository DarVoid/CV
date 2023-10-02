package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// TODO: fazer um componente para dar handle desta lista
type Entries struct {
	De          string
	Ate         string
	JobTitle    string
	Company     string
	Where       string
	WhereLink   string
	Description []string
	Stack       string
}

func handleExperienceTemplate(resp http.ResponseWriter, req *http.Request) {
	experienceTemplate := "components/experience-template.html" // The file you want to read

	var total []string
	entries := []Entries{
		{
			De:  "03/2022",
			Ate: "Present",
			Description: []string{
				"Developed, tested, and maintained a web App for Infrastructure as Code.",
				"Created and managed GitHub actions and reusable workflows to ensure code quality.",
				"Collaborated with Product Owner and Stakeholders to define project vision and roadmap.",
			},
			// "Assumed the role of scrum master (scrum knight) and developer, enabling the delivery of value and ensuring scrum events.",
			JobTitle:  "Fullstack Developer",
			Company:   "Critical Techworks",
			Where:     "Porto",
			WhereLink: "https://goo.gl/maps/pBKxevEVCuREvH3y8",
			Stack:     "GraphQL, Golang, NodeJS, React, Python, AWS + Terraform, DynamoDB, Mongo, Kubernetes, React, Docker, Linux, GitHub",
		},
		{
			De:  "10/2022",
			Ate: "02/2023",
			Description: []string{
				"Developed event-driven micro-services to support a loan Web app for Portugal and Spain markets.",
			},
			JobTitle:  "Backend Developer",
			Company:   "CapGemini",
			Where:     "Lisbon",
			WhereLink: "https://bit.ly/3DUpYy6",
			Stack:     "Golang, Postgres, RabbitMQ, Docker, GCP, Kubernetes, Travis",
		},
		{
			De:  "04/2021",
			Ate: "10/2022",
			Description: []string{
				"Planned and developed an app to manage Industrial Property.",
				"Built CI/CD pipelines and containers for application deployment.",
			},
			JobTitle:  "Fullstack Developer",
			Company:   "Techframe, SA",
			Where:     "Abrantes",
			WhereLink: "https://bit.ly/3n53FNo",
			Stack:     "Angular, Docker, Golang",
		},
		{
			De:  "07/2020",
			Ate: "03/2021",
			Description: []string{
				"Created ETL scripts and refactored APIs for internal frontend applications.",
			},
			JobTitle:  "Fullstack Developer",
			Company:   "Softinsa, IBM",
			Where:     "Tomar",
			WhereLink: "https://bit.ly/35vZgeX",
			Stack:     "T-SQL, Node.js, Python",
		},
	}

	contentBytes, err := os.ReadFile(experienceTemplate)
	if err != nil {
		fmt.Println("Error reading the file:", err)
	}
	formatString := string(contentBytes)
	for _, entry := range entries {
		total = append(total,
			fmt.Sprintf(formatString,
				entry.De,
				entry.Ate,
				entry.JobTitle,
				entry.Company,
				entry.WhereLink,
				entry.Where,
				entry.Description,
				entry.Stack))
	}
	if req.Method == "GET" {
		fmt.Fprintf(resp, "%v", strings.Join(total, "\n"))
	}

}
func handleEducationTemplate(resp http.ResponseWriter, req *http.Request) {
	experienceTemplate := "components/education-template.html" // The file you want to read

	var total []string
	entries := []Entries{
		{
			De:          "2017",
			Ate:         "2020",
			Description: []string{"cenas"},
			JobTitle:    "Computer Engineering",
			Company:     "Bachelor's Degree",
			Where:       "IPT",
			WhereLink:   "https://bit.ly/2IsX53k",
			Stack:       "Ruby on rails,Angular, Travis, Docker, Vagrant, MatLab, Python, Java, CentOS, etc...",
		},
	}

	contentBytes, err := os.ReadFile(experienceTemplate)
	if err != nil {
		fmt.Println("Error reading the file:", err)
	}
	formatString := string(contentBytes)
	for _, entry := range entries {
		total = append(total,
			fmt.Sprintf(formatString,
				entry.De,
				entry.Ate,
				entry.JobTitle,
				entry.Company,
				entry.WhereLink,
				entry.Where,
				entry.Description,
				entry.Stack))
	}

	if req.Method == "GET" {
		fmt.Fprintf(resp, "%v", strings.Join(total, "\n"))
	}

}