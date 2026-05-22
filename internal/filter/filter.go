package filter

import (
	"strings"

	"github.com/user/cronmap/internal/parser"
)

// Options holds filtering criteria for cron entries.
type Options struct {
	DayOfWeek int    // -1 means no filter
	Hour      int    // -1 means no filter
	Command   string // empty means no filter
}

// DefaultOptions returns an Options with all filters disabled.
func DefaultOptions() Options {
	return Options{DayOfWeek: -1, Hour: -1}
}

// ByDay returns only entries that run on the given day (0=Sunday … 6=Saturday).
func ByDay(entries []*parser.Entry, day int) []*parser.Entry {
	var out []*parser.Entry
	for _, e := range entries {
		if e == nil {
			continue
		}
		for _, d := range e.DayOfWeek {
			if d == day {
				out = append(out, e)
				break
			}
		}
	}
	return out
}

// ByHour returns only entries whose hour list contains the given hour.
func ByHour(entries []*parser.Entry, hour int) []*parser.Entry {
	var out []*parser.Entry
	for _, e := range entries {
		if e == nil {
			continue
		}
		for _, h := range e.Hour {
			if h == hour {
				out = append(out, e)
				break
			}
		}
	}
	return out
}

// ByCommand returns entries whose command contains the given substring (case-insensitive).
func ByCommand(entries []*parser.Entry, substr string) []*parser.Entry {
	if substr == "" {
		return entries
	}
	lower := strings.ToLower(substr)
	var out []*parser.Entry
	for _, e := range entries {
		if e == nil {
			continue
		}
		if strings.Contains(strings.ToLower(e.Command), lower) {
			out = append(out, e)
		}
	}
	return out
}

// Apply applies all non-default options from opts to entries sequentially.
func Apply(entries []*parser.Entry, opts Options) []*parser.Entry {
	result := entries
	if opts.DayOfWeek >= 0 {
		result = ByDay(result, opts.DayOfWeek)
	}
	if opts.Hour >= 0 {
		result = ByHour(result, opts.Hour)
	}
	if opts.Command != "" {
		result = ByCommand(result, opts.Command)
	}
	return result
}
