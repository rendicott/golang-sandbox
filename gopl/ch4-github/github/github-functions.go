// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues

package github

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
    q := url.QueryEscape(strings.Join(terms, " "))
    fullURL := IssueURL + "?q=" + q
    // fmt.Printf("fullURL: '%s'\n", fullURL)
    resp, err := http.Get(fullURL)
    if err != nil {
        return nil, err
    }

    // We must close resp.Body on all execution paths
    // (Chapter 5 presents 'defer' which makes this simpler
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("search query failed: %s", resp.Status)
    }

    var result IssuesSearchResult
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        return nil, err
    }
    resp.Body.Close()
    return &result, nil
}