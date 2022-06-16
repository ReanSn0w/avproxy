package models_test

import (
	"testing"

	"github.com/ReanSn0w/avproxy/models"
)

func TestTitleSplitter(t *testing.T) {
	cases := []struct {
		Text          string
		Title         string
		OriginalTitle string
		Episodes      string
	}{
		{
			Text:          "Непутёвый ученик в школе магии: Воспоминания / Mahouka Koukou no Rettousei: Tsuioku-hen [1 из 1]",
			Title:         "Непутёвый ученик в школе магии: Воспоминания",
			OriginalTitle: "Mahouka Koukou no Rettousei: Tsuioku-hen",
			Episodes:      "1 из 1",
		},
		{
			Text:          "Непутёвый ученик в школе магии / Mahouka Koukou no Rettousei [1-26 из 26]",
			Title:         "Непутёвый ученик в школе магии",
			OriginalTitle: "Mahouka Koukou no Rettousei",
			Episodes:      "1-26 из 26",
		},
		{
			Text:          "Перестану быть героем / Yuusha, Yamemasu [1-11 из 12] [12 серия - 21 июня]",
			Title:         "Перестану быть героем",
			OriginalTitle: "Yuusha, Yamemasu",
			Episodes:      "1-11 из 12",
		},
		{
			Text:          "Восхождение героя щита (второй сезон) / Tate no Yuusha no Nariagari 2nd Season [1-10 из 12+] [11 серия - 15 июня]",
			Title:         "Восхождение героя щита (второй сезон)",
			OriginalTitle: "Tate no Yuusha no Nariagari 2nd Season",
			Episodes:      "1-10 из 12+",
		},
	}

	for _, c := range cases {
		title, original, episodes := models.PrepareAVTitle(c.Text)

		if title != c.Title {
			t.Logf("Проверка заголовка закончилась неудачно. Ошибка в заголовке: %v", c.Title)
			t.Logf("%s != %s", title, c.Title)
			t.Fail()
		}

		if original != c.OriginalTitle {
			t.Logf("Проверка оригинального заголовка закончилась неудачно. Ошибка в заголовке: %v", c.OriginalTitle)
			t.Logf("%s != %s", original, c.OriginalTitle)
			t.Fail()
		}

		if episodes != c.Episodes {
			t.Logf("Проверка списка эпизодов закончилась неудачно. Ошибка в заголовке: %v", c.Episodes)
			t.Logf("%s != %s", episodes, c.Episodes)
			t.Fail()
		}
	}
}

func TestPlaylistSort(t *testing.T) {
	pl := models.Playlist{
		{Name: "1 ep"},
		{Name: "3 ep"},
		{Name: "2 ep"},
		{Name: "OVA"},
		{Name: "OVA 2"},
	}

	res := []string{"1 ep", "2 ep", "3 ep", "OVA", "OVA 2"}

	pl.Sort()
	for index := range pl {
		if pl[index].Name != res[index] {
			t.Logf("Ошибка в сортировке элементов на позиции %v, тайтл %s, вместо %s", index, pl[index].Name, res[index])
			t.Fail()
		}
	}
}
