package services

import (
	"os"
	"text/template"
)

type notification struct {
	OpenPrs     []PullRequest
	ClosedPrs   []PullRequest
	Destination string
	Sender      string
	Subject     string
}

type EmailNotificationFormatter struct {
	template    string
	destination string
	sender      string
}

func (s *EmailNotificationFormatter) Notify(pullRequests []PullRequest) error {
	templater := template.New("Console")
	var open []PullRequest
	var closed []PullRequest

	t1, err := templater.Parse(s.template)

	if err != nil {
		return err
	}

	for _, pr := range pullRequests {
		if pr.State == "open" {
			open = append(open, pr)
		} else {
			closed = append(closed, pr)
		}
	}

	data := notification{
		Destination: s.destination,
		Sender:      s.sender,
		OpenPrs:     open,
		ClosedPrs:   closed,
		Subject:     "Pull Requests Log",
	}

	err = t1.Execute(os.Stdout, data)

	if err != nil {
		return err
	}

	return nil
}

func NewEmailNotificationFormatter(to string, from string) *EmailNotificationFormatter {
	t := `
<h1>[To] {{.Destination}}</h1>
<h1>[From] {{.Sender}}</h1>
<h1>[Subject] {{.Subject}}</h1>

<h3>Open Pull Requests</h1>
<table>
	<thead>
		<tr>
			<th>Title</th>
			<th>Status</th>
			<th>Created On</th>
		</th>
	</thead>
	<tbody>
{{ range .OpenPrs }}		<tr>
			<td><a href="{{ .Url -}}">{{ .Title -}}</a></td>
			<td>{{ .State -}}</td>
			<td>{{ .CreatedOn -}}</td>
		</tr>
{{ end }}	</tbody>
</table>

<h3>Closed Pull Requests</h1>
<table>
	<thead>
		<tr>
			<th>Title</th>
			<th>Status</th>
			<th>Created On</th>
		</th>
	</thead>
	<tbody>
{{ range .ClosedPrs }}		<tr>
			<td><a href="{{ .Url -}}">{{ .Title -}}</a></td>
			<td>{{ .State -}}</td>
			<td>{{ .CreatedOn -}}</td>
		</tr>
{{ end }}	</tbody>
</table>

`
	return &EmailNotificationFormatter{template: t, destination: to, sender: from}
}
