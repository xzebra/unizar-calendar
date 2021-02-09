// +build js

// WebAssembly wrapper functions to use from a web client.
package main

import (
	"syscall/js"

	"github.com/xzebra/unizar-calendar/internal/exports"
	"github.com/xzebra/unizar-calendar/internal/semester"
	"github.com/xzebra/unizar-calendar/pkg/schedules"
)

func calendarWrapper() js.Func {
	// calendarFunc(semesterNum int, subjects, schedule, exportType, calendarData string)
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 5 {
			return "Invalid number of arguments passed"
		}

		// Parameters parsing
		semesterNum := args[0].Int()
		if semesterNum < 1 || semesterNum > 2 {
			return "Invalid semester number"
		}
		subjects := args[1].String()
		if len(subjects) == 0 {
			return "Empty subjects"
		}
		schedule := args[2].String()
		if len(schedule) < 0 {
			return "Empty schedule"
		}
		var exportType exports.ExportType
		err := exportType.Set(args[3].String())
		if err != nil {
			return err.Error()
		}
		calendarData := args[4].String()
		if len(calendarData) == 0 {
			return "Empty calendar data"
		}

		// Get semester from github stored data
		sem, err := semester.NewSemesterFromData([]byte(calendarData))
		if err != nil {
			return err.Error()
		}

		// Parse user input
		parsed, err := schedules.ParseSemesterStrings(&schedules.SemesterFiles{
			Subjects: subjects,
			Schedule: schedule,
		})
		if err != nil {
			return err.Error()
		}

		// Generate merged data from user input and calendar data from
		// github.
		data, err := semester.NewData(sem, parsed, semesterNum)
		if err != nil {
			return err.Error()
		}

		// Export with user export option
		return exports.Export(data, exportType)
	})
}

func main() {
	js.Global().Set("calendar", calendarWrapper())
	select {} // don't let process end
}
