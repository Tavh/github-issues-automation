package issues

func Execute(githubToken string, projectUrl string, action string, issueFieldsToUdpate map[IssueField]any, issueNodeId string) {
	switch IssueAction(action) {
	case Update:
		executeUpdate(githubToken, projectUrl, issueFieldsToUdpate, issueNodeId)
	}
}
