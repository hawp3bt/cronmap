package exporter

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/user/cronmap/internal/schedule"
)

// XMLSchedule is the root XML element.
type XMLSchedule struct {
	XMLName xml.Name  `xml:"schedule"`
	Days    []XMLDay  `xml:"day"`
}

// XMLDay represents a single day in the XML output.
type XMLDay struct {
	Name  string    `xml:"name,attr"`
	Slots []XMLSlot `xml:"slot"`
}

// XMLSlot represents a single scheduled slot.
type XMLSlot struct {
	Hour    int    `xml:"hour,attr"`
	Minute  int    `xml:"minute,attr"`
	Command string `xml:"command"`
}

// ToXML converts a weekly schedule map into an XML string.
func ToXML(week map[string][]schedule.Slot) (string, error) {
	dayOrder := []string{
		"Sunday", "Monday", "Tuesday", "Wednesday",
		"Thursday", "Friday", "Saturday",
	}

	xs := XMLSchedule{}

	for _, day := range dayOrder {
		slots, ok := week[day]
		if !ok {
			continue
		}
		xd := XMLDay{Name: day}
		for _, s := range slots {
			xd.Slots = append(xd.Slots, XMLSlot{
				Hour:    s.Hour,
				Minute:  s.Minute,
				Command: s.Command,
			})
		}
		xs.Days = append(xs.Days, xd)
	}

	out, err := xml.MarshalIndent(xs, "", "  ")
	if err != nil {
		return "", fmt.Errorf("xml marshal: %w", err)
	}

	var sb strings.Builder
	sb.WriteString(xml.Header)
	sb.Write(out)
	return sb.String(), nil
}
