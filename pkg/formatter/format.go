package formatter

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"gitlab.com/soerenschneider/ssh-login-notification/internal"
	"text/template"
)

const (
	templ        = "New login on {{ .Host }} for {{ .User }} from {{ .Ip }}{{ if or .Dns (or .Geo.Isp .Geo.Org )}} ({{ .PrettyPrintProvider }}){{ end }}{{ if or .Geo.City .Geo.Country .Geo.Region }} {{ .PrettyPrintLocation }}{{ end }}"
	templateName = "defaultTemplate"
)

// Format accepts a struct containing the scraped ip information and
// templates to return the final message string.
func Format(a internal.RemoteUserInfo) string {
	t, err := template.New(templateName).Parse(templ)
	if err != nil {
		log.Panicf("Template is faulty: %v", err)
	}

	output := bytes.Buffer{}
	t.ExecuteTemplate(&output, templateName, &a)

	return output.String()
}
