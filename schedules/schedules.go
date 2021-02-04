package schedules

import (
	"encoding/csv"
	"io"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type hour struct {
	Hour, Minute int
}

func (h *hour) AddTo(t time.Time) time.Time {
	return t.Add(
		time.Duration(h.Hour)*time.Hour +
			time.Duration(h.Minute)*time.Minute)
}

func (h *hour) UnmarshalCSV(csv string) error {
	t, err := time.Parse("15:04", csv)
	if err != nil {
		return err
	}

	h.Hour = t.Hour()
	h.Minute = t.Minute()
	return nil
}

type ClassNames map[string]*ClassName

type ClassName struct {
	ID   string `csv:"class_id"`
	Name string `csv:"class_name"`
	Desc string `csv:"class_desc"`
}

// Schedule is a map that given a day type (La, Lb, Ma...) returns the
// schedule of that day type.
type Schedule map[string][]*ScheduleClass

type ScheduleClass struct {
	Weekday string `csv:"weekday"`
	ID      string `csv:"class_id"`
	Start   hour   `csv:"start_hour"`
	End     hour   `csv:"end_hour"`
}

func ParseClassNames(filename string) (out ClassNames, err error) {
	out = make(ClassNames)

	f, err := os.Open(filename)
	if err != nil {
		return out, err
	}

	var classes []*ClassName
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		r.Comment = '#'
		return r
	})
	if err = gocsv.UnmarshalFile(f, &classes); err != nil {
		return
	}

	for _, class := range classes {
		out[class.ID] = class
	}

	return
}

func ParseSchedule(filename string) (out Schedule, err error) {
	out = make(Schedule)

	f, err := os.Open(filename)
	if err != nil {
		return out, err
	}

	var classes []*ScheduleClass
	if err = gocsv.UnmarshalFile(f, &classes); err != nil {
		return
	}

	for _, class := range classes {
		if class.Weekday[1] == 'x' {
			weekdayA := class.Weekday[:1] + "a"
			weekdayB := class.Weekday[:1] + "b"
			out[weekdayA] = append(out[weekdayA], class)
			out[weekdayB] = append(out[weekdayB], class)
		} else {
			out[class.Weekday] = append(out[class.Weekday], class)
		}
	}

	return
}

type SemesterFiles struct {
	Subjects string
	Schedule string
}

type ParsedSemesterFiles struct {
	Names    ClassNames
	Schedule Schedule
}

func ParseSemesterFiles(files *SemesterFiles) (sem ParsedSemesterFiles, err error) {
	sem.Names, err = ParseClassNames(files.Subjects)
	if err != nil {
		return
	}

	sem.Schedule, err = ParseSchedule(files.Schedule)

	return
}
