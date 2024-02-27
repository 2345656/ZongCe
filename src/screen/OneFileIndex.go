package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// 单文件界面的容器
var oneFileContent *fyne.Container

func setOneFileIndex() {
	//界面不需要修改
	// 清除原有内容
	MyWin.Win.Content().(*fyne.Container).Objects = nil
	// 设置单文件界面的组件
	oneFileContent = createOneFileContent()
	MyWin.Win.SetContent(oneFileContent)
}

// 设置单文件界面的组件
func createOneFileContent() *fyne.Container {
	//组件制作
	imFileBtn := widget.NewButtonWithIcon("", theme.FileIcon(), func() {
		//按钮功能
	})
	imFileBtn.Move(fyne.NewPos(200, 20))
	imFileBtn.Resize(fyne.NewSize(30, 30))

	//容器
	con := container.New(nil, imFileBtn)
	return con
}
