package ch4

import (
	"log"
	"os"
	"text/template"
	"time"
)

func TextTemplate() {
	const templateString = `
Total Issues: {{.TotalCount}}  {{range .Items}}------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{ .CreatedAt | daysAgo }} days
{{ end }}
`
	report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}

	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
