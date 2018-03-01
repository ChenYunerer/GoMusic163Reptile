package main

import (
	"reptile/src/util"
	"reptile/src/reptile"
)

var (
	log              = util.Log{}
	TAG              = "MAIN"
)

func main() {
	log.I(TAG, "start")
	//先抓取临时数据
	//reptile.ReptilePlaylistTemp()
	//进入详情页抓取详细数据
	//reptile.ReptilePlaylistInfo()
	//获取所有歌单的歌曲id
	songIds, err := reptile.QueryAllPlaylistSongIdFromDB()
	if err != nil {
		log.E(TAG, err)
		return
	}
	//抓取歌曲信息
	songs, err := reptile.ReptileSongById(songIds)
	if err != nil {
		log.E(TAG, err)
		return
	}
	//保存歌曲信息到数据库
	err = reptile.SaveSong2DB(songs)
	if err != nil {
		log.E(TAG, err)
	}
	log.I(TAG, "end")
}

