package git

import "time"

type PullRequest struct {
	URL                string        `json:"url"`
	ID                 int           `json:"id"`
	NodeID             string        `json:"node_id"`
	HTMLURL            string        `json:"html_url"`
	DiffURL            string        `json:"diff_url"`
	PatchURL           string        `json:"patch_url"`
	IssueURL           string        `json:"issue_url"`
	Number             int           `json:"number"`
	State              string        `json:"state"`
	Locked             bool          `json:"locked"`
	Title              string        `json:"title"`
	User               User          `json:"user"`
	Body               interface{}   `json:"body"`
	CreatedAt          time.Time     `json:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
	ClosedAt           interface{}   `json:"closed_at"`
	MergedAt           interface{}   `json:"merged_at"`
	MergeCommitSha     string        `json:"merge_commit_sha"`
	Assignee           interface{}   `json:"assignee"`
	Assignees          []interface{} `json:"assignees"`
	RequestedReviewers []interface{} `json:"requested_reviewers"`
	RequestedTeams     []interface{} `json:"requested_teams"`
	Labels             []interface{} `json:"labels"`
	Milestone          interface{}   `json:"milestone"`
	Draft              bool          `json:"draft"`
	CommitsURL         string        `json:"commits_url"`
	ReviewCommentsURL  string        `json:"review_comments_url"`
	ReviewCommentURL   string        `json:"review_comment_url"`
	CommentsURL        string        `json:"comments_url"`
	StatusesURL        string        `json:"statuses_url"`
	Head               Head          `json:"head"`
	Base               Base          `json:"base"`
	Links              Links         `json:"_links"`
	AuthorAssociation  string        `json:"author_association"`
	AutoMerge          interface{}   `json:"auto_merge"`
	ActiveLockReason   interface{}   `json:"active_lock_reason"`
}

type Repo struct {
	ID                       int           `json:"id"`
	NodeID                   string        `json:"node_id"`
	Name                     string        `json:"name"`
	FullName                 string        `json:"full_name"`
	Private                  bool          `json:"private"`
	Owner                    User          `json:"owner"`
	HTMLURL                  string        `json:"html_url"`
	Description              interface{}   `json:"description"`
	Fork                     bool          `json:"fork"`
	URL                      string        `json:"url"`
	ForksURL                 string        `json:"forks_url"`
	KeysURL                  string        `json:"keys_url"`
	CollaboratorsURL         string        `json:"collaborators_url"`
	TeamsURL                 string        `json:"teams_url"`
	HooksURL                 string        `json:"hooks_url"`
	IssueEventsURL           string        `json:"issue_events_url"`
	EventsURL                string        `json:"events_url"`
	AssigneesURL             string        `json:"assignees_url"`
	BranchesURL              string        `json:"branches_url"`
	TagsURL                  string        `json:"tags_url"`
	BlobsURL                 string        `json:"blobs_url"`
	GitTagsURL               string        `json:"git_tags_url"`
	GitRefsURL               string        `json:"git_refs_url"`
	TreesURL                 string        `json:"trees_url"`
	StatusesURL              string        `json:"statuses_url"`
	LanguagesURL             string        `json:"languages_url"`
	StargazersURL            string        `json:"stargazers_url"`
	ContributorsURL          string        `json:"contributors_url"`
	SubscribersURL           string        `json:"subscribers_url"`
	SubscriptionURL          string        `json:"subscription_url"`
	CommitsURL               string        `json:"commits_url"`
	GitCommitsURL            string        `json:"git_commits_url"`
	CommentsURL              string        `json:"comments_url"`
	IssueCommentURL          string        `json:"issue_comment_url"`
	ContentsURL              string        `json:"contents_url"`
	CompareURL               string        `json:"compare_url"`
	MergesURL                string        `json:"merges_url"`
	ArchiveURL               string        `json:"archive_url"`
	DownloadsURL             string        `json:"downloads_url"`
	IssuesURL                string        `json:"issues_url"`
	PullsURL                 string        `json:"pulls_url"`
	MilestonesURL            string        `json:"milestones_url"`
	NotificationsURL         string        `json:"notifications_url"`
	LabelsURL                string        `json:"labels_url"`
	ReleasesURL              string        `json:"releases_url"`
	DeploymentsURL           string        `json:"deployments_url"`
	CreatedAt                time.Time     `json:"created_at"`
	UpdatedAt                time.Time     `json:"updated_at"`
	PushedAt                 time.Time     `json:"pushed_at"`
	GitURL                   string        `json:"git_url"`
	SSHURL                   string        `json:"ssh_url"`
	CloneURL                 string        `json:"clone_url"`
	SvnURL                   string        `json:"svn_url"`
	Homepage                 interface{}   `json:"homepage"`
	Size                     int           `json:"size"`
	StargazersCount          int           `json:"stargazers_count"`
	WatchersCount            int           `json:"watchers_count"`
	Language                 string        `json:"language"`
	HasIssues                bool          `json:"has_issues"`
	HasProjects              bool          `json:"has_projects"`
	HasDownloads             bool          `json:"has_downloads"`
	HasWiki                  bool          `json:"has_wiki"`
	HasPages                 bool          `json:"has_pages"`
	HasDiscussions           bool          `json:"has_discussions"`
	ForksCount               int           `json:"forks_count"`
	MirrorURL                interface{}   `json:"mirror_url"`
	Archived                 bool          `json:"archived"`
	Disabled                 bool          `json:"disabled"`
	OpenIssuesCount          int           `json:"open_issues_count"`
	License                  interface{}   `json:"license"`
	AllowForking             bool          `json:"allow_forking"`
	IsTemplate               bool          `json:"is_template"`
	WebCommitSignoffRequired bool          `json:"web_commit_signoff_required"`
	Topics                   []interface{} `json:"topics"`
	Visibility               string        `json:"visibility"`
	Forks                    int           `json:"forks"`
	OpenIssues               int           `json:"open_issues"`
	Watchers                 int           `json:"watchers"`
	DefaultBranch            string        `json:"default_branch"`
}

type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Link struct {
	Href string `json:"href"`
}

type Links struct {
	Self           Link `json:"self"`
	HTML           Link `json:"html"`
	Issue          Link `json:"issue"`
	Comments       Link `json:"comments"`
	ReviewComments Link `json:"review_comments"`
	ReviewComment  Link `json:"review_comment"`
	Commits        Link `json:"commits"`
	Statuses       Link `json:"statuses"`
}

type Head struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	Sha   string `json:"sha"`
	User  User   `json:"user"`
	Repo  Repo   `json:"repo"`
}

type Base struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	Sha   string `json:"sha"`
	User  User   `json:"user"`
	Repo  Repo   `json:"repo"`
}
