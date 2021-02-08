package gcal

import (
	"fmt"
	"strings"
	"time"

	"google.golang.org/api/calendar/v3"
)

var daysByWeekday = map[string]byte{
	"lunes":     'L',
	"martes":    'M',
	"miércoles": 'X',
	"jueves":    'J',
	"viernes":   'V',
}

type GcalService interface {
	GetEventsList(calendarID string, timeMin, timeMax time.Time) (events *calendar.Events, err error)
}

type GoogleCalendar struct {
	srv GcalService
}

func NewGoogleCalendar() (cal *GoogleCalendar, err error) {
	srv, err := newServiceImpl()
	return &GoogleCalendar{
		srv: srv,
	}, err
}

// NewGoogleCalendarFromService used to give the calendar parser a
// different Google Calendar endpoint. Actually used for mockups.
func NewGoogleCalendarFromService(srv GcalService) (cal *GoogleCalendar) {
	return &GoogleCalendar{
		srv: srv,
	}
}

type Event struct {
	Name       string
	Start, End time.Time
}

// GetCalendarEvents returns a list of all events in the calendar.
func (c *GoogleCalendar) GetCalendarEvents(id string, timeMin, timeMax time.Time) (out []*Event, err error) {
	events, err := c.srv.GetEventsList(id, timeMin, timeMax)
	if err != nil {
		return nil, err
	}

	for _, event := range events.Items {
		var start, end time.Time

		if event.Start.DateTime == "" {
			start, err = time.Parse("2006-01-02", event.Start.Date)
		} else {
			// If includes a start time, not only a date
			start, err = time.Parse(time.RFC3339, event.Start.DateTime)
		}
		if err != nil {
			return
		}

		if event.End.DateTime == "" {
			end, err = time.Parse("2006-01-02", event.End.Date)
		} else {
			// If includes an end time, not only a date
			end, err = time.Parse(time.RFC3339, event.End.DateTime)
		}

		if err != nil {
			return
		}

		out = append(out, &Event{
			Name:  event.Summary,
			Start: start,
			End:   end,
		})
	}

	return
}

func (c *GoogleCalendar) GetYearCalendarEvents(id string, year int) (out []*Event, err error) {
	// Create a timestamp at the beggining of the year.
	timeBegin := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	// Only fetch calendar events from the beggining of the year.
	return c.GetCalendarEvents(id, timeBegin, time.Time{})
}

// GetCalendarDays returns a list of all days with an event in the calendar.
func (c *GoogleCalendar) GetCalendarDays(id string, timeMin, timeMax time.Time) (out []time.Time, err error) {
	events, err := c.srv.GetEventsList(id, timeMin, timeMax)
	if err != nil {
		return nil, err
	}

	for _, item := range events.Items {
		date, err := time.Parse("2006-01-02", item.Start.Date)
		if err != nil {
			return out, err
		}
		out = append(out, date)
	}

	return
}

// EventDays is a map that given a date returns the day type (La, Lb, Lx...).
type EventDays map[time.Time]string

// GetCalendarEventDays returns a map that given a date, returns the
// event type (La, Lb...) that day belongs to.
func (c *GoogleCalendar) GetCalendarEventDays(id string, timeMin, timeMax time.Time) (out EventDays, err error) {
	out = make(EventDays)

	events, err := c.srv.GetEventsList(id, timeMin, timeMax)
	if err != nil {
		return nil, err
	}

	for _, item := range events.Items {
		date, err := time.Parse("2006-01-02", item.Start.Date)
		if err != nil {
			return out, err
		}

		if len(item.Summary) < 2 {
			return out, fmt.Errorf("date contains event with wrong description")
		}

		var eventType string
		// Day type change in non practical days
		if strings.HasPrefix(item.Summary, "Horario de ") {
			eventType = fmt.Sprintf("%c%c",
				daysByWeekday[strings.TrimPrefix(item.Summary, "Horario de ")],
				'x',
			)
		} else {
			// Extract event type (La, Mb...) from day events (La1, Mb2...).
			eventType = item.Summary[:2]
		}
		out[date] = eventType
	}

	return
}

type EventsMask map[time.Time]bool

// GetCalendarEventMask returns a mask that given a day returns true
// if has an event.
func (c *GoogleCalendar) GetCalendarEventMask(id string, timeMin, timeMax time.Time) (out EventsMask, err error) {
	out = make(EventsMask)

	events, err := c.srv.GetEventsList(id, timeMin, timeMax)
	if err != nil {
		return nil, err
	}

	for _, item := range events.Items {
		dateStart, err := time.Parse("2006-01-02", item.Start.Date)
		if err != nil {
			return out, err
		}

		dateEnd, err := time.Parse("2006-01-02", item.End.Date)
		if err != nil {
			// Add at least the start of the event
			out[dateStart] = true
			return out, err
		}

		for day := dateStart; !day.After(dateEnd); day = day.Add(24 * time.Hour) {
			out[day] = true
		}
	}

	return
}
