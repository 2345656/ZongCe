package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
)

// Config 设计软件配置
type Config struct {
	App       fyne.App
	Win       fyne.Window
	Container *fyne.Container
}

var MyApp = &Config{}

// 屏幕大小
const indexScreenHight = 500
const indexScreenWeight = 800

func SetIndex() {
	MyApp.App = app.New()
	MyApp.Win = MyApp.App.NewWindow("校易综测")
	//制作软件界面
	MyApp.setWinScreen()
	//制作界面组件
	MyApp.setWinWidget()
	//功能栏
	MyApp.setMenu()
	//窗口展示
	MyApp.Win.Show()
}

// 制作软件界面
func (MyWin *Config) setWinScreen() {
	//1.将界面堆好
	MyWin.Win.SetIcon(setIcon())
	MyWin.Win.Resize(fyne.Size{
		Width:  float32(indexScreenWeight),
		Height: float32(indexScreenHight),
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
		setOneFileIndex()
	})
	//多文件类型
	moreFBtn := widget.NewButtonWithIcon("多文件", nil, func() {
		setMoreFileIndex()
	})
	//设置
	setBtn := widget.NewButtonWithIcon("设置", nil, func() {
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
	aboutMe := fyne.NewMenuItem("联系方式", func() {
		setAboutDialog(MyWin.Win)
	})

	about := fyne.NewMenu("关于", aboutMe)

	mainMenu := fyne.NewMainMenu(about)

	MyWin.Win.SetMainMenu(mainMenu)
}

// 对话框
func setAboutDialog(window fyne.Window) {
	title := "本人联系方式"
	file, err := os.ReadFile("image/about/me.png")
	if err != nil {
		log.Println(err)
	}
	newIcon := widget.NewIcon(fyne.NewStaticResource("微信", file))
	newIcon.Resize(fyne.NewSize(300, 300))
	newIcon.Move(fyne.NewPos(30, 0))
	con := container.New(nil, newIcon)
	con.Resize(fyne.NewSize(400, 410))
	custom := dialog.NewCustom(title, "关闭", con, window)
	custom.Resize(fyne.NewSize(400, 410))
	custom.Show()
}
