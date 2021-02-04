package exports

import (
	"fmt"
	"strings"
	"unizar-calendar/semester"
)

const (
	gcalCSVSeparator  = ","
	gcalCSVDateFormat = "2006-01-02"
	gcalCSVTimeFormat = "03:04:05 PM"
)

var gcalCSVHeader = strings.Join([]string{
	"Subject",
	"Start Date",
	"Start Time",
	"End Date",
	"End Time",
	"Description",
	"Private",
}, gcalCSVSeparator)

// toGcal exports calendar to Google Calendar csv format. See the docs:
// https://support.google.com/calendar/answer/37118?co=GENIE.Platform%3DDesktop&hl=en#zippy=%2Ccreate-or-edit-a-csv-file
func toGcal(s *semester.Data) string {
	var out strings.Builder

	out.WriteString(gcalCSVHeader + "\n")

	for class, times := range s.Merged {
		name := s.Classes[class].Name
		desc := s.Classes[class].Desc

		for _, time := range times {
			// Subject
			out.WriteString(fmt.Sprintf("\"%s\"%s", name, gcalCSVSeparator))
			// Start Date
			out.WriteString(time.Start.Format(gcalCSVDateFormat) + gcalCSVSeparator)
			// Start Time
			out.WriteString(time.Start.Format(gcalCSVTimeFormat) + gcalCSVSeparator)
			// End Date
			out.WriteString(time.End.Format(gcalCSVDateFormat) + gcalCSVSeparator)
			// End Time
			out.WriteString(time.End.Format(gcalCSVTimeFormat) + gcalCSVSeparator)
			// Description
			out.WriteString(fmt.Sprintf("\"%s\"%s", desc, gcalCSVSeparator))
			// Private
			out.WriteString("True\n")
		}
	}

	return out.String()
}
