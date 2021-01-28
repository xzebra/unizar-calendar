package gcal

import (
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

type GoogleCalendar struct {
	srv *calendar.Service
}

func NewGoogleCalendar() (cal *GoogleCalendar, err error) {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		return nil, fmt.Errorf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.New(client)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve Calendar client: %v", err)
	}

	return &GoogleCalendar{
		srv: srv,
	}, nil
}

func (c *GoogleCalendar) getEventsList(id string, timeMin, timeMax time.Time) (events *calendar.Events, err error) {
	listCall := c.srv.Events.List(id).
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(timeMin.Format(time.RFC3339)).
		OrderBy("startTime")

	if !timeMax.IsZero() {
		listCall = listCall.TimeMax(timeMax.Format(time.RFC3339))
	}

	events, err = listCall.Do()
	if err != nil {
		return nil, fmt.Errorf("unable to get calendar: %v", err)
	}
	return
}

type Event struct {
	Name       string
	Start, End time.Time
}

// GetCalendarEvents returns a list of all events in the calendar.
func (c *GoogleCalendar) GetCalendarEvents(id string, timeMin, timeMax time.Time) (out []*Event, err error) {
	events, err := c.getEventsList(id, timeMin, timeMax)
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
	events, err := c.getEventsList(id, timeMin, timeMax)
	if err != nil {
		return nil, err
	}

	if len(events.Items) == 0 {
		return out, nil
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

type EventDays map[string][]time.Time

// GetCalendarEventDays returns a map that given an event type (La, Lb...)
// returns the type of event that occurs that day.
func (c *GoogleCalendar) GetCalendarEventDays(id string, timeMin, timeMax time.Time) (out EventDays, err error) {
	out = make(EventDays)

	events, err := c.getEventsList(id, timeMin, timeMax)
	if err != nil {
		return nil, err
	}

	if len(events.Items) == 0 {
		return out, nil
	}

	for _, item := range events.Items {
		date, err := time.Parse("2006-01-02", item.Start.Date)
		if err != nil {
			return out, err
		}

		if len(item.Summary) < 2 {
			return out, fmt.Errorf("date contains event with wrong description")
		}
		eventType := item.Summary[:2]

		out[eventType] = append(out[eventType], date)
	}

	return
}

type EventsMask map[time.Time]bool

// GetCalendarEventMask returns a mask that given a day returns true
// if has an event.
func (c *GoogleCalendar) GetCalendarEventMask(id string, timeMin, timeMax time.Time) (out EventsMask, err error) {
	out = make(EventsMask)

	events, err := c.getEventsList(id, timeMin, timeMax)
	if err != nil {
		return nil, err
	}

	if len(events.Items) == 0 {
		return out, nil
	}

	for _, item := range events.Items {
		date, err := time.Parse("2006-01-02", item.Start.Date)
		if err != nil {
			return out, err
		}

		out[date] = true
	}

	return
}
