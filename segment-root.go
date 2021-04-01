package main

import pwl "github.com/rafa-edreams/powerline-go/powerline"

func segmentRoot(p *powerline) []pwl.Segment {
	var foreground, background uint8
	if p.cfg.PrevError == 0 || p.cfg.StaticPromptIndicator {
		foreground = p.theme.CmdPassedFg
		background = p.theme.CmdPassedBg
	} else {
		foreground = p.theme.CmdFailedFg
		background = p.theme.CmdFailedBg
	}

	return []pwl.Segment{{
		Name:       "root",
		Content:    p.shell.RootIndicator,
		Foreground: foreground,
		Background: background,
	}}
}
