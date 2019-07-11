package reptile

import "GoMusic163Reptile/src/util"

var (
	log              = util.Log{}
	TAG              = "Reptile"
	baseUrl          = "http://music.163.com"   //网易云BaseUrl
	fileName         = "./playlistTempFile.txt" //歌单数据临时文件
	SEPARATOR        = "<-->"                   //歌单数据字段间分隔符
	playlistTempList = make([]PlaylistTemp, 0)
	dbUserName       = "root"
	dbPassword       = "Qmkj@0622."
	dataSourceName   = dbUserName + ":" + dbPassword + "@tcp(39.98.55.88:10030)/music163_reptile?charset=utf8"
)
