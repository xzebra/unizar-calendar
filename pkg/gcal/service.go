package gcal

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/calendar/v3"
)

type serviceImpl struct {
	srv *calendar.Service
}

func newServiceImpl() (*serviceImpl, error) {
	srv, err := calendar.NewService(context.Background())
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
