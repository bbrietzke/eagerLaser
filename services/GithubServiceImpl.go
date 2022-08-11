package services

import (
	"context"
	"github.com/google/go-github/v45/github"
	"time"
)

type GithubPRService struct {
	project string
	owner   string
}

func (p *GithubPRService) GetPullRequests() ([]PullRequest, error) {
	ctx := context.TODO()
	var retVal []PullRequest
	age := time.Now().Add(-360 * time.Hour)

	client := github.NewClient(nil)
	options := &github.PullRequestListOptions{State: "all", Sort: "created", Direction: "desc"}

	prs, _, err := client.PullRequests.List(ctx, p.owner, p.project, options)

	if err != nil {
		return nil, err
	}

	for _, pr := range prs {

		if pr.CreatedAt.After(age) {
			item := PullRequest{
				Url:       pr.GetURL(),
				State:     *pr.State,
				Title:     *pr.Title,
				CreatedOn: *pr.CreatedAt,
			}

			retVal = append(retVal, item)
		}
	}

	return retVal, nil
}

// GithubPullRequests
func GithubPullRequests(ownerName string, projectName string) *GithubPRService {
	return &GithubPRService{project: projectName, owner: ownerName}
}
