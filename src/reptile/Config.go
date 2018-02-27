package reptile

import "reptile/src/util"

var (
	log              = util.Log{}
	TAG              = "Reptile"
	baseUrl          = "http://music.163.com"   //网易云BaseUrl
	fileName         = "./playlistTempFile.txt" //歌单数据临时文件
	SEPARATOR        = "<-->"                   //歌单数据字段间分隔符
	playlistTempList = make([]PlaylistTemp, 0)
	dataSourceName   = ""
)
