package main

import (
	"image"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("Hello World")
	window.Resize(fyne.NewSize(500, 400))
	//label
	// label := widget.NewLabel("New Content")
	label := canvas.NewText("Hello Gopher", color.RGBA{R: 0, G: 255, B: 255, A: 255})
	label.Alignment = fyne.TextAlignCenter // テキストの中央揃え
	label.TextSize = 24

	button := widget.NewButton("Hello button", func() {})
	buttonContainer := container.New(layout.NewCenterLayout(), button)
	buttonContainer.Resize(fyne.NewSize(300, 60)) // ボタンのサイズを設定

	//画像を開く
	file, err := os.Open("./go.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//画像をでコード
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	fyneImage := canvas.NewImageFromImage(img)
	// fyneImage.FillMode = canvas.ImageFillContain
	fyneImage.FillMode = canvas.ImageFillContain
	fyneImage.SetMinSize(fyne.NewSize(200, 200)) // 必要に応じて画像の最小サイズを設定

	vBox := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		label,
		fyneImage,
		layout.NewSpacer(),
		buttonContainer,
		layout.NewSpacer(),
	)
	window.SetContent(vBox)
	window.ShowAndRun()
}
