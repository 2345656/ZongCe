package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"src/common"
	"src/controller"
	"src/interfaces"
)

// Config 设计软件配置
type Config struct {
	interfaces.SetWin
	Win       fyne.Window
	Container *fyne.Container
}

var MyWin = &Config{}

// 屏幕大小
const indexScreenHeight = 500
const indexScreenWeight = 800

func SetIndex() {
	MyWin.Win = common.MyApp.NewWindow("校易综测")
	//制作软件界面
	MyWin.setWinScreen()
	//制作界面组件
	MyWin.setWinWidget()
	//功能栏
	MyWin.setMenu()
	//窗口展示
	MyWin.Win.Show()
}

// 制作软件界面
func (MyWin *Config) setWinScreen() {
	//1.将界面堆好
	MyWin.Win.SetIcon(setIcon())
	MyWin.Win.Resize(fyne.Size{
		Width:  float32(indexScreenWeight),
		Height: float32(indexScreenHeight),
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
	oneFBtn := widget.NewButtonWithIcon("File单", nil, func() {
		setOneFileIndex()
	})
	//多文件类型
	moreFBtn := widget.NewButtonWithIcon("Files", nil, func() {
		setMoreFileIndex()
	})
	//设置
	setBtn := widget.NewButtonWithIcon("Setting", nil, func() {
		settingIndex()
	})
	//添加进组件数组
	widgets := []fyne.CanvasObject{oneFBtn, moreFBtn, setBtn}

	//容器设置
	MyWin.Container = container.NewCenter(container.NewGridWrap(fyne.NewSize(100, 100), widgets...))
	MyWin.Win.SetContent(MyWin.Container)
}

// SetMenu 功能栏
func (MyWin *Config) setMenu() {
	//小菜单
	aboutMe := fyne.NewMenuItem("info", func() {
		controller.SetAboutDialog(MyWin.Win)
	})

	about := fyne.NewMenu("About", aboutMe)

	mainMenu := fyne.NewMainMenu(about)

	MyWin.Win.SetMainMenu(mainMenu)
}
