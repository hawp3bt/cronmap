package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/user/cronmap/internal/exporter"
	"github.com/user/cronmap/internal/filter"
	"github.com/user/cronmap/internal/formatter"
	"github.com/user/cronmap/internal/parser"
	"github.com/user/cronmap/internal/schedule"
)

func main() {
	filePath := flag.String("f", "", "path to crontab file (default: stdin)")
	outputFmt := flag.String("o", "text", "output format: text | json")
	filterDay := flag.Int("day", -1, "filter by day of week (0=Sun … 6=Sat, -1=all)")
	filterHour := flag.Int("hour", -1, "filter by hour (0-23, -1=all)")
	filterCmd := flag.String("cmd", "", "filter by command substring (case-insensitive)")
	flag.Parse()

	var src io.Reader = os.Stdin
	if *filePath != "" {
		f, err := os.Open(*filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cronmap: cannot open file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		src = f
	}

	raw, err := io.ReadAll(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cronmap: read error: %v\n", err)
		os.Exit(1)
	}

	entries, err := parser.Parse(strings.NewReader(string(raw)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "cronmap: parse error: %v\n", err)
		os.Exit(1)
	}

	opts := filter.Options{
		DayOfWeek: *filterDay,
		Hour:      *filterHour,
		Command:   *filterCmd,
	}
	entries = filter.Apply(entries, opts)

	week := schedule.Build(entries)

	switch *outputFmt {
	case "json":
		out, err := exporter.ToJSON(week)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cronmap: json error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(out)
	default:
		slots := formatter.HumanReadable(week)
		fmt.Print(formatter.FormatAll(slots))
	}
}
