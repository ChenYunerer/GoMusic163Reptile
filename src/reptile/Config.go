package reptile

import "reptile/src/util"

var (
	log              = util.Log{}
	TAG              = "Reptile"
	baseUrl          = "http://music.163.com"   //网易云BaseUrl
	fileName         = "./playlistTempFile.txt" //歌单数据临时文件
	SEPARATOR        = "<-->"                   //歌单数据字段间分隔符
	playlistTempList = make([]PlaylistTemp, 0)
	dbUserName       = "root"
	dbPassword       = "12345cyCY"
	dataSourceName   = dbUserName + ":" + dbPassword + "@tcp(rm-2ze5sllf9zl6e9j39uo.mysql.rds.aliyuncs.com:3306)/opern?charset=utf8"
)
