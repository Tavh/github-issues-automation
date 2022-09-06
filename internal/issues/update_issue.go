package issues

import (
	"github.com/pkg/errors"
)

type Field string

const (
	Status Field = "status"
	Label  Field = "label"
)

func (issuesClient *issuesClient) executeUpdate(projectUrl string, fieldToNewValue map[Field]any, issueNodeId string) error {
	anyIssueLevelFieldsPresent := handleIssueLevelFieldsIfPresent(fieldToNewValue)
	anyItemLevelFieldsPresent := handleItemLevelFieldsIfPresent(fieldToNewValue)

	if !anyIssueLevelFieldsPresent && !anyItemLevelFieldsPresent {
		return errors.Errorf("Action %s was requested but no new field values were provided\n", Update)
	}

	return nil
}

func handleItemLevelFieldsIfPresent(fieldToNewValue map[Field]any) bool {
	isAnyFieldValuePresent := false

	if fieldToNewValue[Status] != nil {
		isAnyFieldValuePresent = true
	}

	return isAnyFieldValuePresent
}

func handleIssueLevelFieldsIfPresent(fieldToNewValue map[Field]any) bool {
	isAnyFieldValuePresent := false

	if fieldToNewValue[Label] != nil {
		isAnyFieldValuePresent = true
	}

	return isAnyFieldValuePresent
}

func updateItemLevelField(Field) {
	// TODO: Implement actual update
}
