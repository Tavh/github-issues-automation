package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/google/go-github/v47/github"
	"github.com/pkg/errors"
)

func getIssuesEvent() github.IssuesEvent {
	var jsonFilePath string
	_, ok := os.LookupEnv("GITHUB_ACTION_LOCAL")
	if ok {
		var err error
		jsonFilePath, err = filepath.Abs("./payload/issue_event.json")
		if err != nil {
			log.Printf("[ERROR] %+v\n", err)
		}
	} else {
		jsonFilePath = os.Getenv("GITHUB_EVENT_PATH")
	}
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		log.Printf("[ERROR] %+v\n", errors.Wrap(err, "Failed to open json"))
	}
	defer jsonFile.Close()

	// read opened jsonFile as a byte array.
	jsonByte, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Printf("[ERROR] %+v\n", errors.Wrap(err, "Failed to read json as a byte array"))
	}

	payload := github.IssuesEvent{}
	err = json.Unmarshal(jsonByte, &payload)
	if err != nil {
		log.Printf("[ERROR] %+v\n", errors.Wrap(err, "Failed to unmarshal JSON to Go Object"))
	}

	return payload
}

func getIssueNodeId(event github.IssuesEvent) (string, error) {
	issueNodeId := event.GetIssue().GetNodeID()

	if issueNodeId == "" {
		return "", errors.New("Issue ID is \"\". Failed to get issue id properly")
	}

	return issueNodeId, nil
}
