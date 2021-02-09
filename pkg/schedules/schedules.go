package schedules

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
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
	Weekday     string `csv:"weekday"`
	ID          string `csv:"class_id"`
	Start       hour   `csv:"start_hour"`
	End         hour   `csv:"end_hour"`
	IsPractical bool   `csv:"is_practical"`
}

func ParseClassNames(in io.Reader) (out ClassNames, err error) {
	out = make(ClassNames)

	var classes []*ClassName
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		r.Comment = '#'
		return r
	})
	if err = gocsv.Unmarshal(in, &classes); err != nil {
		return
	}

	for _, class := range classes {
		out[class.ID] = class
	}

	return
}

func ParseSchedule(in io.Reader) (out Schedule, err error) {
	out = make(Schedule)

	var classes []*ScheduleClass
	if err = gocsv.Unmarshal(in, &classes); err != nil {
		return
	}

	for _, class := range classes {
		out[class.Weekday] = append(out[class.Weekday], class)
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

// ParseSemesterFiles parses the files referenced by <data>.Subjects
// and <data>.Schedule filenames.
func ParseSemesterFiles(files *SemesterFiles) (sem *ParsedSemesterFiles, err error) {
	sem = &ParsedSemesterFiles{}

	// Get io.Readers from files
	subjects, err := os.Open(files.Subjects)
	if err != nil {
		return
	}
	defer subjects.Close()

	schedules, err := os.Open(files.Schedule)
	if err != nil {
		return
	}
	defer schedules.Close()

	// Parse given readers
	sem.Names, err = ParseClassNames(subjects)
	if err != nil {
		return
	}

	sem.Schedule, err = ParseSchedule(schedules)

	return
}

// ParseSemesterStrings parses <data>.Subjects and <data>.Schedule as
// strings instead of as files.
func ParseSemesterStrings(data *SemesterFiles) (sem *ParsedSemesterFiles, err error) {
	sem = &ParsedSemesterFiles{}

	sem.Names, err = ParseClassNames(strings.NewReader(data.Subjects))
	if err != nil {
		return
	}

	sem.Schedule, err = ParseSchedule(strings.NewReader(data.Schedule))
	return
}
