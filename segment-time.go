package main

import (
	pwl "github.com/rafa-edreams/powerline-go/powerline"
	"time"
)

func segmentTime(p *powerline) []pwl.Segment {
	return []pwl.Segment{{
		Name:       "time",
		Content:    time.Now().Format("15:04:05"),
		Foreground: p.theme.TimeFg,
		Background: p.theme.TimeBg,
	}}
}
