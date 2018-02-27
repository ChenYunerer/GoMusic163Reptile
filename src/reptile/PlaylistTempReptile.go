package reptile

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"reptile/src/net"
	"reptile/src/util"
	"os"
	"io"
)

type PlaylistTemp struct {
	title     string
	img       string
	url       string
	playCount string
}

var (
	firstUrl = "http://music.163.com/discover/playlist" //爬虫开始访问的第一个URL
)

//爬取歌单临时数据
func ReptilePlaylistTemp() {
	reptilePlaylistTemp(firstUrl)
	write2File()
}

//递归爬取数据
func reptilePlaylistTemp(url string) {
	if url == "" {
		return
	}
	reader, err := net.GetRequestForReader(url)
	if err != nil {
		log.I(TAG, err.Error())
	}
	document, _ := goquery.NewDocumentFromReader(reader)
	document.Find(".u-cover, u-cover-1").Each(func(i int, selection *goquery.Selection) {
		img, _ := selection.Find("img").First().Attr("src")
		title, _ := selection.Find("a").First().Attr("title")
		href, _ := selection.Find("a").First().Attr("href")
		views := selection.Find(".nb").First().Text()
		playlist := PlaylistTemp{
			title:     title,
			img:       img,
			url:       baseUrl + href,
			playCount: views,
		}
		playlistTempList = append(playlistTempList, playlist)
		log.I(TAG, title+" "+href+" "+img+" "+views)
	})
	nextPageHref, exists := document.Find(".zbtn, znxt").Last().Attr("href")
	if exists && nextPageHref != "" && !strings.Contains(nextPageHref, "javascript:void(0)") {
		reptilePlaylistTemp(baseUrl + nextPageHref)
	} else {
		return
	}
}

//将临时数据存入临时文件
func write2File() {
	var f *os.File
	if util.CheckFileIsExist(fileName) {
		os.Remove(fileName)
		f, _ = os.OpenFile(fileName, os.O_CREATE, 0666)
	} else {
		f, _ = os.Create(fileName)
	}
	defer f.Close()
	for _, playlist := range playlistTempList {
		var outPut = playlist.title + SEPARATOR + playlist.img + SEPARATOR + playlist.url + SEPARATOR + playlist.playCount + "\n"
		log.I(TAG, outPut)
		io.WriteString(f, outPut)
	}
}
