package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	githubToken := os.Getenv("GITHUB_TOKEN")
	eventName := os.Getenv("GITHUB_EVENT_NAME")
	fmt.Printf("github token: %s\n", githubToken)
	fmt.Printf("event name: %s\n", eventName)
	payload := getIssuesEvent()
	issueNodeId, err := getIssueNodeId(payload)
	if err != nil {
		log.Printf("[ERROR] %+v\n", errors.Wrap(err, "Failed to extract issue id"))
	}
	fmt.Printf("issue id: %s\n", issueNodeId)
}
