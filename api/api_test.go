package api_test

import (
	"strings"
	"testing"

	"github.com/ReanSn0w/avproxy/api"
)

var (
	data = api.NewAVapi()
)

func Test_LastElements(t *testing.T) {
	_, err := data.Last(1, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func Test_SearchElements(t *testing.T) {
	res, err := data.Search("непут")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	var even bool = false
	for _, item := range res {
		if strings.Contains(strings.ToLower(item.Title), "непут") {
			even = true
		}
	}
	if !even {
		t.Log("поиск завершился неудачно")
		t.Fail()
	}
}

func Test_ElementInfo(t *testing.T) {
	_, err := data.Info(16)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func Test_ElementPlaylist(t *testing.T) {
	_, err := data.Playlist(16)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
