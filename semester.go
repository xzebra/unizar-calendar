package main

import (
	"fmt"
	"strings"
	"time"
	"unizar-calendar/gcal"
	"unizar-calendar/schedules"
)

var (
	daysA      = "eina.unizar.es_hlti3ac2pou7knidr6e6267g4s@group.calendar.google.com"
	daysB      = "eina.unizar.es_ri3mten96cc0s8am0hm080bi94@group.calendar.google.com"
	holidays   = "eina.unizar.es_nvgat6f556c48fmtk7llb5i5l0@group.calendar.google.com"
	exams      = "eina.unizar.es_8g43cd660rntvu09n32g4hsonk@group.calendar.google.com"
	evaluation = "eina.unizar.es_9vuatq1d533o3aoknsej9vbiv8@group.calendar.google.com"
)

var months = map[string]int{
	"Ene":  1,
	"Feb":  2,
	"Mar":  3,
	"Abr":  4,
	"May":  5,
	"Jun":  6,
	"Jul":  7,
	"Ago":  8,
	"Sept": 9,
	"Oct":  10,
	"Nov":  11,
	"Dic":  12,
}

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
}

func NewSemester(cal *gcal.GoogleCalendar, number int) (*Semester, error) {
	semester := &Semester{
		Number: number,
	}

	year := time.Now().Year()

	// Find begind and end of semester in current or previous year.
	for year >= time.Now().Year()-1 {
		fmt.Println("trying with year", year)
		eval, err := cal.GetYearCalendarEvents(evaluation, year)
		if err != nil {
			return nil, err
		}

		// Find the start and end of semester events.
		for _, evalEvent := range eval {
			fmt.Println(year, evalEvent.Start, evalEvent.Name)
			switch evalEvent.Name {
			case semesterEvents[number].Begin:
				semester.Begin = evalEvent.Start
			case semesterEvents[number].End:
				semester.End = evalEvent.Start
			}
		}
		fmt.Println("-----------")
		fmt.Println(semester.Begin)
		fmt.Println(semester.End)
		fmt.Println("-----------")

		// Try with previous year.
		year--
	}

	return semester, nil
}

type timeRange struct {
	Start, End time.Time
}

type SemesterData struct {
	Semester *Semester
	Schedule schedules.Schedule
	Classes  schedules.ClassNames
	days     gcal.EventDays

	// merged is an association between class ids and a list of all
	// days when the class should occur.
	merged map[string][]timeRange
}

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
func getCalendarDays(
	cal *gcal.GoogleCalendar,
	semester *Semester,
) (gcal.EventDays, error) {
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

func NewSemesterData(files *schedules.SemesterFiles, number int) (*SemesterData, error) {
	parsed, err := schedules.ParseSemesterFiles(files)
	if err != nil {
		return nil, err
	}

	cal, err := gcal.NewGoogleCalendar()
	if err != nil {
		return nil, err
	}

	semester, err := NewSemester(cal, number)
	if err != nil {
		return nil, err
	}

	fmt.Println(semester.Begin.Date())

	daysType, err := getCalendarDays(cal, semester)
	if err != nil {
		return nil, err
	}

	s := &SemesterData{
		Semester: semester,
		Schedule: parsed.Schedule,
		Classes:  parsed.Names,
		days:     daysType,
	}
	s.mergeClassesDays()

	return s, nil
}

func (s *SemesterData) mergeClassesDays() {
	s.merged = make(map[string][]timeRange)

	// For each type of day
	for dayType, days := range s.days {
		// Get classes on that day type
		sched := s.Schedule[dayType]

		// For each class on that day type
		for _, class := range sched {
			// For each day that matches that type of day
			for _, day := range days {
				// Add times associated to class
				s.merged[class.ID] = append(s.merged[class.ID], timeRange{
					Start: class.Start.AddTo(day),
					End:   class.End.AddTo(day),
				})
			}
		}
	}
}

// ToOrg exports to Emacs org-mode (https://orgmode.org/) syntaxis, to
// be read by org-agenda.
//
// Ex:
// * Inteligencia artificial
// <2020-09-29 15:00-16:00 Tue>
// <2020-09-30 15:00-16:00 Tue>
// :STYLE: habit
func (s *SemesterData) ToOrg() string {
	var out strings.Builder

	for class, times := range s.merged {
		out.WriteString("* " + s.Classes[class] + "\n")

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
	}

	return out.String()
}
