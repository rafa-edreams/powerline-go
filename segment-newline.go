package main

import pwl "github.com/rafa-edreams/powerline-go/powerline"

func segmentNewline(p *powerline) []pwl.Segment {
	return []pwl.Segment{{NewLine: true}}
}
