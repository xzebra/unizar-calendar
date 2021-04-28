package exports

import (
	"fmt"
	"time"

	"github.com/arran4/golang-ical"
	"github.com/google/uuid"
	"github.com/xzebra/unizar-calendar/internal/semester"
)

const (
	icalTimeFormat                = "20060102T150405"
	componentPropertyRecurrenceId = ics.ComponentProperty(ics.PropertyRecurrenceId)
)

// toGcalICS exports calendar to Google Calendar csv format. See the docs:
// https://support.google.com/calendar/answer/37118?co=GENIE.Platform%3DDesktop&hl=en#zippy=%2Ccreate-or-edit-an-icalendar-file
func toGcalICS(s *semester.Data) string {
	cal := ics.NewCalendar()
	cal.SetProductId("-//Zebra Apps//Unizar Calendar")
	// I suppose everybody that uses this app has the Spain timezone.
	cal.SetXWRTimezone("Europe/Madrid")

	// NOTE: Not sure about which method do I have to choose. I used publish
	// because Google Calendar exports have this `METHOD:PUBLISH`.
	cal.SetMethod(ics.MethodPublish)

	for class, times := range s.Merged {
		if len(times) == 0 {
			continue
		}

		name := s.Classes[class].Name
		desc := s.Classes[class].Desc

		eventUUID := uuid.NewString()

		classTime := times[0]
		startingEvent := cal.AddEvent(eventUUID)

		startingEventStartDate := classTime.Start
		startingEventStart := startingEventStartDate.UTC().Format(icalTimeFormat)
		startingEvent.SetProperty(ics.ComponentPropertyDtStart, startingEventStart)
		startingEventEndDate := classTime.End
		startingEventEnd := startingEventEndDate.UTC().Format(icalTimeFormat)
		startingEvent.SetProperty(ics.ComponentPropertyDtEnd, startingEventEnd)

		startingEvent.SetCreatedTime(time.Now())
		startingEvent.SetDtStampTime(time.Now())
		startingEvent.SetModifiedAt(time.Now())

		startingEvent.SetSummary(name)
		startingEvent.SetDescription(desc)

		// As google exports
		startingEvent.SetSequence(0)
		startingEvent.SetTimeTransparency(ics.TransparencyOpaque)

		// RDATE for recurrence of events:
		// https://www.kanzaki.com/docs/ical/rdate.html
		//
		// Google Calendar seems to only understand comma separated RDATEs.
		rdate := startingEventStart

		for _, classTime := range times[1:] {
			start := classTime.Start.UTC().Format(icalTimeFormat)
			end := classTime.End.UTC().Format(icalTimeFormat)

			if sameHour(classTime.Start, startingEventStartDate) &&
				sameHour(classTime.End, startingEventEndDate) {
				rdate += fmt.Sprintf(",%s", start)
				continue
			}

			event := cal.AddEvent(eventUUID)
			event.SetProperty(ics.ComponentPropertyDtStart, start)
			event.SetProperty(ics.ComponentPropertyDtEnd, end)

			event.SetCreatedTime(time.Now())
			event.SetDtStampTime(time.Now())
			event.SetModifiedAt(time.Now())

			// link to parent rdate
			event.SetProperty(componentPropertyRecurrenceId, start)

			event.SetSummary(name)
			event.SetDescription(desc)

			// As google exports
			event.SetSequence(0)
			event.SetTimeTransparency(ics.TransparencyOpaque)
			event.SetStatus(ics.ObjectStatusConfirmed)
		}

		// Add total RDATE
		startingEvent.AddRdate(rdate)
	}

	return cal.Serialize()
}

func sameHour(a time.Time, b time.Time) bool {
	return a.Hour() == b.Hour() && a.Minute() == b.Minute()
}
