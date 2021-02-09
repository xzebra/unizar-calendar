// Usage: webdata semester1 semester2
//
// <semester1> and <semester2> are the files where each semester info
// will be stored.
package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/xzebra/unizar-calendar/internal/semester"
	"github.com/xzebra/unizar-calendar/pkg/gcal"
)

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func dumpSemester(cal *gcal.GoogleCalendar, n int) {
	semester, err := semester.NewSemester(cal, n)
	assert(err)
	data, err := json.Marshal(semester)
	assert(err)
	err = ioutil.WriteFile(os.Args[n], data, 0644)
	assert(err)
}

func main() {
	log.SetFlags(0)

	if len(os.Args) != 3 {
		log.Fatal(errors.New("not enough parameters"))
	}

	cal, err := gcal.NewGoogleCalendar()
	assert(err)

	dumpSemester(cal, 1)
	dumpSemester(cal, 2)
}
