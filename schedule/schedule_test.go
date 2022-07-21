package schedule_test

import (
	"log"
	"testing"
	"time"

	"github.com/ReanSn0w/avproxy/schedule"
)

func Test_Interval(t *testing.T) {
	interval := schedule.NewSchedule().NextInterval()
	t.Logf("interval is %v", interval)

	timestamp := time.Now().Add(interval)

	if timestamp.Weekday() != 1 {
		t.Log("weekday not monday")
		t.Fail()
	}

	if timestamp.Hour() != 10 {
		t.Log("hour not 10")
		t.Fail()
	}

	if timestamp.Unix() < time.Now().Unix() {
		t.Log("timestamp must be larger then current time")
		t.Fail()
	}

	if t.Failed() {
		t.Logf("timestamp is %v", timestamp)
	}
}

func Test_MakeSchedule(t *testing.T) {
	s := schedule.NewSchedule()
	s.Refresh()

	items, err := s.MakeSchedule("Europe/Moscow")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	for day, elements := range items {
		log.Printf("Day is: %s\n", day)

		for index, element := range elements {
			log.Printf("%v) %s", index, element.Title)
		}
	}
}
