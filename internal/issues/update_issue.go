package issues

import (
	"github.com/pkg/errors"
	"github.com/tavh/github-issues-automation/internal/logs"
)

type IssueField string

const (
	Status IssueField = "status"
)

func executeUpdate(githubToken string, projectUrl string, issueFieldsToNewValues map[IssueField]any, issueNodeId string) {
	err := validateFields(issueFieldsToNewValues)
	if validateFields(issueFieldsToNewValues) != nil {
		logs.Error(err)
		return
	}
}

func validateFields(issueFieldsToNewValues map[IssueField]any) error {
	isAnyFieldValuePresent := false

	if issueFieldsToNewValues[Status] != nil {
		isAnyFieldValuePresent = true
	}

	if !isAnyFieldValuePresent {
		return errors.Errorf("Action %s was requested but no new field values were provided\n", Update)
	}

	return nil
}
