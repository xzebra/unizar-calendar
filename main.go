package main

import (
	"errors"
	"os"
	"strings"
	"unizar-calendar/schedules"

	"flag"
	"log"
)

var (
	ErrInvalidSemester = errors.New("semester parameter invalid")
)

func invalidFormat() {
	log.Fatalf(
		"Invalid format.\nUse: %s [options] subjectsFile scheduleFile",
		os.Args[0],
	)
}

func main() {
	log.SetFlags(0)

	var semester int
	flag.IntVar(&semester, "s", 1, "semester (1 or 2)")
	flag.Parse()

	if semester != 1 && semester != 2 {
		log.Fatal(ErrInvalidSemester)
	}

	if len(flag.Args()) != 2 {
		invalidFormat()
	}

	subjectsFile := flag.Args()[0]
	scheduleFile := flag.Args()[1]

	if strings.HasPrefix(subjectsFile, "-") ||
		strings.HasPrefix(scheduleFile, "-") {

		invalidFormat()
	}

	sem, err := NewSemesterData(
		&schedules.SemesterFiles{
			Subjects: subjectsFile,
			Schedule: scheduleFile,
		}, semester)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(sem.ToOrg())
}
