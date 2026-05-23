package exporter

import (
	"fmt"
	"html"
	"strings"

	"github.com/example/cronmap/internal/schedule"
)

// ToHTML renders the weekly schedule as a self-contained HTML table.
func ToHTML(week schedule.Week) string {
	var sb strings.Builder

	sb.WriteString("<!DOCTYPE html>\n")
	sb.WriteString("<html lang=\"en\">\n<head>\n")
	sb.WriteString("<meta charset=\"UTF-8\">\n")
	sb.WriteString("<title>cronmap – Weekly Schedule</title>\n")
	sb.WriteString("<style>\n")
	sb.WriteString("body{font-family:sans-serif;margin:2rem;}\n")
	sb.WriteString("table{border-collapse:collapse;width:100%;}\n")
	sb.WriteString("th,td{border:1px solid #ccc;padding:0.4rem 0.8rem;text-align:left;}\n")
	sb.WriteString("th{background:#f0f0f0;}\n")
	sb.WriteString("tr:nth-child(even){background:#fafafa;}\n")
	sb.WriteString("</style>\n</head>\n<body>\n")
	sb.WriteString("<h1>cronmap – Weekly Schedule</h1>\n")
	sb.WriteString("<table>\n<thead>\n")
	sb.WriteString("<tr><th>Day</th><th>Time</th><th>Command</th></tr>\n")
	sb.WriteString("</thead>\n<tbody>\n")

	for _, day := range schedule.DayNames {
		slots, ok := week[day]
		if !ok || len(slots) == 0 {
			continue
		}
		for _, slot := range slots {
			sb.WriteString(fmt.Sprintf(
				"<tr><td>%s</td><td>%02d:%02d</td><td>%s</td></tr>\n",
				html.EscapeString(day),
				slot.Hour,
				slot.Minute,
				html.EscapeString(slot.Command),
			))
		}
	}

	sb.WriteString("</tbody>\n</table>\n</body>\n</html>\n")
	return sb.String()
}
