package issues

import (
	"context"

	"github.com/machinebox/graphql"
)

var GITHUB_GRAPHQL_ENDPOINT = "https://api.github.com/graphql"

type IssuesClient interface {
	Execute(projectUrl string, issueAction string, fieldToNewValue map[Field]any, issueNodeId string)
}

type issuesClient struct {
	gqlClient *graphql.Client
	ctx       context.Context
}

func NewIssuesClient() IssuesClient {
	return &issuesClient{
		gqlClient: graphql.NewClient(GITHUB_GRAPHQL_ENDPOINT),
		ctx:       context.Background(),
	}
}

func (issuesClient *issuesClient) Execute(projectUrl string, issueAction string, fieldToNewValue map[Field]any, issueNodeId string) {
	switch IssueAction(issueAction) {
	case Update:
		issuesClient.executeUpdate(projectUrl, fieldToNewValue, issueNodeId)
	}
}
