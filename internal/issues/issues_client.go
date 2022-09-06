package issues

import (
	"github.com/machinebox/graphql"
	"github.com/tavh/github-issues-automation/internal/logs"

	"context"
)

var GITHUB_GRAPHQL_ENDPOINT = "https://api.github.com/graphql"

type IssuesClient interface {
	Execute(projectUrl string, issueAction string, issueFieldsToUdpate map[IssueField]any, issueNodeId string)
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

func (issuesClient *issuesClient) Execute(projectUrl string, issueAction string, issueFieldsToUdpate map[IssueField]any, issueNodeId string) {
	switch IssueAction(issueAction) {
	case Update:
		issuesClient.executeUpdate(projectUrl, issueFieldsToUdpate, issueNodeId)
	}
}

func (issuesClient *issuesClient) executeUpdate(projectUrl string, issueFieldsToNewValues map[IssueField]any, issueNodeId string) {
	err := validateFields(issueFieldsToNewValues)
	if validateFields(issueFieldsToNewValues) != nil {
		logs.Error(err)
		return
	}
}
