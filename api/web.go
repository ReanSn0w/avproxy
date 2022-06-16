package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type HTTPError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func SendOK(w http.ResponseWriter, status int, obj interface{}) {
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(obj)
	if err != nil {
		log.Println(err)
	}
}

func SendError(w http.ResponseWriter, status int, err error) {
	log.Println(err)

	SendOK(w, status, HTTPError{
		Code:        status,
		Description: err.Error(),
	})
}

func IntFromQuery(r *http.Request, name string) (int, error) {
	return strconv.Atoi(r.URL.Query().Get(name))
}

func PageInfo(r *http.Request) (int, int) {
	page, pageErr := IntFromQuery(r, "page")
	count, countErr := IntFromQuery(r, "count")

	res := func(res int, def int, err error) int {
		if err != nil {
			return def
		} else {
			return res
		}
	}

	return res(page, 1, pageErr), res(count, 20, countErr)
}

func SearchInfo(r *http.Request) (string, error) {
	req := r.URL.Query().Get("q")
	if len(req) < 3 {
		return req, errors.New("текст запроса слишком короткий")
	}

	return req, nil
}
