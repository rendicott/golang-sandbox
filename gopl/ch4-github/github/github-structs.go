// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues

package github

import "time"

// IssueURL is the url of the github api for issue tracking
const IssueURL = "https://api.github.com/search/issues"

// IssuesSearchResult holds information from the query
// search results
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue holds info regarding a github issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User Holds info about the github user
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
