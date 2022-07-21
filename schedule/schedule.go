package schedule

import (
	"log"
	"sort"
	"time"

	"github.com/ReanSn0w/avproxy/api"
	"github.com/ReanSn0w/avproxy/models"
)

func NewSchedule() *schedule {
	s := &schedule{
		api: api.NewAPI(),
	}

	s.Refresh()

	return s
}

type schedule struct {
	api *api.API

	items []models.ScheduleItem

	timer *time.Ticker
}

func (s *schedule) MakeSchedule(loc string) (map[string][]models.ScheduleItem, error) {
	res := map[string][]models.ScheduleItem{}
	timezone, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}

	for _, item := range s.items {
		time := time.Unix(item.Timestamp, 0).In(timezone)
		weekday := WeekDayString(time.Weekday())

		if res[weekday] == nil {
			res[weekday] = []models.ScheduleItem{}
		}

		res[weekday] = append(res[weekday], item)
	}

	return res, nil
}

// Метод немедленно обновляет расписание
func (s *schedule) Refresh() error {
	items, err := s.findTitles()
	if err != nil {
		return err
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Timestamp < items[j].Timestamp
	})

	s.items = items
	return nil
}

// Поиск тайтлов для составления расписания
func (s *schedule) findTitles() ([]models.ScheduleItem, error) {
	elementsPerPage := 50
	items := []models.ScheduleItem{}

	for {
		apiElements, err := s.api.Last(1, 50)
		if err != nil {
			return items, err
		}

		for _, item := range apiElements {
			if item.Counters.NextEpisode == 0 {
				continue
			}

			items = append(items, models.ScheduleItem{
				ID:        item.ID,
				Title:     item.Title,
				Timestamp: item.Counters.NextEpisode,
			})
		}

		if len(items) < elementsPerPage {
			break
		}
	}

	return items, nil
}

// Запуск механизма обновления расписания
func (s *schedule) Run() {
	s.timer = time.NewTicker(s.NextInterval())
	go s.ticker()
}

// Плавная остановка механизма автообновления
func (s *schedule) Stop() {
	s.timer.Stop()
	s.timer = nil
}

func (s *schedule) ticker() {
	for s.timer != nil {
		<-s.timer.C

		log.Println("Запуск обновления расписания по таймеру")
		s.Refresh()

		if s.timer != nil {
			s.timer = time.NewTicker(s.NextInterval())
			break
		}
	}
}

// Возвращает интервал до 10 часов следующего понедельника
func (s *schedule) NextInterval() time.Duration {
	t := time.Now()
	t = t.Add(-(time.Second * time.Duration(t.Second())))
	t = t.Add(-(time.Minute * time.Duration(t.Minute())))
	t = t.Add(-(time.Hour * time.Duration(t.Hour())))
	t = t.Add(-(time.Hour * 24 * time.Duration(int64(t.Weekday()))))
	t = t.Add(time.Hour * 34) // 10ч понедельник

	if t.Unix() < time.Now().Unix() {
		// Если взят прошедший понедельник, то добавляется 7 дней
		t = t.Add(time.Hour * 24 * 7)
	}

	duration := t.Unix() - time.Now().Unix()
	return time.Second * time.Duration(duration)
}

// Возвращает строковый идентификатор для дня недели
func WeekDayString(day time.Weekday) string {
	switch day {
	case 0:
		return "sun"
	case 1:
		return "mon"
	case 2:
		return "tue"
	case 3:
		return "wed"
	case 4:
		return "thu"
	case 5:
		return "fri"
	case 6:
		return "sat"
	default:
		// По идее данный кейс не должен произойти никогда.
		log.Println("Для данного дня недели нет представления в виде текста")
		return "not_a_day"
	}
}
