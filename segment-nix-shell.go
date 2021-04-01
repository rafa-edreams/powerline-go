package main

import (
	pwl "github.com/rafa-edreams/powerline-go/powerline"
	"os"
)

func segmentNixShell(p *powerline) []pwl.Segment {
	var nixShell string
	nixShell, _ = os.LookupEnv("IN_NIX_SHELL")
	if nixShell == "" {
		return []pwl.Segment{}
	}
	return []pwl.Segment{{
		Name:       "nix-shell",
		Content:    nixShell,
		Foreground: p.theme.NixShellFg,
		Background: p.theme.NixShellBg,
	}}
}
