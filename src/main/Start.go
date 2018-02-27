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
	reptile.ReptilePlaylistTemp()
	//进入详情页抓取详细数据
	reptile.ReptilePlaylistInfo()
	log.I(TAG, "end")
}

