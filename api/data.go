package api

import "github.com/ReanSn0w/avproxy/models"

type API struct {
	avapi *avapi
}

func NewAPI() *API {
	return &API{
		avapi: NewAVapi(),
	}
}

func (api *API) Last(page, count int) ([]models.Title, error) {
	return api.converter(api.avapi.Last(page, count))
}

func (api *API) Search(query string) ([]models.Title, error) {
	return api.converter(api.avapi.Search(query))
}

func (api *API) Info(id int) (*models.Title, error) {
	item, err := api.avapi.Info(id)
	if err != nil {
		return nil, err
	}

	title := item.ConvertToTitle()
	playlist, _ := api.avapi.Playlist(id)
	playlist.Sort()

	title.Playlist = playlist
	return &title, nil
}

func (api *API) converter(items []models.AVItem, err error) ([]models.Title, error) {
	res := []models.Title{}
	if err != nil {
		return res, err
	}

	for _, item := range items {
		res = append(res, item.ConvertToTitle())
	}

	return res, nil
}
