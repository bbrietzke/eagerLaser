package services

import (
	"os"
	"strings"
	"text/template"
)

type ConsoleNotificationService struct {
	template string
}

func (s *ConsoleNotificationService) Notify(pullRequests []PullRequest) error {
	templater := template.New("Console").Funcs(template.FuncMap{
		"columnTitle": columnTitle,
		"columnState": columnState,
	})

	t1, err := templater.Parse(s.template)

	if err != nil {
		return err
	}

	err = t1.Execute(os.Stdout, pullRequests)

	if err != nil {
		return err
	}

	return nil
}

func NewConsoleNotifier() *ConsoleNotificationService {
	t := `
Title						State      Url
-------------------------------------------------------------------------
{{range . -}}
{{.Title 	| columnTitle }} {{.State | columnState}}{{.Url }}
{{end -}}
`
	return &ConsoleNotificationService{template: t}
}

func columnTitle(text string) string {
	var l = len(text)
	var retVal = []string{}

	if l > 45 {
		var r = strings.Split(text, "")[:40]
		retVal = append(r, "...   ")
	} else {
		var c = 45 - l
		retVal = append(retVal, text)
		for i := 0; i <= c; i++ {
			retVal = append(retVal, " ")
		}
	}

	return strings.Join(retVal, "")
}

func columnState(text string) string {
	var l = len(text)
	var retVal = []string{}

	if l > 10 {
		var r = strings.Split(text, "")[:10]
		retVal = append(r, "...   ")
	} else {
		var c = 10 - l
		retVal = append(retVal, text)
		for i := 0; i <= c; i++ {
			retVal = append(retVal, " ")
		}
	}

	return strings.Join(retVal, "")
}
