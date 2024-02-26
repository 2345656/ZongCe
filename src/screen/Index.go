package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Config 设计软件配置
type Config struct {
	App       fyne.App
	Win       fyne.Window
	Container *fyne.Container
}

var MyApp = &Config{}

func SetIndex() {
	MyApp.App = app.New()
	MyApp.Win = MyApp.App.NewWindow("校易综测")
	//制作软件界面
	MyApp.setWinScreen()
	//制作界面组件
	MyApp.setWinWidget()
	//窗口展示
	MyApp.Win.Show()
}

// 制作软件界面
func (MyWin *Config) setWinScreen() {
	//1.将界面堆好
	MyWin.Win.SetIcon(SetIcon())
	MyWin.Win.Resize(fyne.Size{
		Width:  1000,
		Height: 700,
	})
	MyWin.Win.CenterOnScreen()
	MyWin.Win.SetFixedSize(true)
}

/*
*制作界面组件（三个按钮
*按钮一：单文件类型
*按钮二：多文件类型
*按钮三：设置按钮
 */
func (MyWin *Config) setWinWidget() {
	//按钮制作
	//单文件类型
	oneFBtn := widget.NewButtonWithIcon("单文件", nil, func() {

	})
	oneFBtn.Resize(fyne.NewSize(100, 100))
	oneFBtn.Move(fyne.NewPos(100, 100))
	//多文件类型
	moreFBtn := widget.NewButtonWithIcon("多文件", nil, func() {

	})
	moreFBtn.Resize(fyne.NewSize(100, 100))
	moreFBtn.Move(fyne.NewPos(300, 100))
	//设置
	setBtn := widget.NewButtonWithIcon("设置", nil, func() {

	})
	setBtn.Resize(fyne.NewSize(100, 100))
	setBtn.Move(fyne.NewPos(400, 100))
	//容器设置
	MyWin.Container = container.New(nil, oneFBtn, moreFBtn, setBtn)
	MyWin.Win.SetContent(MyWin.Container)
}
