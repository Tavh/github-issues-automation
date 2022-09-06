package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/tavh/github-issues-automation/logs"
)

func main() {
	githubToken := os.Getenv("GITHUB_TOKEN")
	eventName := os.Getenv("GITHUB_EVENT_NAME")
	logs.Debug("github token: %s\n", githubToken)
	logs.Debug("github event name: %s\n", eventName)
	issueEvent := getIssuesEvent()
	issueNodeId, err := getIssueNodeId(issueEvent)
	if err != nil {
		logs.Error(errors.Wrap(err, "Failed to extract issue id"))
	}
	fmt.Printf("issue id: %s\n", issueNodeId)
}
