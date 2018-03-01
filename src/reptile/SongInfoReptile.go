package reptile

import (
	"reptile/src/net"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"strconv"
	"database/sql"
)

type Song struct {
	id           int
	title        string
	subtitle     string
	img          string
	artist       []string
	artistId     []int
	album        string
	albumId      int
	commentCount int
	lyric        string
}

/**
 * 从数据库查询所有歌单的歌曲id
 */
func QueryAllPlaylistSongIdFromDB() ([]int, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.E(TAG, err)
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT DISTINCT songId FROM tbl_163music_playlist_song_info")
	if err != nil {
		log.E(TAG, err)
		return nil, err
	}
	defer rows.Close()
	ids := make([]int, 0)
	for rows.Next() {
		var songId int
		err := rows.Scan(&songId)
		if err != nil {
			log.E(TAG, err)
		} else {
			ids = append(ids, songId)
		}
	}
	return ids, nil
}

/**
 * 抓取歌曲数据
 */
func ReptileSongById(songIds []int) ([]Song, error) {
	var songs = make([]Song, 0)
	for index, songId := range songIds {
		log.I(TAG, "当前进度", index, "of", len(songIds))
		url := "http://music.163.com/song?id=" + strconv.Itoa(songId)
		log.I(url)
		reader, err := net.GetRequestForReader(url)
		if err != nil {
			log.E(TAG, err)
			return nil, err
		}
		document, _ := goquery.NewDocumentFromReader(reader)
		log.I(document.Html())
		id := songId
		title := document.Find(".f-ff2").First().Text()
		subtitle := document.Find(".subtit.f-fs1.f-ff2").First().Text()
		img, _ := document.Find(".j-img").First().Attr("src")
		var artist []string
		var artistId []int

		document.Find(".des.s-fc4").First().Find("a").Each(func(index int, selection *goquery.Selection) {
			artist = append(artist, selection.Text())
			href, _ := selection.Attr("href")
			idStr := strings.Split(href, "id=")[1]
			id, _ := strconv.Atoi(idStr)
			artistId = append(artistId, id)
		})
		log.I("", document.Find(".des.s-fc4").First().Next().Text())
		album := document.Find(".des.s-fc4").First().Next().Find("a").Text()
		href, _ := document.Find(".des.s-fc4").First().Next().Find("a").Attr("href")
		albumIdStr := strings.Split(href, "id=")[1]
		albumId, _ := strconv.Atoi(albumIdStr)
		//todo commentCount lyric 获取
		commentCount := -1
		lyric := "todo"
		/*commentCountStr := document.Find("#cnt_comment_count").Text()
		commentCount, _ := strconv.Atoi(commentCountStr)
		lyric := document.Find("#lyric-content").Text()*/
		song := Song{
			id:           id,
			title:        title,
			subtitle:     subtitle,
			img:          img,
			artist:       artist,
			artistId:     artistId,
			album:        album,
			albumId:      albumId,
			commentCount: commentCount,
			lyric:        lyric,
		}
		songs = append(songs, song)
		log.I(TAG, song)
	}
	return songs, nil
}

/**
 * 保存歌曲到数据库
 */
func SaveSong2DB(songs []Song) error {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.E(TAG, err)
		return err
	}
	defer db.Close()
	for index, song := range songs {
		log.I(TAG, "当前进度:", index)
		var artist string
		for i, singleArtist := range song.artist {
			if i == 0 {
				artist = artist + singleArtist
			} else {
				artist = artist + "/" + singleArtist
			}
		}
		result, err := db.Exec("REPLACE INTO tbl_163music_song_info VALUES (?,?,?,?,?,?,?,?,?)", song.id, song.title, song.subtitle, song.img, artist, song.album, song.albumId, song.commentCount, song.lyric)
		//todo 保存albumId到数据库
		if err != nil {
			log.E(TAG, err)
		} else {
			log.I(TAG, result)
		}
	}
	return nil
}
