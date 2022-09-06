package issues

func Execute(githubToken string, projectUrl string, issueAction string, issueFieldsToUdpate map[IssueField]any, issueNodeId string) {
	switch IssueAction(issueAction) {
	case Update:
		executeUpdate(githubToken, projectUrl, issueFieldsToUdpate, issueNodeId)
	}
}
