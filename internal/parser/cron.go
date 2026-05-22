package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// CronEntry represents a parsed crontab entry.
type CronEntry struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Command    string
}

// Parse parses a single crontab line into a CronEntry.
// Returns an error if the line is malformed.
func Parse(line string) (*CronEntry, error) {
	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, "#") {
		return nil, nil
	}

	fields := strings.Fields(line)
	if len(fields) < 6 {
		return nil, fmt.Errorf("invalid cron entry: expected at least 6 fields, got %d", len(fields))
	}

	return &CronEntry{
		Minute:     fields[0],
		Hour:       fields[1],
		DayOfMonth: fields[2],
		Month:      fields[3],
		DayOfWeek:  fields[4],
		Command:    strings.Join(fields[5:], " "),
	}, nil
}

// ExpandField expands a cron field (e.g. "*/5", "1-3", "2,4") into a slice of ints
// within [min, max].
func ExpandField(field string, min, max int) ([]int, error) {
	var result []int

	for _, part := range strings.Split(field, ",") {
		if part == "*" {
			for i := min; i <= max; i++ {
				result = append(result, i)
			}
			continue
		}

		if strings.Contains(part, "/") {
			sub := strings.SplitN(part, "/", 2)
			step, err := strconv.Atoi(sub[1])
			if err != nil || step <= 0 {
				return nil, fmt.Errorf("invalid step in field %q", part)
			}
			start := min
			if sub[0] != "*" {
				start, err = strconv.Atoi(sub[0])
				if err != nil {
					return nil, fmt.Errorf("invalid start in field %q", part)
				}
			}
			for i := start; i <= max; i += step {
				result = append(result, i)
			}
			continue
		}

		if strings.Contains(part, "-") {
			bounds := strings.SplitN(part, "-", 2)
			lo, err1 := strconv.Atoi(bounds[0])
			hi, err2 := strconv.Atoi(bounds[1])
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("invalid range in field %q", part)
			}
			for i := lo; i <= hi; i++ {
				result = append(result, i)
			}
			continue
		}

		v, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid value in field %q", part)
		}
		result = append(result, v)
	}

	return result, nil
}
