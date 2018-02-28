package reptile

import (
	"reptile/src/net"
	"github.com/PuerkitoBio/goquery"
	"encoding/json"
)

type Song struct {
	id           int
	title        string
	img          string
	artist       []string
	artistId     []int
	album        string
	albumId      int
	commentCount int
	lyric        string
}

func ReptileSongById(songIds ...int) ([]Song, error) {
	for index, songId := range songIds {
		log.I(TAG, index)
		reader, err := net.GetRequestForReader("http://music.163.com/song?id=" + string(songId))
		if err != nil {
			log.E(TAG, err)
			return nil, err
		}
		document, _ := goquery.NewDocumentFromReader(reader)
		id := songId
		var title string
		jsonStr := document.Find("script").First().Text()
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
			title = dat["title"].(string)
		} else {
			title = ""
			log.E(TAG, err)
		}
		img, _ := document.Find(".j-img").First().Attr("src")
	}
}
