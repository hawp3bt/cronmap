// Package exporter provides functions to export a weekly schedule
// to various output formats including YAML.
package exporter

import (
	"fmt"
	"strings"

	"github.com/user/cronmap/internal/schedule"
)

// ToYAML serialises the weekly schedule as a YAML document.
// Each day is rendered as a top-level key whose value is a list of
// slot objects with 'time' and 'command' fields.
//
// Example output:
//
//	schedule:
//	  Monday:
//	    - time: "08:00"
//	      command: /usr/bin/backup.sh
func ToYAML(week schedule.Week) string {
	var sb strings.Builder

	sb.WriteString("schedule:\n")

	for _, day := range schedule.DayNames {
		slots, ok := week[day]
		if !ok || len(slots) == 0 {
			continue
		}

		fmt.Fprintf(&sb, "  %s:\n", day)

		for _, slot := range slots {
			time := fmt.Sprintf("%02d:%02d", slot.Hour, slot.Minute)
			fmt.Fprintf(&sb, "    - time: %q\n", time)
			fmt.Fprintf(&sb, "      command: %s\n", escapeYAML(slot.Command))
		}
	}

	return sb.String()
}

// escapeYAML wraps a string in double quotes and escapes characters that
// would otherwise break YAML parsing: backslashes, double-quotes, and
// control characters.
func escapeYAML(s string) string {
	// If the value contains no special characters we can emit it bare,
	// which keeps the output readable for the common case.
	special := false
	for _, ch := range s {
		if ch == ':' || ch == '#' || ch == '"' || ch == '\'' ||
			ch == '{' || ch == '}' || ch == '[' || ch == ']' ||
			ch == ',' || ch == '&' || ch == '*' || ch == '?' ||
			ch == '|' || ch == '-' || ch == '<' || ch == '>' ||
			ch == '=' || ch == '!' || ch == '%' || ch == '@' ||
			ch == '`' || ch == '\\' || ch < 0x20 {
			special = true
			break
		}
	}
	if !special {
		return s
	}

	// Wrap in double quotes and escape internals.
	var sb strings.Builder
	sb.WriteByte('"')
	for _, ch := range s {
		switch ch {
		case '\\':
			sb.WriteString(`\\`)
		case '"':
			sb.WriteString(`\"`)
		case '\n':
			sb.WriteString(`\n`)
		case '\r':
			sb.WriteString(`\r`)
		case '\t':
			sb.WriteString(`\t`)
		default:
			if ch < 0x20 {
				fmt.Fprintf(&sb, `\u%04x`, ch)
			} else {
				sb.WriteRune(ch)
			}
		}
	}
	sb.WriteByte('"')
	return sb.String()
}
