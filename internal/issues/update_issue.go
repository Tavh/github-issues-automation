package issues

import (
	"github.com/machinebox/graphql"
	"github.com/pkg/errors"

	"github.com/tavh/github-issues-automation/internal/logs"
)

type Field string

const (
	Status Field = "status"
	Label  Field = "label"
)

func (issuesClient *issuesClient) executeUpdate(fieldToNewValue map[Field]any, issueNodeId string) error {
	anyIssueLevelFieldsPresent := issuesClient.handleIssueLevelFieldsIfPresent(fieldToNewValue, issueNodeId)
	anyItemLevelFieldsPresent := issuesClient.handleItemLevelFieldsIfPresent(fieldToNewValue, issueNodeId)

	if !anyIssueLevelFieldsPresent && !anyItemLevelFieldsPresent {
		return errors.Errorf("Action %s was requested but no new field values were provided\n", Update)
	}

	return nil
}

func (issuesClient *issuesClient) handleItemLevelFieldsIfPresent(fieldToNewValue map[Field]any, issueNodeId string) bool {
	isAnyFieldValuePresent := false

	statusValue := fieldToNewValue[Status]
	if statusValue != nil {
		isAnyFieldValuePresent = true
		issuesClient.updateItemLevelField(Status, statusValue, issueNodeId)
	}

	return isAnyFieldValuePresent
}

func (issuesClient *issuesClient) handleIssueLevelFieldsIfPresent(fieldToNewValue map[Field]any, issueNodeId string) bool {
	isAnyFieldValuePresent := false

	if fieldToNewValue[Label] != nil {
		isAnyFieldValuePresent = true
	}

	return isAnyFieldValuePresent
}

func (issuesClient *issuesClient) updateItemLevelField(field Field, fieldNewValue any, issueNodeId string) {
	req := graphql.NewRequest(
		`query($organization: String!, $projectNumber: Int!) {
			organization(login: $organization){
				projectV2(number: $projectNumber) {
					id
					fields(first:100) {
					nodes {
						... on ProjectV2SingleSelectField {
							id
							name                        
							options {
								id
								name
							}
						}
					}
				}
			}
		}`,
	)
	req.Var("organization", issuesClient.organization)
	req.Var("projectNumber", issuesClient.projectNumber)

	var res map[string]any

	logs.Debug("gql request: %v", req)
	issuesClient.gqlClient.Run(issuesClient.ctx, req, &res)

	logs.Debug("gql response: %v", res)
}
