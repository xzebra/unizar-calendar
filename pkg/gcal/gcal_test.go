package gcal_test

import (
	"fmt"
	"github.com/xzebra/unizar-calendar/pkg/gcal"
	"os"
	"path"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/calendar/v3"
)

var (
	ignoreID string = "ignored by mockup"
)

// ----------------------------------------------------------
// Test mockup data
// ----------------------------------------------------------
type mockupData map[int][12][]*calendar.Event

var testdata = mockupData{
	2016: [12][]*calendar.Event{
		{},
		{},
		{},
		{ // April
			{
				Summary: "San Jorge. Día de Aragón",
				Start:   &calendar.EventDateTime{Date: "2016-04-23"},
				End:     &calendar.EventDateTime{Date: "2016-04-23"},
			},
		},
		{ // May
			{
				Summary: "La5",
				Start:   &calendar.EventDateTime{Date: "2016-05-03"},
				End:     &calendar.EventDateTime{Date: "2016-05-03"},
			},
			{
				Summary: "Ma5",
				Start:   &calendar.EventDateTime{Date: "2016-05-04"},
				End:     &calendar.EventDateTime{Date: "2016-05-04"},
			},
			{
				Summary: "Lb6",
				Start:   &calendar.EventDateTime{Date: "2016-05-11"},
				End:     &calendar.EventDateTime{Date: "2016-05-11"},
			},
			{
				Summary: "Mb6",
				Start:   &calendar.EventDateTime{Date: "2016-05-12"},
				End:     &calendar.EventDateTime{Date: "2016-05-12"},
			},
			{
				Summary: "La6",
				Start:   &calendar.EventDateTime{Date: "2016-05-17"},
				End:     &calendar.EventDateTime{Date: "2016-05-17"},
			},
		},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	},
}

// ----------------------------------------------------------
// Google API Service Mockup
// ----------------------------------------------------------

type serviceMockup struct {
	// events stores a list of events of every month given a year.
	events mockupData
}

func newServiceMockup() *serviceMockup {
	return &serviceMockup{
		events: make(mockupData),
	}
}

func (c *serviceMockup) GetEventsList(_ string, timeMin, timeMax time.Time) (events *calendar.Events, err error) {
	events = &calendar.Events{}
	startMonth := timeMin.Month()

	for year := timeMin.Year(); year <= timeMax.Year(); year++ {
		yearData, ok := c.events[year]
		if !ok {
			return nil, fmt.Errorf("mockup data does not contain information about year %d", year)
		}

		// Get limit month
		endMonth := timeMax.Month()
		// If it is the month of another year, get until end of year.
		if timeMax.Year() != year {
			endMonth = time.December
		}

		for month := startMonth - 1; month < endMonth; month++ {
			var start time.Time
			var err error
			for _, event := range yearData[month] {
				if event.Start.DateTime == "" {
					start, err = time.Parse("2006-01-02", event.Start.Date)
				} else {
					// If includes a start time, not only a date
					start, err = time.Parse(time.RFC3339, event.Start.DateTime)
				}

				if err != nil {
					return nil, fmt.Errorf("mockup data contains a wrong formatted time: %v", err)
				}

				if !start.After(timeMax) {
					events.Items = append(events.Items, event)
				}
			}
		}

	}

	return
}

// ----------------------------------------------------------
// Test functions
// ----------------------------------------------------------

var cal *gcal.GoogleCalendar
var mockup *serviceMockup

func TestGetCalendarEvents(t *testing.T) {
	timeBegin := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	timeEnd := time.Date(2016, time.December, 31, 0, 0, 0, 0, time.UTC)

	data := struct {
		ID               string
		TimeMin, TimeMax time.Time
		// Event that must be in the returned data. Only Start date
		// (Day and Month) will be used.
		Event gcal.Event
	}{
		ID:      ignoreID,
		TimeMin: timeBegin,
		TimeMax: timeEnd,
		Event: gcal.Event{
			Name:  "San Jorge. Día de Aragón",
			Start: time.Date(2016, time.April, 23, 0, 0, 0, 0, time.UTC),
		},
	}

	events, err := cal.GetCalendarEvents(data.ID, data.TimeMin, data.TimeMax)
	assert.Nil(t, err)

	found := false
	for _, event := range events {
		if event.Name != data.Event.Name {
			continue
		}

		if event.Name == data.Event.Name &&
			event.Start.Day() == data.Event.Start.Day() &&
			event.Start.Month() == data.Event.Start.Month() {

			found = true
		}
	}

	assert.Equal(t, true, found)
}

func TestGetCalendarEventDays(t *testing.T) {
	timeBegin := time.Date(2016, time.May, 1, 0, 0, 0, 0, time.UTC)
	timeEnd := time.Date(2016, time.May, 31, 0, 0, 0, 0, time.UTC)

	testData := map[string][]time.Time{
		"La": {
			time.Date(2016, time.May, 3, 0, 0, 0, 0, time.UTC),
			time.Date(2016, time.May, 17, 0, 0, 0, 0, time.UTC),
		},
		"Ma": {
			time.Date(2016, time.May, 4, 0, 0, 0, 0, time.UTC),
		},
		"Lb": {
			time.Date(2016, time.May, 11, 0, 0, 0, 0, time.UTC),
		},
		"Mb": {
			time.Date(2016, time.May, 12, 0, 0, 0, 0, time.UTC),
		},
	}

	eventDays, err := cal.GetCalendarEventDays(
		ignoreID,
		timeBegin,
		timeEnd,
	)
	assert.Nil(t, err)

	for k, got := range eventDays {
		expected, ok := testData[k]
		assert.True(t, ok)

		assert.Equal(t, expected, got)
	}
}

func TestGetCalendarDays(t *testing.T) {
	timeBegin := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	timeEnd := time.Date(2016, time.December, 31, 0, 0, 0, 0, time.UTC)

	// We asume GetCalendarEventMask is working fine.
	eventDayMask, err := cal.GetCalendarEventMask(ignoreID, timeBegin, timeEnd)
	assert.Nil(t, err)

	returned, err := cal.GetCalendarDays(ignoreID, timeBegin, timeEnd)
	assert.Nil(t, err)

	assert.Equal(t, len(eventDayMask), len(returned))

	for _, event := range returned {
		assert.True(t, eventDayMask[event])
	}
}

// ----------------------------------------------------------
// Test initialization
// ----------------------------------------------------------

func TestMain(t *testing.M) {
	// Run everything from project root folder
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("error returning to root folder: ", err)
		os.Exit(1)
	}

	mockup = newServiceMockup()

	// Create a Google Calendar API mockup
	cal = gcal.NewGoogleCalendarFromService(
		&serviceMockup{
			events: testdata,
		},
	)

	os.Exit(t.Run())
}
