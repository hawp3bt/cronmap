package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/example/cronmap/internal/exporter"
	"github.com/example/cronmap/internal/parser"
	"github.com/example/cronmap/internal/schedule"
)

func main() {
	jsonOut := flag.Bool("json", false, "output schedule as JSON")
	file := flag.String("file", "", "path to crontab file (default: stdin)")
	flag.Parse()

	var src *os.File
	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cronmap: cannot open file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		src = f
	} else {
		src = os.Stdin
	}

	var entries []*parser.Entry
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		e, err := parser.Parse(line)
		if err != nil {
			// skip unparseable lines silently
			continue
		}
		if e != nil {
			entries = append(entries, e)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "cronmap: read error: %v\n", err)
		os.Exit(1)
	}

	week := schedule.Build(entries)

	if *jsonOut {
		if err := exporter.ToJSON(os.Stdout, week); err != nil {
			fmt.Fprintf(os.Stderr, "cronmap: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Print(schedule.Render(week))
	}
}
