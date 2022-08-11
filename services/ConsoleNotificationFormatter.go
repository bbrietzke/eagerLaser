package services

import (
	"os"
	"text/template"
)

type ConsoleNotificationService struct {
	template string
}

func (s *ConsoleNotificationService) Notify(pullRequests []PullRequest) error {
	templater := template.New("Console")

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
Title						State			Url
-------------------------------------------------------------------------
{{range . -}}
{{.Title }}					{{.State }}		{{.Url }}
{{end -}}
`
	return &ConsoleNotificationService{template: t}
}
