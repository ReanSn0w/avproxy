package main

import (
	"net/http"

	"github.com/ReanSn0w/avproxy/api"
	"github.com/go-chi/chi"
)

var (
	data = api.NewAPI()
)

// @title          AnimeVost API
// @version        1.0
// @description    Прокси сервер для api animevost
// @contact.name   Дмитрий Папков
// @contact.email  papkovda@me.com
// @host           reansn0w.ru
// @BasePath       /avproxy/v1
func main() {
	r := chi.NewRouter()

	r.Use(CORS)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/last", lastTitles)
		r.Get("/search", searchTitles)
		r.Get("/info", titleInfo)
	})

	http.ListenAndServe(":8080", r)
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := w.Header()
		headers.Add("Access-Control-Allow-Origin", "*")
		headers.Add("Access-Control-Allow-Methods", "GET, OPTIONS")
		headers.Add("Access-Control-Allow-Headers", "Content-Type")
		headers.Add("Access-Control-Max-Age", "86400")

		next.ServeHTTP(w, r)
	})
}

// @Summary      Список обновлений
// @Description  Вернет список новый аниме которые, были опубликованы на сайте
// @Accept       json
// @Produce      json
// @Param        page  query  int  false  "Номер страницы для вывода элементов. В случае отсутствия будет: 1"
// @Param        count query  int  false  "Кол-во элементов на странице. В случае отсутствия будет: 20"
// @Success      200  {object}  []models.Title
// @Failure      500  {object}  api.HTTPError
// @Router       /last [get]
func lastTitles(w http.ResponseWriter, r *http.Request) {
	page, count := api.PageInfo(r)

	items, err := data.Last(page, count)
	if err != nil {
		api.SendError(w, 500, err)
		return
	}

	api.SendOK(w, 200, items)
}

// @Summary      Поиск сериалов
// @Description  Вернет список аниме, которые соответствуют поисковуму запросу
// @Accept       json
// @Produce      json
// @Param        q    query  string  true  "Текст для поиска. Должен содержать не менее 3 символов"
// @Success      200  {object}  []models.Title
// @Failure      400  {object}  api.HTTPError
// @Failure      500  {object}  api.HTTPError
// @Router       /search [get]
func searchTitles(w http.ResponseWriter, r *http.Request) {
	query, err := api.SearchInfo(r)
	if err != nil {
		api.SendError(w, 400, err)
		return
	}

	items, err := data.Search(query)
	if err != nil {
		api.SendError(w, 500, err)
		return
	}

	api.SendOK(w, 200, items)
}

// @Summary      Информация о сериале
// @Description  Вернет информацию о сериале со списком серий
// @Accept       json
// @Produce      json
// @Param        id   query  int  true  "Идентификатор сериала"
// @Success      200  {object}  models.Title
// @Failure      400  {object}  api.HTTPError
// @Failure      500  {object}  api.HTTPError
// @Router       /info [get]
func titleInfo(w http.ResponseWriter, r *http.Request) {
	id, err := api.IntFromQuery(r, "id")
	if err != nil {
		api.SendError(w, 400, err)
		return
	}

	title, err := data.Info(id)
	if err != nil {
		api.SendError(w, 500, err)
		return
	}

	api.SendOK(w, 200, title)
}
