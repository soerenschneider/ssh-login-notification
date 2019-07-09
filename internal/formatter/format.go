package formatter

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"sshnot/internal"
	"text/template"
)

const (
	templ = "New login on {{ .Host }} for {{ .User }} from {{ .Ip }}{{ if or .Dns (or .Geo.Isp .Geo.Org )}} ({{ .PrettyPrintProvider }}){{ end }}{{ if or .Geo.City .Geo.Country .Geo.Region }} {{ .PrettyPrintLocation }}{{ end }}"
	templateName = "defaultTemplate"
)

func Format(a internal.SshLoginNotification) string {
	t, err := template.New(templateName).Parse(templ)
	if err != nil {
		log.Panicf("Template is faulty: %v", err)
	}

	output := bytes.Buffer{}
	t.ExecuteTemplate(&output, templateName,  &a)

	return output.String()
}
