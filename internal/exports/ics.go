package exports

import (
	"time"

	"github.com/arran4/golang-ical"
	"github.com/xzebra/unizar-calendar/internal/semester"
)

const (
	icalTimeFormat                = "20060102T150405"
	componentPropertyRecurrenceId = ics.ComponentProperty(ics.PropertyRecurrenceId)
)

// toGcalICS exports calendar to Google Calendar csv format. See the docs:
// https://support.google.com/calendar/answer/37118?co=GENIE.Platform%3DDesktop&hl=en#zippy=%2Ccreate-or-edit-an-icalendar-file
func toGcalICS(s *semester.Data) string {
	events := newEventStorage()

	for class, times := range s.Merged {
		if len(times) == 0 {
			continue
		}

		name := s.Classes[class].Name
		desc := s.Classes[class].Desc

		for _, classTime := range times {

			if !events.exists(classTime.UUID) {
				events.createEvent(name, desc, classTime)
			} else {
				events.addRecurrence(classTime)
			}
		}

	}

	return events.serialize()
}

func sameHour(a time.Time, b time.Time) bool {
	return a.Hour() == b.Hour() && a.Minute() == b.Minute()
}

type eventStorage struct {
	cal    *ics.Calendar
	events map[string]*icsEvent
}

type icsEvent struct {
	*ics.VEvent

	// RDATE for recurrence of events:
	// https://www.kanzaki.com/docs/ical/rdate.html
	//
	// Google Calendar seems to only understand comma separated RDATEs.
	rdate string
}

func newEventStorage() *eventStorage {
	cal := ics.NewCalendar()
	cal.SetProductId("-//Zebra Apps//Unizar Calendar")
	// I suppose everybody that uses this app has the Spain timezone.
	cal.SetXWRTimezone("Europe/Madrid")

	// NOTE: Not sure about which method do I have to choose. I used publish
	// because Google Calendar exports have this `METHOD:PUBLISH`.
	cal.SetMethod(ics.MethodPublish)

	return &eventStorage{
		cal: cal,
		// Map events by UUID
		events: make(map[string]*icsEvent),
	}
}

func (e *eventStorage) exists(uuid string) bool {
	_, exists := e.events[uuid]
	return exists
}

func (e *eventStorage) createEvent(name, desc string, classTime *semester.TimeRange) {
	uuid := classTime.UUID
	start := classTime.Start.UTC().Format(icalTimeFormat)
	end := classTime.End.UTC().Format(icalTimeFormat)

	event := e.cal.AddEvent(uuid)

	event.SetProperty(ics.ComponentPropertyDtStart, start)
	event.SetProperty(ics.ComponentPropertyDtEnd, end)

	event.SetCreatedTime(time.Now())
	event.SetDtStampTime(time.Now())
	event.SetModifiedAt(time.Now())

	event.SetSummary(name)
	event.SetDescription(desc)

	// As google exports
	event.SetSequence(0)
	event.SetTimeTransparency(ics.TransparencyOpaque)

	e.events[uuid] = &icsEvent{
		rdate:  start,
		VEvent: event,
	}
}

func (e *eventStorage) addRecurrence(classTime *semester.TimeRange) {
	uuid := classTime.UUID
	start := classTime.Start.UTC().Format(icalTimeFormat)

	e.events[uuid].rdate += "," + start
}

func (e *eventStorage) serialize() string {
	for _, event := range e.events {
		// Add total RDATE
		event.AddRdate(event.rdate)
	}

	return e.cal.Serialize()
}
