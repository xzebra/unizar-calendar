package gcal

import (
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

type serviceImpl struct {
	srv *calendar.Service
}

func newServiceImpl() (*serviceImpl, error) {
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

	return &serviceImpl{
		srv: srv,
	}, nil
}

func (c *serviceImpl) GetEventsList(id string, timeMin, timeMax time.Time) (events *calendar.Events, err error) {
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
