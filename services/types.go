package services

import "time"

type PullRequest struct {
	Url       string
	Id        int
	State     string
	Title     string
	CreatedOn time.Time
}

type PRService interface {
	GetPullRequests() (PullRequest, error)
}

type NotificationFormatter interface {
	Notify([]PullRequest) error
}
