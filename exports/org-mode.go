package exports

import (
	"fmt"
	"strings"
	"unizar-calendar/semester"
)

// toOrgMode exports to Emacs org-mode (https://orgmode.org/)
// syntaxis, to be read by org-agenda.
//
// Ex:
// * Inteligencia artificial
// <2020-09-29 15:00-16:00 Tue>
// <2020-09-30 15:00-16:00 Tue>
// :STYLE: habit
func toOrgMode(s *semester.Data) string {
	var out strings.Builder

	for class, times := range s.Merged {
		out.WriteString("* " + s.Classes[class].Name + "\n")

		for _, time := range times {
			out.WriteString(
				fmt.Sprintf("<%s %s-%s>\n",
					time.Start.Format("2006-01-02"),
					time.Start.Format("15:04"),
					time.End.Format("15:04"),
				),
			)
		}

		out.WriteString(":STYLE: habit\n\n")
		out.WriteString(s.Classes[class].Desc)
		out.WriteString("\n")
	}

	return out.String()
}
