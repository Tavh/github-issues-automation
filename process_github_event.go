package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/google/go-github/v47/github"
	"github.com/pkg/errors"

	"github.com/tavh/github-issues-automation/logs"
)

func getIssuesEvent() github.IssuesEvent {
	logs.Debug("github event path: %s\n", os.Getenv("GITHUB_EVENT_PATH"))

	var jsonFilePath string = os.Getenv("GITHUB_EVENT_PATH")

	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		logs.Error(errors.Wrap(err, "Failed to open json"))
	}
	defer jsonFile.Close()

	jsonByte, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logs.Error(errors.Wrap(err, "Failed to read json as a byte array"))
	}

	issueEvent := github.IssuesEvent{}
	err = json.Unmarshal(jsonByte, &issueEvent)
	if err != nil {
		logs.Error(errors.Wrap(err, "Failed to unmarshal JSON to Go Object"))
	}

	logs.Debug("github issue event: %v\n", issueEvent)
	return issueEvent
}

func getIssueNodeId(issueEvent github.IssuesEvent) (string, error) {
	issue := issueEvent.GetIssue()
	logs.Debug("issue: %s\n", issue)
	issueNodeId := issueEvent.GetIssue().GetNodeID()
	logs.Debug("issue nodeId: %s\n", issueNodeId)

	if issueNodeId == "" {
		return "", errors.New("Issue ID is \"\". Failed to get issue id properly")
	}

	return issueNodeId, nil
}
