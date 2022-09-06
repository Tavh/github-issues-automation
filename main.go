package main

import (
	"os"

	"github.com/machinebox/graphql"
	"github.com/pkg/errors"
	"github.com/tavh/github-issues-automation/internal/issues"
	"github.com/tavh/github-issues-automation/internal/logs"
)

var ISSUES_EVENT_NAME = "issues"
var GITHUB_GRAPHQL_ENDPOINT = "https://api.github.com/graphql"

func main() {
	logs.Init()

	eventName := os.Getenv("GITHUB_EVENT_NAME")
	logs.Debug("github event name: %s\n", eventName)

	if eventName == ISSUES_EVENT_NAME {
		githubToken := os.Getenv("GITHUB_TOKEN")
		logs.Debug("github token: %s\n", githubToken)

		projectUrl := os.Getenv("PROJECT_URL")
		logs.Debug("project url: %s\n", projectUrl)

		targetStatus := os.Getenv("TARGET_STATUS")
		logs.Debug("target-status: %s\n", targetStatus)

		issueAction := os.Getenv("ISSUE_ACTION")
		logs.Debug("issue action: %s\n", issueAction)

		issueNodeId, err := GetIssueNodeId()
		if err != nil {
			logs.Error(errors.Wrap(err, "Failed to extract issue id"))
		}
		logs.Debug("issue id: %s\n", issueNodeId)

		client := issues.IssuesClient{
			GQLClient: graphql.NewClient(GITHUB_GRAPHQL_ENDPOINT),
		}

		client.Execute(projectUrl, issueAction, constructIssueFieldsToNewValues(targetStatus), issueNodeId)
	} else {
		logs.Error(errors.Errorf("Action triggered from invalid event, only supports events of type '%s'\n", ISSUES_EVENT_NAME))
	}
}

func constructIssueFieldsToNewValues(status string) map[issues.IssueField]any {
	return map[issues.IssueField]any{
		issues.Status: status,
	}
}
