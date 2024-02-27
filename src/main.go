package main

import (
	"os"
	"src/common"
	"src/screen"
)

func init() {
	//中文字体设置
	err := os.Setenv("FYNE_FONT", "fonts/SourceHanSansCN-VF-2.otf")
	if err != nil {
		println(err)
		return
	}
}

func main() {
	//首页
	screen.SetIndex()
	//软件运行
	common.MyApp.Run()
}
