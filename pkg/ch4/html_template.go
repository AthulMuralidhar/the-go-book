package ch4

import (
	"html/template"
	"log"
	"os"
)

func HTMLTemplate() {
	const HTMLTemplate = `
<h1>Total Issues: {{.TotalCount}}  </h1> 
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{ range .Items }}
<tr>
<td>
<a href='{{.HTMLURL}}'>
{{.Number}}
</td>
<td> 
{{.State}}
</td>
<td>
<a href='{{.User.HTMLURL}}'>
{{.User.Login}}
</a></td>
<td>
<a href='{{.HTMLURL}}'>{{.Title}}</a>
</td>
</tr>
{{ end }}
</table>
`
	report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(HTMLTemplate)
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
