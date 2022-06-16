// В файле представлен список изначальных моделей для API
// Данные модели предоставляет API сайта AnimeVost в ответ на запросы к нему,
// для удобства дальнейшего использования модели пересобираются в более удобный для работы формат.
//
// К примеру следующим изменениям подвергается название сериалов:
// происходит разбиение названия на составные части, когда в отдельном поле выделено оригинальное название сериала,
// локализированное название и и кол-во доспупных эпизодов
package models

import (
	"sort"
	"strconv"
	"strings"
)

// Структура для получения ответа от сервера со списком сериалов
type AVListResponse struct {
	State State   `json:"state"`
	Data  AVItems `json:"data"`
}

type AVItems []AVItem

// Элемент списка описывающие сериал
type AVItem struct {
	ScreenImage     []string    `json:"screenImage"`
	Rating          int64       `json:"rating"`
	Description     string      `json:"description"`
	Series          string      `json:"series"`
	Director        string      `json:"director"`
	URLImagePreview string      `json:"urlImagePreview"`
	Year            string      `json:"year"`
	Genre           string      `json:"genre"`
	ID              int64       `json:"id"`
	Votes           int64       `json:"votes"`
	IsFavorite      int64       `json:"isFavorite"`
	Title           string      `json:"title"`
	Timer           interface{} `json:"timer"`
	Type            string      `json:"type"`
	IsLikes         int64       `json:"isLikes"`
}

func (item AVItem) ConvertToTitle() Title {
	name, original, episodes := PrepareAVTitle(item.Title)

	return Title{
		ID:          item.ID,
		Title:       name,
		Original:    original,
		Episodes:    episodes,
		Description: item.ClearDescription(),
		Director:    item.Director,
		Year:        item.Year,
		Genres:      item.ClearGenres(),
		Counters: Counters{
			Votes:       item.Votes,
			Favorites:   item.IsFavorite,
			Likes:       item.IsLikes,
			Rating:      item.Rating,
			NextEpisode: item.ClearTimer(),
		},
		Cover:       item.URLImagePreview,
		Screenshots: item.ClearScreenshots(),
	}
}

func (item *AVItem) ClearDescription() string {
	return item.Description
}

func (item *AVItem) ClearGenres() []string {
	return []string{item.Genre}
}

func (item *AVItem) ClearTimer() int64 {
	val, b := item.Timer.(string)
	if !b {
		return 0
	}

	v, _ := strconv.Atoi(val)
	return int64(v)
}

func (item *AVItem) ClearScreenshots() []string {
	res := []string{}

	for _, item := range item.ScreenImage {
		res = append(res, "https://static.openni.ru"+item)
	}

	return res
}

type State struct {
	Status string `json:"status"`
	Rek    int64  `json:"rek"`
	Page   int64  `json:"page"`
	Count  int64  `json:"count"`
}

type Timer struct {
	Integer *int64
	String  *string
}

type Playlist []PlaylistElement

type PlaylistElement struct {
	Std     string `json:"std"`
	Preview string `json:"preview"`
	Name    string `json:"name"`
	HD      string `json:"hd"`

	number *int
}

func (el *PlaylistElement) Index() int {
	if el.number != nil {
		return *el.number
	}

	res := strings.Split(el.Name, " ")
	num, _ := strconv.Atoi(res[0])
	el.number = &num
	return el.Index()
}

func (pl Playlist) Sort() {
	sort.Slice(pl, func(i, j int) bool {
		first := pl[i].Index()
		second := pl[j].Index()

		if first == 0 {
			return false
		}

		return first < second
	})
}
