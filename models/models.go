package models

type Title struct {
	ID          int64    `json:"id"`                 // идентификатор
	Title       string   `json:"title"`              // название
	Original    string   `json:"original"`           // оригинальное название
	Episodes    string   `json:"episodes"`           // доступные эпизоды
	Description string   `json:"description"`        // описание
	Director    string   `json:"director"`           // режиссер
	Year        string   `json:"year"`               // год выхода
	Genres      []string `json:"genres"`             // жанры
	Counters    Counters `json:"counters"`           // счетчики
	Playlist    Playlist `json:"playlist,omitempty"` // плейлист сериала
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
