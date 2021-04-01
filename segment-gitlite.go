package main

import (
	"strings"

	pwl "github.com/rafa-edreams/powerline-go/powerline"
)

func segmentGitLite(p *powerline) []pwl.Segment {
	if len(p.ignoreRepos) > 0 {
		out, err := runGitCommand("git", "rev-parse", "--show-toplevel")
		if err != nil {
			return []pwl.Segment{}
		}
		out = strings.TrimSpace(out)
		if p.ignoreRepos[out] {
			return []pwl.Segment{}
		}
	}

	out, err := runGitCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return []pwl.Segment{}
	}

	status := strings.TrimSpace(out)
	var branch string

	if status == "HEAD" {
		branch = getGitDetachedBranch(p)
	} else {
		branch = status
	}

	return []pwl.Segment{{
		Name:       "git-branch",
		Content:    branch,
		Foreground: p.theme.RepoCleanFg,
		Background: p.theme.RepoCleanBg,
	}}
}
