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

// GetCalendarDays returns a list of all days with an event in the calendar.
func (c *GoogleCalendar) GetCalendarDays(id string) (out []time.Time, err error) {
	t := time.Now().Format(time.RFC3339)
	events, err := c.srv.Events.List(id).
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(t).
		OrderBy("startTime").
		Do()
	if err != nil {
		return out, fmt.Errorf("unable to get calendar: %v", err)
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
func (c *GoogleCalendar) GetCalendarEventDays(id string) (out EventDays, err error) {
	out = make(EventDays)

	t := time.Now().Format(time.RFC3339)
	events, err := c.srv.Events.List(id).
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(t).
		OrderBy("startTime").
		Do()
	if err != nil {
		return out, fmt.Errorf("unable to get calendar: %v", err)
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
