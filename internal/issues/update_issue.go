package issues

import (
	"github.com/pkg/errors"
)

type IssueField string

const (
	Status IssueField = "status"
)

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
