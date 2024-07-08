package models

import "time"

type Event struct {
	ID          int
	Name        string    `"binding":"required"`
	Description string    `"binding":"required"`
	Location    string    `"binding":"required"`
	DateTime    time.Time `"binding":"required"`
	UserId      int
}

var Events []Event

func Save(event Event) {
	Events = append(Events, event)
}

func GetAllEvent() []Event {
	return Events

}
