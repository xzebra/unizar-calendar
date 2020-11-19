package schedules

import (
	"encoding/csv"
	"io"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type csvTime struct {
	time.Time
}

func (t *csvTime) UnmarshalCSV(csv string) (err error) {
	t.Time, err = time.Parse("15:04", csv)
	return err
}

type ClassNames map[string]string

type ClassName struct {
	ID   string `csv:"class_id"`
	Name string `csv:"class_name"`
}

type Schedule map[string][]*ScheduleClass

type ScheduleClass struct {
	Weekday     string  `csv:"weekday"`
	ID          string  `csv:"class_id"`
	Start       csvTime `csv:"start_hour"`
	End         csvTime `csv:"end_hour"`
	Description string  `csv:"desc"`
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
		return r
	})
	if err = gocsv.UnmarshalFile(f, &classes); err != nil {
		return
	}

	for _, class := range classes {
		out[class.ID] = class.Name
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
