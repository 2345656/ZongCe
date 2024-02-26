package main

import (
	"os"
	"src/screen"
)

func main() {
	//字体设置
	os.Setenv("FYNE_FONT", "fonts/cjkFonts_allseto_v1.11.ttf")
	//首页
	screen.SetIndex()
	//软件运行
	screen.MyApp.App.Run()
}
