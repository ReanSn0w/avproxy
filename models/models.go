package models

type Title struct {
	ID          int64    `json:"id"`                 // Идентификатор
	Title       string   `json:"title"`              // Название
	Original    string   `json:"original"`           // Оригинальное название
	Episodes    string   `json:"episodes"`           // Доступные эпизоды
	Description string   `json:"description"`        // Описание
	Director    string   `json:"director"`           // Режиссер
	Year        string   `json:"year"`               // Год выхода
	Genres      []string `json:"genres"`             // Жанры
	Counters    Counters `json:"counters"`           // Счетчики
	Playlist    Playlist `json:"playlist,omitempty"` // Плейлист сериала
	Cover       string   `json:"cover"`              // Ссылка на обложку
	Screenshots []string `json:"screenshots"`        // Ссылка на скриншоты
}

type Counters struct {
	Votes       int64 `json:"votes"`        // Кол-во голосов
	Favorites   int64 `json:"favorites"`    // Кол-во добавлений в избранное
	Likes       int64 `json:"likes"`        // Кол-во лайков
	Rating      int64 `json:"rating"`       // Рейтинг
	NextEpisode int64 `json:"next_episode"` // Временная метка по unix содержащая дату выхода нового эпизода
}

// Элемент для составления списка расписания
type ScheduleItem struct {
	ID        int64  `json:"id"`        // Идентификатор
	Title     string `json:"title"`     // Название
	Timestamp int64  `json:"timestamp"` // Временная метка по unix содержащая дату выхода нового эпизода
}
