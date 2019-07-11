package reptile

import (
	"GoMusic163Reptile/src/net"
	"GoMusic163Reptile/src/util"
	"bufio"
	"database/sql"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type PlaylistInfo struct {
	id           int
	title        string
	img          string
	url          string
	uploadTime   string
	playCount    int
	commentCount int
	tags         []string
	description  string
	musicIds     []int
}

var (
	playlistInfoList = make([]PlaylistInfo, 0)
	wg               = sync.WaitGroup{}
)

//爬取歌单详细数据
func ReptilePlaylistInfo() {
	readFromFile()
	threadNum := 12
	a := len(playlistTempList) / threadNum
	for i := 0; i < threadNum; i++ {
		var c []PlaylistTemp
		if i == threadNum-1 {
			c = playlistTempList[i*a:]
		} else {
			c = playlistTempList[i*a : (i+1)*a]
		}

		go func() {
			wg.Add(1)
			defer wg.Done()
			reptilePlaylistInfo(c)
		}()
	}
	wg.Wait()
	savePlaylist2DB()
}

//爬取歌单数据
func reptilePlaylistInfo(date []PlaylistTemp) {
	for index, playlist := range date {
		log.I(TAG, index)
		reader, err := net.GetRequestForReader(playlist.url)
		if err != nil {
			log.I(TAG, err.Error())
		}
		document, _ := goquery.NewDocumentFromReader(reader)
		idStr, _ := document.Find("#content-operation").Attr("data-rid")
		id, _ := strconv.Atoi(idStr)
		title := document.Find(".f-ff2.f-brk").First().Text()
		img, _ := document.Find(".j-img").First().Attr("src")
		url := playlist.url
		var uploadTime string
		jsonStr := document.Find("script").First().Text()
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
			uploadTime = dat["pubDate"].(string)
		} else {
			uploadTime = ""
			log.E(TAG, err)
		}
		playCountStr := document.Find("#play-count").Text()
		playCount, _ := strconv.Atoi(playCountStr)
		commentCountStr := document.Find("#cnt_comment_count").Text()
		commentCount, _ := strconv.Atoi(commentCountStr)
		var tags []string
		document.Find(".u-tag").Each(func(i int, selection *goquery.Selection) {
			tags = append(tags, selection.Children().First().Text())
		})
		description := document.Find("#album-desc-more").Text()
		var musicIds []int
		document.Find("#song-list-pre-cache").Find("a").Each(func(i int, selection *goquery.Selection) {
			musicHref, _ := selection.Attr("href")
			splits := strings.Split(musicHref, "id=")
			musicIdStr := splits[1]
			musicId, _ := strconv.Atoi(musicIdStr)
			musicIds = append(musicIds, musicId)
		})
		playlistInfo := PlaylistInfo{
			id:           id,
			title:        title,
			img:          img,
			url:          url,
			uploadTime:   uploadTime,
			playCount:    playCount,
			commentCount: commentCount,
			tags:         tags,
			description:  description,
			musicIds:     musicIds,
		}
		playlistInfoList = append(playlistInfoList, playlistInfo)
		log.I(TAG, playlistInfo)
	}
	log.I(TAG, len(playlistTempList))
}

//保存到数据库
func savePlaylist2DB() {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.E(TAG, err)
		return
	}
	defer db.Close()
	for index, item := range playlistInfoList {
		log.I(TAG, "当前数据index", index)
		res, err := db.Exec("INSERT INTO music163_playlist_info_tbl VALUES (?,?,?,?,?,?,?,?)", item.id, item.title, item.img, item.url, item.uploadTime, item.playCount, item.commentCount, item.description)
		if err != nil {
			log.E(TAG, err)
			continue
		}
		for _, tag := range item.tags {
			_, err := db.Exec("INSERT INTO music163_playlist_tag_info_tbl(playListId, playListTag) VALUES (?,?)", item.id, tag)
			if err != nil {
				log.E(TAG, err)
			}
		}
		for _, musicId := range item.musicIds {
			_, err := db.Exec("INSERT INTO music163_playlist_song_info_tbl(playListId, songId) VALUES (?,?)", item.id, musicId)
			if err != nil {
				log.E(TAG, err)
			}
		}
		if err != nil {
			log.E(TAG, err)
		} else {
			row, _ := res.RowsAffected()
			log.I(TAG, row, "tags len:", len(item.tags), "songs len:", len(item.musicIds))
		}
	}
}

//从文件读取歌单临时数据
func readFromFile() {
	playlistTempList = make([]PlaylistTemp, 0)
	if !util.CheckFileIsExist(fileName) {
		return
	}
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		log.I(TAG, line)
		splits := strings.Split(line, SEPARATOR)
		playlistTemp := PlaylistTemp{
			title:     splits[0],
			img:       splits[1],
			url:       splits[2],
			playCount: splits[3],
		}
		playlistTempList = append(playlistTempList, playlistTemp)
		if len(playlistTempList) == 20 {
			return
		}
	}
}
