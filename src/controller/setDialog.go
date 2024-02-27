package controller

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
)

// SetAboutDialog 对话框
func SetAboutDialog(window fyne.Window) {
	title := "authorId"
	file, err := os.ReadFile("image/about/me.png")
	if err != nil {
		log.Println(err)
	}
	newIcon := widget.NewIcon(fyne.NewStaticResource("WX", file))
	newIcon.Resize(fyne.NewSize(300, 300))
	newIcon.Move(fyne.NewPos(30, 0))
	con := container.New(nil, newIcon)
	con.Resize(fyne.NewSize(400, 410))
	custom := dialog.NewCustom(title, "close", con, window)
	custom.Resize(fyne.NewSize(400, 410))
	custom.Show()
}
