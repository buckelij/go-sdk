package models
// The default value for a merge commit title.- `PR_TITLE` - default to the pull request's title.- `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., "Merge pull request #123 from branch-name").
type PullRequestWebhook_merge_commit_title int

const (
    PR_TITLE_PULLREQUESTWEBHOOK_MERGE_COMMIT_TITLE PullRequestWebhook_merge_commit_title = iota
    MERGE_MESSAGE_PULLREQUESTWEBHOOK_MERGE_COMMIT_TITLE
)

func (i PullRequestWebhook_merge_commit_title) String() string {
    return []string{"PR_TITLE", "MERGE_MESSAGE"}[i]
}
func ParsePullRequestWebhook_merge_commit_title(v string) (any, error) {
    result := PR_TITLE_PULLREQUESTWEBHOOK_MERGE_COMMIT_TITLE
    switch v {
        case "PR_TITLE":
            result = PR_TITLE_PULLREQUESTWEBHOOK_MERGE_COMMIT_TITLE
        case "MERGE_MESSAGE":
            result = MERGE_MESSAGE_PULLREQUESTWEBHOOK_MERGE_COMMIT_TITLE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePullRequestWebhook_merge_commit_title(values []PullRequestWebhook_merge_commit_title) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PullRequestWebhook_merge_commit_title) isMultiValue() bool {
    return false
}
