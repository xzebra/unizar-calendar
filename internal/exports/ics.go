package exports

import (
	"fmt"
	"strings"
	"time"

	"github.com/arran4/golang-ical"
	"github.com/google/uuid"
	"github.com/xzebra/unizar-calendar/internal/semester"
)

const (
	icalTimeFormat = "20060102T150405Z"
)

// toGcalICS exports calendar to Google Calendar csv format. See the docs:
// https://support.google.com/calendar/answer/37118?co=GENIE.Platform%3DDesktop&hl=en#zippy=%2Ccreate-or-edit-an-icalendar-file
func toGcalICS(s *semester.Data) string {
	cal := ics.NewCalendar()
	cal.SetProductId("-//Github: @xzebra//Unizar Calendar")

	// NOTE: Not sure about which method do I have to choose. I used publish
	// because Google Calendar exports have this `METHOD:PUBLISH`.
	cal.SetMethod(ics.MethodPublish)

	for class, times := range s.Merged {
		if len(times) == 0 {
			continue
		}

		name := s.Classes[class].Name
		desc := s.Classes[class].Desc

		event := cal.AddEvent(uuid.NewString())
		event.SetCreatedTime(time.Now())
		event.SetDtStampTime(time.Now())
		event.SetModifiedAt(time.Now())
		event.SetStartAt(times[0].Start)
		event.SetEndAt(times[0].End)
		event.SetSummary(name)
		event.SetDescription(desc)

		for _, time := range times {
			// RDATE for recurrence of events:
			// https://www.kanzaki.com/docs/ical/rdate.html
			event.AddRdate(fmt.Sprintf(
				// Set the value as period to have different durations and start
				// times.
				"RDATE;VALUE=PERIOD:%s/%s",
				time.Start.UTC().Format(icalTimeFormat), // Start time
				time.End.UTC().Format(icalTimeFormat),   // End time
			))
		}
	}

	return cal.Serialize()
}
