package formatter

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"sshnot/internal"
	"text/template"
)

var templ = "New login on {{ .Host }} for {{ .User }} from {{ .Ip }}{{ if or .Dns (or .IpInfo.Isp .IpInfo.Org )}} ({{ .PrettyPrintProvider }}){{ end }} {{ if .IpInfo }} {{ .PrettyPrintLocation }}{{ end }}"

func Format(a internal.SshLoginNotification) string {
	templateName := "default"
	t, err := template.New(templateName).Parse(templ)
	if err != nil {
		log.Panicf("Template is faulty: %v", err)
	}

	output := bytes.Buffer{}
	t.ExecuteTemplate(&output, templateName,  &a)

	return output.String()
}
