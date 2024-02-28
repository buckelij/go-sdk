package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepoSearchResultItem repo Search Result Item
type RepoSearchResultItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allow_auto_merge property
    allow_auto_merge *bool
    // The allow_forking property
    allow_forking *bool
    // The allow_merge_commit property
    allow_merge_commit *bool
    // The allow_rebase_merge property
    allow_rebase_merge *bool
    // The allow_squash_merge property
    allow_squash_merge *bool
    // The archive_url property
    archive_url *string
    // The archived property
    archived *bool
    // The assignees_url property
    assignees_url *string
    // The blobs_url property
    blobs_url *string
    // The branches_url property
    branches_url *string
    // The clone_url property
    clone_url *string
    // The collaborators_url property
    collaborators_url *string
    // The comments_url property
    comments_url *string
    // The commits_url property
    commits_url *string
    // The compare_url property
    compare_url *string
    // The contents_url property
    contents_url *string
    // The contributors_url property
    contributors_url *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The default_branch property
    default_branch *string
    // The delete_branch_on_merge property
    delete_branch_on_merge *bool
    // The deployments_url property
    deployments_url *string
    // The description property
    description *string
    // Returns whether or not this repository disabled.
    disabled *bool
    // The downloads_url property
    downloads_url *string
    // The events_url property
    events_url *string
    // The fork property
    fork *bool
    // The forks property
    forks *int32
    // The forks_count property
    forks_count *int32
    // The forks_url property
    forks_url *string
    // The full_name property
    full_name *string
    // The git_commits_url property
    git_commits_url *string
    // The git_refs_url property
    git_refs_url *string
    // The git_tags_url property
    git_tags_url *string
    // The git_url property
    git_url *string
    // The has_discussions property
    has_discussions *bool
    // The has_downloads property
    has_downloads *bool
    // The has_issues property
    has_issues *bool
    // The has_pages property
    has_pages *bool
    // The has_projects property
    has_projects *bool
    // The has_wiki property
    has_wiki *bool
    // The homepage property
    homepage *string
    // The hooks_url property
    hooks_url *string
    // The html_url property
    html_url *string
    // The id property
    id *int32
    // The is_template property
    is_template *bool
    // The issue_comment_url property
    issue_comment_url *string
    // The issue_events_url property
    issue_events_url *string
    // The issues_url property
    issues_url *string
    // The keys_url property
    keys_url *string
    // The labels_url property
    labels_url *string
    // The language property
    language *string
    // The languages_url property
    languages_url *string
    // License Simple
    license NullableLicenseSimpleable
    // The master_branch property
    master_branch *string
    // The merges_url property
    merges_url *string
    // The milestones_url property
    milestones_url *string
    // The mirror_url property
    mirror_url *string
    // The name property
    name *string
    // The node_id property
    node_id *string
    // The notifications_url property
    notifications_url *string
    // The open_issues property
    open_issues *int32
    // The open_issues_count property
    open_issues_count *int32
    // A GitHub user.
    owner NullableSimpleUserable
    // The permissions property
    permissions RepoSearchResultItem_permissionsable
    // The private property
    private *bool
    // The pulls_url property
    pulls_url *string
    // The pushed_at property
    pushed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The releases_url property
    releases_url *string
    // The score property
    score *float64
    // The size property
    size *int32
    // The ssh_url property
    ssh_url *string
    // The stargazers_count property
    stargazers_count *int32
    // The stargazers_url property
    stargazers_url *string
    // The statuses_url property
    statuses_url *string
    // The subscribers_url property
    subscribers_url *string
    // The subscription_url property
    subscription_url *string
    // The svn_url property
    svn_url *string
    // The tags_url property
    tags_url *string
    // The teams_url property
    teams_url *string
    // The temp_clone_token property
    temp_clone_token *string
    // The text_matches property
    text_matches []Repositoriesable
    // The topics property
    topics []string
    // The trees_url property
    trees_url *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The url property
    url *string
    // The repository visibility: public, private, or internal.
    visibility *string
    // The watchers property
    watchers *int32
    // The watchers_count property
    watchers_count *int32
    // The web_commit_signoff_required property
    web_commit_signoff_required *bool
}
// NewRepoSearchResultItem instantiates a new RepoSearchResultItem and sets the default values.
func NewRepoSearchResultItem()(*RepoSearchResultItem) {
    m := &RepoSearchResultItem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepoSearchResultItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepoSearchResultItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepoSearchResultItem(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepoSearchResultItem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowAutoMerge gets the allow_auto_merge property value. The allow_auto_merge property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetAllowAutoMerge()(*bool) {
    return m.allow_auto_merge
}
// GetAllowForking gets the allow_forking property value. The allow_forking property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetAllowForking()(*bool) {
    return m.allow_forking
}
// GetAllowMergeCommit gets the allow_merge_commit property value. The allow_merge_commit property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetAllowMergeCommit()(*bool) {
    return m.allow_merge_commit
}
// GetAllowRebaseMerge gets the allow_rebase_merge property value. The allow_rebase_merge property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetAllowRebaseMerge()(*bool) {
    return m.allow_rebase_merge
}
// GetAllowSquashMerge gets the allow_squash_merge property value. The allow_squash_merge property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetAllowSquashMerge()(*bool) {
    return m.allow_squash_merge
}
// GetArchived gets the archived property value. The archived property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetArchived()(*bool) {
    return m.archived
}
// GetArchiveUrl gets the archive_url property value. The archive_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetArchiveUrl()(*string) {
    return m.archive_url
}
// GetAssigneesUrl gets the assignees_url property value. The assignees_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetAssigneesUrl()(*string) {
    return m.assignees_url
}
// GetBlobsUrl gets the blobs_url property value. The blobs_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetBlobsUrl()(*string) {
    return m.blobs_url
}
// GetBranchesUrl gets the branches_url property value. The branches_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetBranchesUrl()(*string) {
    return m.branches_url
}
// GetCloneUrl gets the clone_url property value. The clone_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetCloneUrl()(*string) {
    return m.clone_url
}
// GetCollaboratorsUrl gets the collaborators_url property value. The collaborators_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetCollaboratorsUrl()(*string) {
    return m.collaborators_url
}
// GetCommentsUrl gets the comments_url property value. The comments_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetCommentsUrl()(*string) {
    return m.comments_url
}
// GetCommitsUrl gets the commits_url property value. The commits_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetCommitsUrl()(*string) {
    return m.commits_url
}
// GetCompareUrl gets the compare_url property value. The compare_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetCompareUrl()(*string) {
    return m.compare_url
}
// GetContentsUrl gets the contents_url property value. The contents_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetContentsUrl()(*string) {
    return m.contents_url
}
// GetContributorsUrl gets the contributors_url property value. The contributors_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetContributorsUrl()(*string) {
    return m.contributors_url
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *RepoSearchResultItem) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDefaultBranch gets the default_branch property value. The default_branch property
// returns a *string when successful
func (m *RepoSearchResultItem) GetDefaultBranch()(*string) {
    return m.default_branch
}
// GetDeleteBranchOnMerge gets the delete_branch_on_merge property value. The delete_branch_on_merge property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetDeleteBranchOnMerge()(*bool) {
    return m.delete_branch_on_merge
}
// GetDeploymentsUrl gets the deployments_url property value. The deployments_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetDeploymentsUrl()(*string) {
    return m.deployments_url
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *RepoSearchResultItem) GetDescription()(*string) {
    return m.description
}
// GetDisabled gets the disabled property value. Returns whether or not this repository disabled.
// returns a *bool when successful
func (m *RepoSearchResultItem) GetDisabled()(*bool) {
    return m.disabled
}
// GetDownloadsUrl gets the downloads_url property value. The downloads_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetDownloadsUrl()(*string) {
    return m.downloads_url
}
// GetEventsUrl gets the events_url property value. The events_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetEventsUrl()(*string) {
    return m.events_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepoSearchResultItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allow_auto_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAutoMerge(val)
        }
        return nil
    }
    res["allow_forking"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowForking(val)
        }
        return nil
    }
    res["allow_merge_commit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMergeCommit(val)
        }
        return nil
    }
    res["allow_rebase_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowRebaseMerge(val)
        }
        return nil
    }
    res["allow_squash_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowSquashMerge(val)
        }
        return nil
    }
    res["archive_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchiveUrl(val)
        }
        return nil
    }
    res["archived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchived(val)
        }
        return nil
    }
    res["assignees_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssigneesUrl(val)
        }
        return nil
    }
    res["blobs_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlobsUrl(val)
        }
        return nil
    }
    res["branches_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranchesUrl(val)
        }
        return nil
    }
    res["clone_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloneUrl(val)
        }
        return nil
    }
    res["collaborators_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollaboratorsUrl(val)
        }
        return nil
    }
    res["comments_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommentsUrl(val)
        }
        return nil
    }
    res["commits_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitsUrl(val)
        }
        return nil
    }
    res["compare_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompareUrl(val)
        }
        return nil
    }
    res["contents_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentsUrl(val)
        }
        return nil
    }
    res["contributors_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContributorsUrl(val)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["default_branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultBranch(val)
        }
        return nil
    }
    res["delete_branch_on_merge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleteBranchOnMerge(val)
        }
        return nil
    }
    res["deployments_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentsUrl(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["disabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisabled(val)
        }
        return nil
    }
    res["downloads_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadsUrl(val)
        }
        return nil
    }
    res["events_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventsUrl(val)
        }
        return nil
    }
    res["fork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFork(val)
        }
        return nil
    }
    res["forks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForks(val)
        }
        return nil
    }
    res["forks_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForksCount(val)
        }
        return nil
    }
    res["forks_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForksUrl(val)
        }
        return nil
    }
    res["full_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullName(val)
        }
        return nil
    }
    res["git_commits_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitCommitsUrl(val)
        }
        return nil
    }
    res["git_refs_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitRefsUrl(val)
        }
        return nil
    }
    res["git_tags_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitTagsUrl(val)
        }
        return nil
    }
    res["git_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGitUrl(val)
        }
        return nil
    }
    res["has_discussions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasDiscussions(val)
        }
        return nil
    }
    res["has_downloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasDownloads(val)
        }
        return nil
    }
    res["has_issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasIssues(val)
        }
        return nil
    }
    res["has_pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasPages(val)
        }
        return nil
    }
    res["has_projects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasProjects(val)
        }
        return nil
    }
    res["has_wiki"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasWiki(val)
        }
        return nil
    }
    res["homepage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomepage(val)
        }
        return nil
    }
    res["hooks_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHooksUrl(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["is_template"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTemplate(val)
        }
        return nil
    }
    res["issue_comment_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueCommentUrl(val)
        }
        return nil
    }
    res["issue_events_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueEventsUrl(val)
        }
        return nil
    }
    res["issues_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuesUrl(val)
        }
        return nil
    }
    res["keys_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeysUrl(val)
        }
        return nil
    }
    res["labels_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelsUrl(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["languages_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguagesUrl(val)
        }
        return nil
    }
    res["license"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableLicenseSimpleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicense(val.(NullableLicenseSimpleable))
        }
        return nil
    }
    res["master_branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMasterBranch(val)
        }
        return nil
    }
    res["merges_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMergesUrl(val)
        }
        return nil
    }
    res["milestones_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMilestonesUrl(val)
        }
        return nil
    }
    res["mirror_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMirrorUrl(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["notifications_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationsUrl(val)
        }
        return nil
    }
    res["open_issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenIssues(val)
        }
        return nil
    }
    res["open_issues_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenIssuesCount(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepoSearchResultItem_permissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(RepoSearchResultItem_permissionsable))
        }
        return nil
    }
    res["private"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivate(val)
        }
        return nil
    }
    res["pulls_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullsUrl(val)
        }
        return nil
    }
    res["pushed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPushedAt(val)
        }
        return nil
    }
    res["releases_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleasesUrl(val)
        }
        return nil
    }
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["ssh_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSshUrl(val)
        }
        return nil
    }
    res["stargazers_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStargazersCount(val)
        }
        return nil
    }
    res["stargazers_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStargazersUrl(val)
        }
        return nil
    }
    res["statuses_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusesUrl(val)
        }
        return nil
    }
    res["subscribers_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscribersUrl(val)
        }
        return nil
    }
    res["subscription_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionUrl(val)
        }
        return nil
    }
    res["svn_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSvnUrl(val)
        }
        return nil
    }
    res["tags_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTagsUrl(val)
        }
        return nil
    }
    res["teams_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsUrl(val)
        }
        return nil
    }
    res["temp_clone_token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTempCloneToken(val)
        }
        return nil
    }
    res["text_matches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoriesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Repositoriesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Repositoriesable)
                }
            }
            m.SetTextMatches(res)
        }
        return nil
    }
    res["topics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetTopics(res)
        }
        return nil
    }
    res["trees_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTreesUrl(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["visibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val)
        }
        return nil
    }
    res["watchers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatchers(val)
        }
        return nil
    }
    res["watchers_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatchersCount(val)
        }
        return nil
    }
    res["web_commit_signoff_required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebCommitSignoffRequired(val)
        }
        return nil
    }
    return res
}
// GetFork gets the fork property value. The fork property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetFork()(*bool) {
    return m.fork
}
// GetForks gets the forks property value. The forks property
// returns a *int32 when successful
func (m *RepoSearchResultItem) GetForks()(*int32) {
    return m.forks
}
// GetForksCount gets the forks_count property value. The forks_count property
// returns a *int32 when successful
func (m *RepoSearchResultItem) GetForksCount()(*int32) {
    return m.forks_count
}
// GetForksUrl gets the forks_url property value. The forks_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetForksUrl()(*string) {
    return m.forks_url
}
// GetFullName gets the full_name property value. The full_name property
// returns a *string when successful
func (m *RepoSearchResultItem) GetFullName()(*string) {
    return m.full_name
}
// GetGitCommitsUrl gets the git_commits_url property value. The git_commits_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetGitCommitsUrl()(*string) {
    return m.git_commits_url
}
// GetGitRefsUrl gets the git_refs_url property value. The git_refs_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetGitRefsUrl()(*string) {
    return m.git_refs_url
}
// GetGitTagsUrl gets the git_tags_url property value. The git_tags_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetGitTagsUrl()(*string) {
    return m.git_tags_url
}
// GetGitUrl gets the git_url property value. The git_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetGitUrl()(*string) {
    return m.git_url
}
// GetHasDiscussions gets the has_discussions property value. The has_discussions property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetHasDiscussions()(*bool) {
    return m.has_discussions
}
// GetHasDownloads gets the has_downloads property value. The has_downloads property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetHasDownloads()(*bool) {
    return m.has_downloads
}
// GetHasIssues gets the has_issues property value. The has_issues property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetHasIssues()(*bool) {
    return m.has_issues
}
// GetHasPages gets the has_pages property value. The has_pages property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetHasPages()(*bool) {
    return m.has_pages
}
// GetHasProjects gets the has_projects property value. The has_projects property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetHasProjects()(*bool) {
    return m.has_projects
}
// GetHasWiki gets the has_wiki property value. The has_wiki property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetHasWiki()(*bool) {
    return m.has_wiki
}
// GetHomepage gets the homepage property value. The homepage property
// returns a *string when successful
func (m *RepoSearchResultItem) GetHomepage()(*string) {
    return m.homepage
}
// GetHooksUrl gets the hooks_url property value. The hooks_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetHooksUrl()(*string) {
    return m.hooks_url
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *RepoSearchResultItem) GetId()(*int32) {
    return m.id
}
// GetIssueCommentUrl gets the issue_comment_url property value. The issue_comment_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetIssueCommentUrl()(*string) {
    return m.issue_comment_url
}
// GetIssueEventsUrl gets the issue_events_url property value. The issue_events_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetIssueEventsUrl()(*string) {
    return m.issue_events_url
}
// GetIssuesUrl gets the issues_url property value. The issues_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetIssuesUrl()(*string) {
    return m.issues_url
}
// GetIsTemplate gets the is_template property value. The is_template property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetIsTemplate()(*bool) {
    return m.is_template
}
// GetKeysUrl gets the keys_url property value. The keys_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetKeysUrl()(*string) {
    return m.keys_url
}
// GetLabelsUrl gets the labels_url property value. The labels_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetLabelsUrl()(*string) {
    return m.labels_url
}
// GetLanguage gets the language property value. The language property
// returns a *string when successful
func (m *RepoSearchResultItem) GetLanguage()(*string) {
    return m.language
}
// GetLanguagesUrl gets the languages_url property value. The languages_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetLanguagesUrl()(*string) {
    return m.languages_url
}
// GetLicense gets the license property value. License Simple
// returns a NullableLicenseSimpleable when successful
func (m *RepoSearchResultItem) GetLicense()(NullableLicenseSimpleable) {
    return m.license
}
// GetMasterBranch gets the master_branch property value. The master_branch property
// returns a *string when successful
func (m *RepoSearchResultItem) GetMasterBranch()(*string) {
    return m.master_branch
}
// GetMergesUrl gets the merges_url property value. The merges_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetMergesUrl()(*string) {
    return m.merges_url
}
// GetMilestonesUrl gets the milestones_url property value. The milestones_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetMilestonesUrl()(*string) {
    return m.milestones_url
}
// GetMirrorUrl gets the mirror_url property value. The mirror_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetMirrorUrl()(*string) {
    return m.mirror_url
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *RepoSearchResultItem) GetName()(*string) {
    return m.name
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *RepoSearchResultItem) GetNodeId()(*string) {
    return m.node_id
}
// GetNotificationsUrl gets the notifications_url property value. The notifications_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetNotificationsUrl()(*string) {
    return m.notifications_url
}
// GetOpenIssues gets the open_issues property value. The open_issues property
// returns a *int32 when successful
func (m *RepoSearchResultItem) GetOpenIssues()(*int32) {
    return m.open_issues
}
// GetOpenIssuesCount gets the open_issues_count property value. The open_issues_count property
// returns a *int32 when successful
func (m *RepoSearchResultItem) GetOpenIssuesCount()(*int32) {
    return m.open_issues_count
}
// GetOwner gets the owner property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *RepoSearchResultItem) GetOwner()(NullableSimpleUserable) {
    return m.owner
}
// GetPermissions gets the permissions property value. The permissions property
// returns a RepoSearchResultItem_permissionsable when successful
func (m *RepoSearchResultItem) GetPermissions()(RepoSearchResultItem_permissionsable) {
    return m.permissions
}
// GetPrivate gets the private property value. The private property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetPrivate()(*bool) {
    return m.private
}
// GetPullsUrl gets the pulls_url property value. The pulls_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetPullsUrl()(*string) {
    return m.pulls_url
}
// GetPushedAt gets the pushed_at property value. The pushed_at property
// returns a *Time when successful
func (m *RepoSearchResultItem) GetPushedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.pushed_at
}
// GetReleasesUrl gets the releases_url property value. The releases_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetReleasesUrl()(*string) {
    return m.releases_url
}
// GetScore gets the score property value. The score property
// returns a *float64 when successful
func (m *RepoSearchResultItem) GetScore()(*float64) {
    return m.score
}
// GetSize gets the size property value. The size property
// returns a *int32 when successful
func (m *RepoSearchResultItem) GetSize()(*int32) {
    return m.size
}
// GetSshUrl gets the ssh_url property value. The ssh_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetSshUrl()(*string) {
    return m.ssh_url
}
// GetStargazersCount gets the stargazers_count property value. The stargazers_count property
// returns a *int32 when successful
func (m *RepoSearchResultItem) GetStargazersCount()(*int32) {
    return m.stargazers_count
}
// GetStargazersUrl gets the stargazers_url property value. The stargazers_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetStargazersUrl()(*string) {
    return m.stargazers_url
}
// GetStatusesUrl gets the statuses_url property value. The statuses_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetStatusesUrl()(*string) {
    return m.statuses_url
}
// GetSubscribersUrl gets the subscribers_url property value. The subscribers_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetSubscribersUrl()(*string) {
    return m.subscribers_url
}
// GetSubscriptionUrl gets the subscription_url property value. The subscription_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetSubscriptionUrl()(*string) {
    return m.subscription_url
}
// GetSvnUrl gets the svn_url property value. The svn_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetSvnUrl()(*string) {
    return m.svn_url
}
// GetTagsUrl gets the tags_url property value. The tags_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetTagsUrl()(*string) {
    return m.tags_url
}
// GetTeamsUrl gets the teams_url property value. The teams_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetTeamsUrl()(*string) {
    return m.teams_url
}
// GetTempCloneToken gets the temp_clone_token property value. The temp_clone_token property
// returns a *string when successful
func (m *RepoSearchResultItem) GetTempCloneToken()(*string) {
    return m.temp_clone_token
}
// GetTextMatches gets the text_matches property value. The text_matches property
// returns a []Repositoriesable when successful
func (m *RepoSearchResultItem) GetTextMatches()([]Repositoriesable) {
    return m.text_matches
}
// GetTopics gets the topics property value. The topics property
// returns a []string when successful
func (m *RepoSearchResultItem) GetTopics()([]string) {
    return m.topics
}
// GetTreesUrl gets the trees_url property value. The trees_url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetTreesUrl()(*string) {
    return m.trees_url
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *RepoSearchResultItem) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *RepoSearchResultItem) GetUrl()(*string) {
    return m.url
}
// GetVisibility gets the visibility property value. The repository visibility: public, private, or internal.
// returns a *string when successful
func (m *RepoSearchResultItem) GetVisibility()(*string) {
    return m.visibility
}
// GetWatchers gets the watchers property value. The watchers property
// returns a *int32 when successful
func (m *RepoSearchResultItem) GetWatchers()(*int32) {
    return m.watchers
}
// GetWatchersCount gets the watchers_count property value. The watchers_count property
// returns a *int32 when successful
func (m *RepoSearchResultItem) GetWatchersCount()(*int32) {
    return m.watchers_count
}
// GetWebCommitSignoffRequired gets the web_commit_signoff_required property value. The web_commit_signoff_required property
// returns a *bool when successful
func (m *RepoSearchResultItem) GetWebCommitSignoffRequired()(*bool) {
    return m.web_commit_signoff_required
}
// Serialize serializes information the current object
func (m *RepoSearchResultItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allow_auto_merge", m.GetAllowAutoMerge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allow_forking", m.GetAllowForking())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allow_merge_commit", m.GetAllowMergeCommit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allow_rebase_merge", m.GetAllowRebaseMerge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allow_squash_merge", m.GetAllowSquashMerge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("archived", m.GetArchived())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("archive_url", m.GetArchiveUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("assignees_url", m.GetAssigneesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("blobs_url", m.GetBlobsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("branches_url", m.GetBranchesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clone_url", m.GetCloneUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("collaborators_url", m.GetCollaboratorsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("comments_url", m.GetCommentsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commits_url", m.GetCommitsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("compare_url", m.GetCompareUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contents_url", m.GetContentsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contributors_url", m.GetContributorsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("default_branch", m.GetDefaultBranch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("delete_branch_on_merge", m.GetDeleteBranchOnMerge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deployments_url", m.GetDeploymentsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("disabled", m.GetDisabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("downloads_url", m.GetDownloadsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("events_url", m.GetEventsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("fork", m.GetFork())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("forks", m.GetForks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("forks_count", m.GetForksCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("forks_url", m.GetForksUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("full_name", m.GetFullName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_commits_url", m.GetGitCommitsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_refs_url", m.GetGitRefsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_tags_url", m.GetGitTagsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("git_url", m.GetGitUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_discussions", m.GetHasDiscussions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_downloads", m.GetHasDownloads())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_issues", m.GetHasIssues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_pages", m.GetHasPages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_projects", m.GetHasProjects())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_wiki", m.GetHasWiki())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("homepage", m.GetHomepage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hooks_url", m.GetHooksUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issues_url", m.GetIssuesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issue_comment_url", m.GetIssueCommentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issue_events_url", m.GetIssueEventsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_template", m.GetIsTemplate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("keys_url", m.GetKeysUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("labels_url", m.GetLabelsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languages_url", m.GetLanguagesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("license", m.GetLicense())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("master_branch", m.GetMasterBranch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("merges_url", m.GetMergesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("milestones_url", m.GetMilestonesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mirror_url", m.GetMirrorUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notifications_url", m.GetNotificationsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("open_issues", m.GetOpenIssues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("open_issues_count", m.GetOpenIssuesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("private", m.GetPrivate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pulls_url", m.GetPullsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("pushed_at", m.GetPushedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("releases_url", m.GetReleasesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("score", m.GetScore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ssh_url", m.GetSshUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("stargazers_count", m.GetStargazersCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stargazers_url", m.GetStargazersUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("statuses_url", m.GetStatusesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscribers_url", m.GetSubscribersUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscription_url", m.GetSubscriptionUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("svn_url", m.GetSvnUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tags_url", m.GetTagsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teams_url", m.GetTeamsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("temp_clone_token", m.GetTempCloneToken())
        if err != nil {
            return err
        }
    }
    if m.GetTextMatches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTextMatches()))
        for i, v := range m.GetTextMatches() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("text_matches", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTopics() != nil {
        err := writer.WriteCollectionOfStringValues("topics", m.GetTopics())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("trees_url", m.GetTreesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("visibility", m.GetVisibility())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("watchers", m.GetWatchers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("watchers_count", m.GetWatchersCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("web_commit_signoff_required", m.GetWebCommitSignoffRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepoSearchResultItem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowAutoMerge sets the allow_auto_merge property value. The allow_auto_merge property
func (m *RepoSearchResultItem) SetAllowAutoMerge(value *bool)() {
    m.allow_auto_merge = value
}
// SetAllowForking sets the allow_forking property value. The allow_forking property
func (m *RepoSearchResultItem) SetAllowForking(value *bool)() {
    m.allow_forking = value
}
// SetAllowMergeCommit sets the allow_merge_commit property value. The allow_merge_commit property
func (m *RepoSearchResultItem) SetAllowMergeCommit(value *bool)() {
    m.allow_merge_commit = value
}
// SetAllowRebaseMerge sets the allow_rebase_merge property value. The allow_rebase_merge property
func (m *RepoSearchResultItem) SetAllowRebaseMerge(value *bool)() {
    m.allow_rebase_merge = value
}
// SetAllowSquashMerge sets the allow_squash_merge property value. The allow_squash_merge property
func (m *RepoSearchResultItem) SetAllowSquashMerge(value *bool)() {
    m.allow_squash_merge = value
}
// SetArchived sets the archived property value. The archived property
func (m *RepoSearchResultItem) SetArchived(value *bool)() {
    m.archived = value
}
// SetArchiveUrl sets the archive_url property value. The archive_url property
func (m *RepoSearchResultItem) SetArchiveUrl(value *string)() {
    m.archive_url = value
}
// SetAssigneesUrl sets the assignees_url property value. The assignees_url property
func (m *RepoSearchResultItem) SetAssigneesUrl(value *string)() {
    m.assignees_url = value
}
// SetBlobsUrl sets the blobs_url property value. The blobs_url property
func (m *RepoSearchResultItem) SetBlobsUrl(value *string)() {
    m.blobs_url = value
}
// SetBranchesUrl sets the branches_url property value. The branches_url property
func (m *RepoSearchResultItem) SetBranchesUrl(value *string)() {
    m.branches_url = value
}
// SetCloneUrl sets the clone_url property value. The clone_url property
func (m *RepoSearchResultItem) SetCloneUrl(value *string)() {
    m.clone_url = value
}
// SetCollaboratorsUrl sets the collaborators_url property value. The collaborators_url property
func (m *RepoSearchResultItem) SetCollaboratorsUrl(value *string)() {
    m.collaborators_url = value
}
// SetCommentsUrl sets the comments_url property value. The comments_url property
func (m *RepoSearchResultItem) SetCommentsUrl(value *string)() {
    m.comments_url = value
}
// SetCommitsUrl sets the commits_url property value. The commits_url property
func (m *RepoSearchResultItem) SetCommitsUrl(value *string)() {
    m.commits_url = value
}
// SetCompareUrl sets the compare_url property value. The compare_url property
func (m *RepoSearchResultItem) SetCompareUrl(value *string)() {
    m.compare_url = value
}
// SetContentsUrl sets the contents_url property value. The contents_url property
func (m *RepoSearchResultItem) SetContentsUrl(value *string)() {
    m.contents_url = value
}
// SetContributorsUrl sets the contributors_url property value. The contributors_url property
func (m *RepoSearchResultItem) SetContributorsUrl(value *string)() {
    m.contributors_url = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *RepoSearchResultItem) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDefaultBranch sets the default_branch property value. The default_branch property
func (m *RepoSearchResultItem) SetDefaultBranch(value *string)() {
    m.default_branch = value
}
// SetDeleteBranchOnMerge sets the delete_branch_on_merge property value. The delete_branch_on_merge property
func (m *RepoSearchResultItem) SetDeleteBranchOnMerge(value *bool)() {
    m.delete_branch_on_merge = value
}
// SetDeploymentsUrl sets the deployments_url property value. The deployments_url property
func (m *RepoSearchResultItem) SetDeploymentsUrl(value *string)() {
    m.deployments_url = value
}
// SetDescription sets the description property value. The description property
func (m *RepoSearchResultItem) SetDescription(value *string)() {
    m.description = value
}
// SetDisabled sets the disabled property value. Returns whether or not this repository disabled.
func (m *RepoSearchResultItem) SetDisabled(value *bool)() {
    m.disabled = value
}
// SetDownloadsUrl sets the downloads_url property value. The downloads_url property
func (m *RepoSearchResultItem) SetDownloadsUrl(value *string)() {
    m.downloads_url = value
}
// SetEventsUrl sets the events_url property value. The events_url property
func (m *RepoSearchResultItem) SetEventsUrl(value *string)() {
    m.events_url = value
}
// SetFork sets the fork property value. The fork property
func (m *RepoSearchResultItem) SetFork(value *bool)() {
    m.fork = value
}
// SetForks sets the forks property value. The forks property
func (m *RepoSearchResultItem) SetForks(value *int32)() {
    m.forks = value
}
// SetForksCount sets the forks_count property value. The forks_count property
func (m *RepoSearchResultItem) SetForksCount(value *int32)() {
    m.forks_count = value
}
// SetForksUrl sets the forks_url property value. The forks_url property
func (m *RepoSearchResultItem) SetForksUrl(value *string)() {
    m.forks_url = value
}
// SetFullName sets the full_name property value. The full_name property
func (m *RepoSearchResultItem) SetFullName(value *string)() {
    m.full_name = value
}
// SetGitCommitsUrl sets the git_commits_url property value. The git_commits_url property
func (m *RepoSearchResultItem) SetGitCommitsUrl(value *string)() {
    m.git_commits_url = value
}
// SetGitRefsUrl sets the git_refs_url property value. The git_refs_url property
func (m *RepoSearchResultItem) SetGitRefsUrl(value *string)() {
    m.git_refs_url = value
}
// SetGitTagsUrl sets the git_tags_url property value. The git_tags_url property
func (m *RepoSearchResultItem) SetGitTagsUrl(value *string)() {
    m.git_tags_url = value
}
// SetGitUrl sets the git_url property value. The git_url property
func (m *RepoSearchResultItem) SetGitUrl(value *string)() {
    m.git_url = value
}
// SetHasDiscussions sets the has_discussions property value. The has_discussions property
func (m *RepoSearchResultItem) SetHasDiscussions(value *bool)() {
    m.has_discussions = value
}
// SetHasDownloads sets the has_downloads property value. The has_downloads property
func (m *RepoSearchResultItem) SetHasDownloads(value *bool)() {
    m.has_downloads = value
}
// SetHasIssues sets the has_issues property value. The has_issues property
func (m *RepoSearchResultItem) SetHasIssues(value *bool)() {
    m.has_issues = value
}
// SetHasPages sets the has_pages property value. The has_pages property
func (m *RepoSearchResultItem) SetHasPages(value *bool)() {
    m.has_pages = value
}
// SetHasProjects sets the has_projects property value. The has_projects property
func (m *RepoSearchResultItem) SetHasProjects(value *bool)() {
    m.has_projects = value
}
// SetHasWiki sets the has_wiki property value. The has_wiki property
func (m *RepoSearchResultItem) SetHasWiki(value *bool)() {
    m.has_wiki = value
}
// SetHomepage sets the homepage property value. The homepage property
func (m *RepoSearchResultItem) SetHomepage(value *string)() {
    m.homepage = value
}
// SetHooksUrl sets the hooks_url property value. The hooks_url property
func (m *RepoSearchResultItem) SetHooksUrl(value *string)() {
    m.hooks_url = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *RepoSearchResultItem) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The id property
func (m *RepoSearchResultItem) SetId(value *int32)() {
    m.id = value
}
// SetIssueCommentUrl sets the issue_comment_url property value. The issue_comment_url property
func (m *RepoSearchResultItem) SetIssueCommentUrl(value *string)() {
    m.issue_comment_url = value
}
// SetIssueEventsUrl sets the issue_events_url property value. The issue_events_url property
func (m *RepoSearchResultItem) SetIssueEventsUrl(value *string)() {
    m.issue_events_url = value
}
// SetIssuesUrl sets the issues_url property value. The issues_url property
func (m *RepoSearchResultItem) SetIssuesUrl(value *string)() {
    m.issues_url = value
}
// SetIsTemplate sets the is_template property value. The is_template property
func (m *RepoSearchResultItem) SetIsTemplate(value *bool)() {
    m.is_template = value
}
// SetKeysUrl sets the keys_url property value. The keys_url property
func (m *RepoSearchResultItem) SetKeysUrl(value *string)() {
    m.keys_url = value
}
// SetLabelsUrl sets the labels_url property value. The labels_url property
func (m *RepoSearchResultItem) SetLabelsUrl(value *string)() {
    m.labels_url = value
}
// SetLanguage sets the language property value. The language property
func (m *RepoSearchResultItem) SetLanguage(value *string)() {
    m.language = value
}
// SetLanguagesUrl sets the languages_url property value. The languages_url property
func (m *RepoSearchResultItem) SetLanguagesUrl(value *string)() {
    m.languages_url = value
}
// SetLicense sets the license property value. License Simple
func (m *RepoSearchResultItem) SetLicense(value NullableLicenseSimpleable)() {
    m.license = value
}
// SetMasterBranch sets the master_branch property value. The master_branch property
func (m *RepoSearchResultItem) SetMasterBranch(value *string)() {
    m.master_branch = value
}
// SetMergesUrl sets the merges_url property value. The merges_url property
func (m *RepoSearchResultItem) SetMergesUrl(value *string)() {
    m.merges_url = value
}
// SetMilestonesUrl sets the milestones_url property value. The milestones_url property
func (m *RepoSearchResultItem) SetMilestonesUrl(value *string)() {
    m.milestones_url = value
}
// SetMirrorUrl sets the mirror_url property value. The mirror_url property
func (m *RepoSearchResultItem) SetMirrorUrl(value *string)() {
    m.mirror_url = value
}
// SetName sets the name property value. The name property
func (m *RepoSearchResultItem) SetName(value *string)() {
    m.name = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *RepoSearchResultItem) SetNodeId(value *string)() {
    m.node_id = value
}
// SetNotificationsUrl sets the notifications_url property value. The notifications_url property
func (m *RepoSearchResultItem) SetNotificationsUrl(value *string)() {
    m.notifications_url = value
}
// SetOpenIssues sets the open_issues property value. The open_issues property
func (m *RepoSearchResultItem) SetOpenIssues(value *int32)() {
    m.open_issues = value
}
// SetOpenIssuesCount sets the open_issues_count property value. The open_issues_count property
func (m *RepoSearchResultItem) SetOpenIssuesCount(value *int32)() {
    m.open_issues_count = value
}
// SetOwner sets the owner property value. A GitHub user.
func (m *RepoSearchResultItem) SetOwner(value NullableSimpleUserable)() {
    m.owner = value
}
// SetPermissions sets the permissions property value. The permissions property
func (m *RepoSearchResultItem) SetPermissions(value RepoSearchResultItem_permissionsable)() {
    m.permissions = value
}
// SetPrivate sets the private property value. The private property
func (m *RepoSearchResultItem) SetPrivate(value *bool)() {
    m.private = value
}
// SetPullsUrl sets the pulls_url property value. The pulls_url property
func (m *RepoSearchResultItem) SetPullsUrl(value *string)() {
    m.pulls_url = value
}
// SetPushedAt sets the pushed_at property value. The pushed_at property
func (m *RepoSearchResultItem) SetPushedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.pushed_at = value
}
// SetReleasesUrl sets the releases_url property value. The releases_url property
func (m *RepoSearchResultItem) SetReleasesUrl(value *string)() {
    m.releases_url = value
}
// SetScore sets the score property value. The score property
func (m *RepoSearchResultItem) SetScore(value *float64)() {
    m.score = value
}
// SetSize sets the size property value. The size property
func (m *RepoSearchResultItem) SetSize(value *int32)() {
    m.size = value
}
// SetSshUrl sets the ssh_url property value. The ssh_url property
func (m *RepoSearchResultItem) SetSshUrl(value *string)() {
    m.ssh_url = value
}
// SetStargazersCount sets the stargazers_count property value. The stargazers_count property
func (m *RepoSearchResultItem) SetStargazersCount(value *int32)() {
    m.stargazers_count = value
}
// SetStargazersUrl sets the stargazers_url property value. The stargazers_url property
func (m *RepoSearchResultItem) SetStargazersUrl(value *string)() {
    m.stargazers_url = value
}
// SetStatusesUrl sets the statuses_url property value. The statuses_url property
func (m *RepoSearchResultItem) SetStatusesUrl(value *string)() {
    m.statuses_url = value
}
// SetSubscribersUrl sets the subscribers_url property value. The subscribers_url property
func (m *RepoSearchResultItem) SetSubscribersUrl(value *string)() {
    m.subscribers_url = value
}
// SetSubscriptionUrl sets the subscription_url property value. The subscription_url property
func (m *RepoSearchResultItem) SetSubscriptionUrl(value *string)() {
    m.subscription_url = value
}
// SetSvnUrl sets the svn_url property value. The svn_url property
func (m *RepoSearchResultItem) SetSvnUrl(value *string)() {
    m.svn_url = value
}
// SetTagsUrl sets the tags_url property value. The tags_url property
func (m *RepoSearchResultItem) SetTagsUrl(value *string)() {
    m.tags_url = value
}
// SetTeamsUrl sets the teams_url property value. The teams_url property
func (m *RepoSearchResultItem) SetTeamsUrl(value *string)() {
    m.teams_url = value
}
// SetTempCloneToken sets the temp_clone_token property value. The temp_clone_token property
func (m *RepoSearchResultItem) SetTempCloneToken(value *string)() {
    m.temp_clone_token = value
}
// SetTextMatches sets the text_matches property value. The text_matches property
func (m *RepoSearchResultItem) SetTextMatches(value []Repositoriesable)() {
    m.text_matches = value
}
// SetTopics sets the topics property value. The topics property
func (m *RepoSearchResultItem) SetTopics(value []string)() {
    m.topics = value
}
// SetTreesUrl sets the trees_url property value. The trees_url property
func (m *RepoSearchResultItem) SetTreesUrl(value *string)() {
    m.trees_url = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *RepoSearchResultItem) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The url property
func (m *RepoSearchResultItem) SetUrl(value *string)() {
    m.url = value
}
// SetVisibility sets the visibility property value. The repository visibility: public, private, or internal.
func (m *RepoSearchResultItem) SetVisibility(value *string)() {
    m.visibility = value
}
// SetWatchers sets the watchers property value. The watchers property
func (m *RepoSearchResultItem) SetWatchers(value *int32)() {
    m.watchers = value
}
// SetWatchersCount sets the watchers_count property value. The watchers_count property
func (m *RepoSearchResultItem) SetWatchersCount(value *int32)() {
    m.watchers_count = value
}
// SetWebCommitSignoffRequired sets the web_commit_signoff_required property value. The web_commit_signoff_required property
func (m *RepoSearchResultItem) SetWebCommitSignoffRequired(value *bool)() {
    m.web_commit_signoff_required = value
}
type RepoSearchResultItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAutoMerge()(*bool)
    GetAllowForking()(*bool)
    GetAllowMergeCommit()(*bool)
    GetAllowRebaseMerge()(*bool)
    GetAllowSquashMerge()(*bool)
    GetArchived()(*bool)
    GetArchiveUrl()(*string)
    GetAssigneesUrl()(*string)
    GetBlobsUrl()(*string)
    GetBranchesUrl()(*string)
    GetCloneUrl()(*string)
    GetCollaboratorsUrl()(*string)
    GetCommentsUrl()(*string)
    GetCommitsUrl()(*string)
    GetCompareUrl()(*string)
    GetContentsUrl()(*string)
    GetContributorsUrl()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDefaultBranch()(*string)
    GetDeleteBranchOnMerge()(*bool)
    GetDeploymentsUrl()(*string)
    GetDescription()(*string)
    GetDisabled()(*bool)
    GetDownloadsUrl()(*string)
    GetEventsUrl()(*string)
    GetFork()(*bool)
    GetForks()(*int32)
    GetForksCount()(*int32)
    GetForksUrl()(*string)
    GetFullName()(*string)
    GetGitCommitsUrl()(*string)
    GetGitRefsUrl()(*string)
    GetGitTagsUrl()(*string)
    GetGitUrl()(*string)
    GetHasDiscussions()(*bool)
    GetHasDownloads()(*bool)
    GetHasIssues()(*bool)
    GetHasPages()(*bool)
    GetHasProjects()(*bool)
    GetHasWiki()(*bool)
    GetHomepage()(*string)
    GetHooksUrl()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetIssueCommentUrl()(*string)
    GetIssueEventsUrl()(*string)
    GetIssuesUrl()(*string)
    GetIsTemplate()(*bool)
    GetKeysUrl()(*string)
    GetLabelsUrl()(*string)
    GetLanguage()(*string)
    GetLanguagesUrl()(*string)
    GetLicense()(NullableLicenseSimpleable)
    GetMasterBranch()(*string)
    GetMergesUrl()(*string)
    GetMilestonesUrl()(*string)
    GetMirrorUrl()(*string)
    GetName()(*string)
    GetNodeId()(*string)
    GetNotificationsUrl()(*string)
    GetOpenIssues()(*int32)
    GetOpenIssuesCount()(*int32)
    GetOwner()(NullableSimpleUserable)
    GetPermissions()(RepoSearchResultItem_permissionsable)
    GetPrivate()(*bool)
    GetPullsUrl()(*string)
    GetPushedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReleasesUrl()(*string)
    GetScore()(*float64)
    GetSize()(*int32)
    GetSshUrl()(*string)
    GetStargazersCount()(*int32)
    GetStargazersUrl()(*string)
    GetStatusesUrl()(*string)
    GetSubscribersUrl()(*string)
    GetSubscriptionUrl()(*string)
    GetSvnUrl()(*string)
    GetTagsUrl()(*string)
    GetTeamsUrl()(*string)
    GetTempCloneToken()(*string)
    GetTextMatches()([]Repositoriesable)
    GetTopics()([]string)
    GetTreesUrl()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetVisibility()(*string)
    GetWatchers()(*int32)
    GetWatchersCount()(*int32)
    GetWebCommitSignoffRequired()(*bool)
    SetAllowAutoMerge(value *bool)()
    SetAllowForking(value *bool)()
    SetAllowMergeCommit(value *bool)()
    SetAllowRebaseMerge(value *bool)()
    SetAllowSquashMerge(value *bool)()
    SetArchived(value *bool)()
    SetArchiveUrl(value *string)()
    SetAssigneesUrl(value *string)()
    SetBlobsUrl(value *string)()
    SetBranchesUrl(value *string)()
    SetCloneUrl(value *string)()
    SetCollaboratorsUrl(value *string)()
    SetCommentsUrl(value *string)()
    SetCommitsUrl(value *string)()
    SetCompareUrl(value *string)()
    SetContentsUrl(value *string)()
    SetContributorsUrl(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDefaultBranch(value *string)()
    SetDeleteBranchOnMerge(value *bool)()
    SetDeploymentsUrl(value *string)()
    SetDescription(value *string)()
    SetDisabled(value *bool)()
    SetDownloadsUrl(value *string)()
    SetEventsUrl(value *string)()
    SetFork(value *bool)()
    SetForks(value *int32)()
    SetForksCount(value *int32)()
    SetForksUrl(value *string)()
    SetFullName(value *string)()
    SetGitCommitsUrl(value *string)()
    SetGitRefsUrl(value *string)()
    SetGitTagsUrl(value *string)()
    SetGitUrl(value *string)()
    SetHasDiscussions(value *bool)()
    SetHasDownloads(value *bool)()
    SetHasIssues(value *bool)()
    SetHasPages(value *bool)()
    SetHasProjects(value *bool)()
    SetHasWiki(value *bool)()
    SetHomepage(value *string)()
    SetHooksUrl(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetIssueCommentUrl(value *string)()
    SetIssueEventsUrl(value *string)()
    SetIssuesUrl(value *string)()
    SetIsTemplate(value *bool)()
    SetKeysUrl(value *string)()
    SetLabelsUrl(value *string)()
    SetLanguage(value *string)()
    SetLanguagesUrl(value *string)()
    SetLicense(value NullableLicenseSimpleable)()
    SetMasterBranch(value *string)()
    SetMergesUrl(value *string)()
    SetMilestonesUrl(value *string)()
    SetMirrorUrl(value *string)()
    SetName(value *string)()
    SetNodeId(value *string)()
    SetNotificationsUrl(value *string)()
    SetOpenIssues(value *int32)()
    SetOpenIssuesCount(value *int32)()
    SetOwner(value NullableSimpleUserable)()
    SetPermissions(value RepoSearchResultItem_permissionsable)()
    SetPrivate(value *bool)()
    SetPullsUrl(value *string)()
    SetPushedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReleasesUrl(value *string)()
    SetScore(value *float64)()
    SetSize(value *int32)()
    SetSshUrl(value *string)()
    SetStargazersCount(value *int32)()
    SetStargazersUrl(value *string)()
    SetStatusesUrl(value *string)()
    SetSubscribersUrl(value *string)()
    SetSubscriptionUrl(value *string)()
    SetSvnUrl(value *string)()
    SetTagsUrl(value *string)()
    SetTeamsUrl(value *string)()
    SetTempCloneToken(value *string)()
    SetTextMatches(value []Repositoriesable)()
    SetTopics(value []string)()
    SetTreesUrl(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetVisibility(value *string)()
    SetWatchers(value *int32)()
    SetWatchersCount(value *int32)()
    SetWebCommitSignoffRequired(value *bool)()
}
