// +build js

// WebAssembly wrapper functions to use from a web client.
package main

import (
	"errors"
	"syscall/js"

	"github.com/xzebra/unizar-calendar/internal/exports"
	"github.com/xzebra/unizar-calendar/internal/semester"
	"github.com/xzebra/unizar-calendar/pkg/schedules"
)

var (
	ErrInvalidArgs     = errors.New("Invalid number of arguments passed")
	ErrEmptySubjects   = errors.New("Empty subjects")
	ErrEmptySchedule   = errors.New("Empty schedule")
	ErrEmptyCalendar   = errors.New("Empty calendar data")
	ErrInvalidSemester = errors.New("Invalid semester number")
)

func errToJS(err error) interface{} {
	errorConstructor := js.Global().Get("Error")
	return errorConstructor.New(err.Error())
}

func calendarWrapper() js.Func {
	// calendarFunc(semesterNum int, subjects, schedule, exportType, calendarData string)
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 5 {
			return errToJS(ErrInvalidArgs)
		}

		// Parameters parsing
		semesterNum := args[0].Int()
		if semesterNum < 1 || semesterNum > 2 {
			return errToJS(ErrInvalidSemester)
		}
		subjects := args[1].String()
		if len(subjects) == 0 {
			return errToJS(ErrEmptySubjects)
		}
		schedule := args[2].String()
		if len(schedule) < 0 {
			return errToJS(ErrEmptySchedule)
		}
		var exportType exports.ExportType
		err := exportType.Set(args[3].String())
		if err != nil {
			return errToJS(err)
		}
		calendarData := args[4].String()
		if len(calendarData) == 0 {
			return errToJS(ErrEmptyCalendar)
		}

		// Get semester from github stored data
		sem, err := semester.NewSemesterFromData([]byte(calendarData))
		if err != nil {
			return errToJS(err)
		}

		// Parse user input
		parsed, err := schedules.ParseSemesterStrings(&schedules.SemesterFiles{
			Subjects: subjects,
			Schedule: schedule,
		})
		if err != nil {
			return errToJS(err)
		}

		// Generate merged data from user input and calendar data from
		// github.
		data, err := semester.NewData(sem, parsed, semesterNum)
		if err != nil {
			return errToJS(err)
		}

		// Export with user export option
		return exports.Export(data, exportType)
	})
}

func main() {
	js.Global().Set("calendar", calendarWrapper())
	select {} // don't let process end
}
