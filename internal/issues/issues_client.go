package issues

import (
	"context"

	"github.com/machinebox/graphql"
)

var GITHUB_GRAPHQL_ENDPOINT = "https://api.github.com/graphql"

type IssuesClient interface {
	Execute(issueAction string, fieldToNewValue map[Field]any, issueNodeId string)
}

type issuesClient struct {
	gqlClient     *graphql.Client
	ctx           context.Context
	organization  string
	projectNumber string
}

func NewIssuesClient(organization string, projectNumber string) IssuesClient {
	return &issuesClient{
		gqlClient:     graphql.NewClient(GITHUB_GRAPHQL_ENDPOINT),
		ctx:           context.Background(),
		organization:  organization,
		projectNumber: projectNumber,
	}
}

func (issuesClient *issuesClient) Execute(issueAction string, fieldToNewValue map[Field]any, issueNodeId string) {
	switch IssueAction(issueAction) {
	case Update:
		issuesClient.executeUpdate(fieldToNewValue, issueNodeId)
	}
}
