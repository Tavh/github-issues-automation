package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/tavh/github-issues-automation/internal/issues"
	"github.com/tavh/github-issues-automation/internal/logs"
)

var ISSUES_EVENT_NAME = "issues"

func main() {
	logs.Init()

	eventName := os.Getenv("GITHUB_EVENT_NAME")
	logs.Debug("github event name: %s\n", eventName)

	if eventName == ISSUES_EVENT_NAME {
		githubToken := os.Getenv("GITHUB_TOKEN")
		logs.Debug("github token: %s\n", githubToken)

		organization := os.Getenv("ORGANIZATION")
		logs.Debug("organization: %s\n", organization)

		projectNumber := os.Getenv("PROJECT_NUMBER")
		logs.Debug("project number: %s\n", projectNumber)

		targetStatus := os.Getenv("TARGET_STATUS")
		logs.Debug("target-status: %s\n", targetStatus)

		issueAction := os.Getenv("ISSUE_ACTION")
		logs.Debug("issue action: %s\n", issueAction)

		issueNodeId, err := GetIssueNodeId()
		if err != nil {
			logs.Error(errors.Wrap(err, "Failed to extract issue id"))
		}
		logs.Debug("issue id: %s\n", issueNodeId)

		client := issues.NewIssuesClient(githubToken, organization, projectNumber)
		client.Execute(issueAction, constructFieldToNewValueMap(targetStatus), issueNodeId)
	} else {
		logs.Error(errors.Errorf("Action triggered from invalid event, only supports events of type '%s'\n", ISSUES_EVENT_NAME))
	}
}

func constructFieldToNewValueMap(status string) map[issues.Field]any {
	return map[issues.Field]any{
		issues.Status: status,
	}
}
