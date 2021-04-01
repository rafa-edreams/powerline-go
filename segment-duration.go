package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	pwl "github.com/rafa-edreams/powerline-go/powerline"
)

const (
	micro  rune = '\u00B5'
	milli  rune = 'm'
	second rune = 's'
	minute rune = 'm'
	hour   rune = 'h'
)

const (
	nanoseconds  int64 = 1
	microseconds int64 = nanoseconds * 1000
	milliseconds int64 = microseconds * 1000
	seconds      int64 = milliseconds * 1000
	minutes      int64 = seconds * 60
	hours        int64 = minutes * 60
)

func segmentDuration(p *powerline) []pwl.Segment {
	if p.cfg.Duration == "" {
		return []pwl.Segment{{
			Name:       "duration",
			Content:    "No duration",
			Foreground: p.theme.DurationFg,
			Background: p.theme.DurationBg,
		}}
	}

	durationValue := strings.Trim(p.cfg.Duration, "'\"")
	durationMinValue := strings.Trim(p.cfg.DurationMin, "'\"")

	hasPrecision := strings.Contains(durationValue, ".")

	durationFloat, err := strconv.ParseFloat(durationValue, 64)
	durationMinFloat, _ := strconv.ParseFloat(durationMinValue, 64)
	if err != nil {
		return []pwl.Segment{{
			Name:       "duration",
			Content:    fmt.Sprintf("Failed to convert '%s' to a number", p.cfg.Duration),
			Foreground: p.theme.DurationFg,
			Background: p.theme.DurationBg,
		}}
	}

	if durationFloat < durationMinFloat {
		return []pwl.Segment{}
	}

	duration := time.Duration(durationFloat * float64(time.Second.Nanoseconds()))

	if duration <= 0 {
		return []pwl.Segment{}
	}

	var content string
	ns := duration.Nanoseconds()
	if ns > hours {
		hrs := ns / hours
		ns -= hrs * hours
		mins := ns / minutes
		content = fmt.Sprintf("%dh %dm", hrs, mins)
	} else if ns > minutes {
		mins := ns / minutes
		ns -= mins * minutes
		secs := ns / seconds
		content = fmt.Sprintf("%dm %ds", mins, secs)
	} else if !hasPrecision {
		secs := ns / seconds
		content = fmt.Sprintf("%ds", secs)
	} else if ns > seconds {
		secs := ns / seconds
		ns -= secs * seconds
		millis := ns / milliseconds
		content = fmt.Sprintf("%ds %dms", secs, millis)
	} else if ns > milliseconds {
		millis := ns / milliseconds
		ns -= millis * milliseconds
		micros := ns / microseconds
		content = fmt.Sprintf("%dms %d\u00B5s", millis, micros)
	} else {
		content = fmt.Sprintf("%d\u00B5s", ns/microseconds)
	}

	return []pwl.Segment{{
		Name:       "duration",
		Content:    content,
		Foreground: p.theme.DurationFg,
		Background: p.theme.DurationBg,
	}}
}
