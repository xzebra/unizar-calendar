package semester

import (
	"encoding/json"
	"time"

	"github.com/xzebra/unizar-calendar/pkg/gcal"
)

var (
	daysA      = "eina.unizar.es_hlti3ac2pou7knidr6e6267g4s@group.calendar.google.com"
	daysB      = "eina.unizar.es_ri3mten96cc0s8am0hm080bi94@group.calendar.google.com"
	holidays   = "eina.unizar.es_nvgat6f556c48fmtk7llb5i5l0@group.calendar.google.com"
	exams      = "eina.unizar.es_8g43cd660rntvu09n32g4hsonk@group.calendar.google.com"
	evaluation = "eina.unizar.es_9vuatq1d533o3aoknsej9vbiv8@group.calendar.google.com"
)

var semesterEvents = map[int]struct{ Begin, End string }{
	1: {
		Begin: "Comienzo clases 1er Semestre",
		End:   "Final clases 1er Semestre",
	},
	2: {
		Begin: "Comienzo clases 2º Semestre",
		End:   "Final clases 2º Semestre",
	},
}

type Semester struct {
	Number     int
	Begin, End time.Time
	Days       gcal.EventDays
}

func (s *Semester) findStartAndEnd(cal *gcal.GoogleCalendar, number int) error {
	year := time.Now().Year()

	// Find begind and end of semester in current or previous year.
	for year >= time.Now().Year()-1 {
		eval, err := cal.GetYearCalendarEvents(evaluation, year)
		if err != nil {
			return err
		}

		// Find the start and end of semester events.
		for _, evalEvent := range eval {
			switch evalEvent.Name {
			case semesterEvents[number].Begin:
				s.Begin = evalEvent.Start
			case semesterEvents[number].End:
				s.End = evalEvent.Start
			}
		}

		// Try with previous year.
		year--
	}

	return nil
}

func NewSemester(cal *gcal.GoogleCalendar, number int) (semester *Semester, err error) {
	semester = &Semester{
		Number: number,
	}
	semester.findStartAndEnd(cal, number)

	// Get type A and B days (days with practical classes).
	semester.Days, err = getCalendarDays(cal, semester)
	if err != nil {
		return nil, err
	}

	// Add non practical classes too.
	err = addRestOfLectiveDays(cal, semester, semester.Days)
	if err != nil {
		return nil, err
	}

	return semester, nil
}

func NewSemesterFromData(data []byte) (semester *Semester, err error) {
	semester = &Semester{}
	err = json.Unmarshal(data, semester)
	return
}

type TimeRange struct {
	UUID       string
	Start, End time.Time
}

// mergeDayTypes joins both EventDays objects.
func mergeDayTypes(a, b gcal.EventDays) (c gcal.EventDays) {
	c = make(gcal.EventDays)
	for k, v := range a {
		c[k] = v
	}
	for k, v := range b {
		c[k] = v
	}
	return
}

// getCalendarDays fetchs event days from Google Calendar and returns
// a merged list of type A and B days.
func getCalendarDays(cal *gcal.GoogleCalendar, semester *Semester) (gcal.EventDays, error) {
	// Get both type of days from different calendars

	daysTypeA, err := cal.GetCalendarEventDays(
		daysA,
		semester.Begin,
		semester.End,
	)
	if err != nil {
		return nil, err
	}

	daysTypeB, err := cal.GetCalendarEventDays(
		daysB,
		semester.Begin,
		semester.End,
	)
	if err != nil {
		return nil, err
	}

	// Merge both calendars
	return mergeDayTypes(daysTypeA, daysTypeB), nil
}

func addRestOfLectiveDays(cal *gcal.GoogleCalendar, semester *Semester, days gcal.EventDays) error {
	holidaysMask, err := cal.GetCalendarEventMask(holidays, semester.Begin, semester.End)
	if err != nil {
		return err
	}

	// Iterate semester days
	for day := semester.Begin; !day.After(semester.End); day = day.Add(time.Hour * 24) {
		// Skip weekends
		if day.Weekday() == time.Saturday || day.Weekday() == time.Sunday {
			continue
		}

		// Skip days which have a day type
		if _, ok := days[day]; ok {
			continue
		}

		// Skip holidays
		if holidaysMask[day] {
			continue
		}

		// Add a wildcard event
		switch day.Weekday() {
		case time.Monday:
			days[day] = "Lx"
		case time.Tuesday:
			days[day] = "Mx"
		case time.Wednesday:
			days[day] = "Xx"
		case time.Thursday:
			days[day] = "Jx"
		case time.Friday:
			days[day] = "Vx"
		}
	}

	return nil
}
