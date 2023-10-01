package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

const (
	BASEURL = "http://localhost:8080"
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

	r.HandleFunc("/content", handleMasterContent).Methods("GET")
	r.HandleFunc("/experience-template", handleExperienceTemplate).Methods("GET")
	r.HandleFunc("/education-template", handleEducationTemplate).Methods("GET")

	r.HandleFunc("/footer", handleFooter).Methods("GET")

	r.HandleFunc("/icons/{icon}", handleIcons).Methods("GET")
	r.HandleFunc("/pic/{pic}", handlePic).Methods("GET")
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

type Entries struct {
	De          string
	Ate         string
	JobTitle    string
	Company     string
	Where       string
	WhereLink   string
	Description []string //TODO: fazer um componente para dar handle desta lista
	Stack       string
}

func handleExperienceTemplate(resp http.ResponseWriter, req *http.Request) {
	experienceTemplate := "components/experience-template.html" // The file you want to read
	// exeperienceTemplate := "components/descriptions.html"        // The file you want to read
	parentComponent := "components/table.html" // The file you want to read

	var total []string
	entries := []Entries{
		{
			De:  "3/2022",
			Ate: "Present",
			Description: []string{
				"Developed, tested, and maintained a web App for Infrastructure as Code.",
				"Created and managed GitHub actions and reusable workflows to ensure code quality.",
				"Collaborated with Product Owner and Stakeholders to define project vision and roadmap.",
				// "Assumed the role of scrum master (scrum knight) and developer, enabling the delivery of value and ensuring scrum events.",
			},
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
	contentBytes, err = os.ReadFile(parentComponent)
	if err != nil {
		fmt.Println("Error reading the file:", err)
	}
	formatString = string(contentBytes)
	resposta := fmt.Sprintf(formatString, "Experience", strings.Join(total, "\n"))

	if req.Method == "GET" {
		fmt.Fprintf(resp, "%v", resposta)
	}

}
func handleEducationTemplate(resp http.ResponseWriter, req *http.Request) {
	experienceTemplate := "components/education-template.html" // The file you want to read
	// exeperienceTemplate := "components/descriptions.html"        // The file you want to read
	parentComponent := "components/table.html" // The file you want to read

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
			Stack:       "GraphQL, Golang, NodeJS, React, Python, AWS + Terraform, DynamoDB, Mongo, Kubernetes, React, Docker, Linux, GitHub",
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
	contentBytes, err = os.ReadFile(parentComponent)
	if err != nil {
		fmt.Println("Error reading the file:", err)
	}
	formatString = string(contentBytes)
	resposta := fmt.Sprintf(formatString, "Education", strings.Join(total, "\n"))

	if req.Method == "GET" {
		fmt.Fprintf(resp, "%v", resposta)
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
func handleTraits(resp http.ResponseWriter, req *http.Request) {
	filePath := "components/traits.html" // The file you want to read

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
func handleFAQ(resp http.ResponseWriter, req *http.Request) {
	filePath := "components/faq.html" // The file you want to read

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
