package issues

import (
	"github.com/machinebox/graphql"
	"github.com/tavh/github-issues-automation/internal/logs"
)

type IssuesClient struct {
	GQLClient *graphql.Client
}

func (issuesClient *IssuesClient) Execute(projectUrl string, issueAction string, issueFieldsToUdpate map[IssueField]any, issueNodeId string) {
	switch IssueAction(issueAction) {
	case Update:
		executeUpdate(projectUrl, issueFieldsToUdpate, issueNodeId)
	}
}

func executeUpdate(projectUrl string, issueFieldsToNewValues map[IssueField]any, issueNodeId string) {
	err := validateFields(issueFieldsToNewValues)
	if validateFields(issueFieldsToNewValues) != nil {
		logs.Error(err)
		return
	}
}
