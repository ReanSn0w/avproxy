package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ReanSn0w/avproxy/models"
)

const (
	baseURL = "https://api.animevost.org/v1"
)

type avapi struct {
	client *http.Client
}

func NewAVapi() *avapi {
	return &avapi{
		client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
}

func (api *avapi) Last(page, count int) (models.AVItems, error) {
	url := fmt.Sprintf("%s/last?page=%v&quantity=%v", baseURL, page, count)
	data := models.AVListResponse{}
	err := api.request(&data, http.MethodGet, url, "")
	return data.Data, err
}

func (api *avapi) Search(q string) (models.AVItems, error) {
	req := baseURL + "/search"
	body := api.body(map[string]string{"name": q})
	data := models.AVListResponse{}
	err := api.request(&data, http.MethodPost, req, body)
	return data.Data, err
}

func (api *avapi) Playlist(id int) (models.Playlist, error) {
	req := baseURL + "/playlist"
	body := api.body(map[string]string{"id": fmt.Sprintf("%v", id)})
	data := models.Playlist{}
	err := api.request(&data, http.MethodPost, req, body)
	return data, err
}

func (api *avapi) Info(id int) (models.AVItem, error) {
	req := baseURL + "/info"
	body := api.body(map[string]string{"id": fmt.Sprintf("%v", id)})
	data := models.AVListResponse{}
	err := api.request(&data, http.MethodPost, req, body)
	if err != nil {
		return models.AVItem{}, err
	}

	if len(data.Data) != 1 {
		return models.AVItem{}, errors.New("не удалось получить нужный элемент")
	}

	return data.Data[0], nil
}

func (api *avapi) request(data interface{}, method string, url string, body string) error {
	var resp *http.Response
	var err error

	switch method {
	case http.MethodGet:
		resp, err = api.client.Get(url)
	case http.MethodPost:
		resp, err = api.client.Post(url, "application/x-www-form-urlencoded", bytes.NewReader([]byte(body)))
	default:
		err = errors.New("данный запрос не поддерживается")
	}

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("запрос завершился неудачно")
	}

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(data)
}

func (api *avapi) body(params map[string]string) string {
	res := ""

	for key, value := range params {
		if len(res) > 0 {
			res = res + "&"
		}

		res = res + key
		res = res + "="
		res = res + value
	}

	return res
}
