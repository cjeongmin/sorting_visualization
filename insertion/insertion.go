package insertion

import (
	"image/color"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"github.com/cjeongmin/sort_visualization/node"
)

var (
	RED    = color.RGBA{R: 255, G: 0, B: 0}
	GREEN  = color.RGBA{R: 0, G: 255, B: 0}
	YELLOW = color.RGBA{R: 255, G: 255, B: 0}
)

func newContainer(data []*node.Node) *fyne.Container {
	container := fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(len(data)))
	for _, node := range data {
		container.AddObject(node.Container)
	}
	return container
}

func Sort(data []*node.Node, window fyne.Window) {
	window.SetContent(newContainer(data))
	for i := 1; i < len(data); i++ {
		key, j := data[i], i-1
		key.Fill(GREEN)
		window.SetContent(newContainer(data))
		time.Sleep(time.Millisecond * 100)
		for j >= 0 && data[j].Data > key.Data {
			data[j].Fill(YELLOW)
			window.SetContent(newContainer(data))
			time.Sleep(time.Millisecond * 250)
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
		window.SetContent(newContainer(data))
		time.Sleep(time.Millisecond * 100)
		for _, v := range data {
			v.Fill(color.White)
		}
		window.SetContent(newContainer(data))
	}
	time.Sleep(time.Second)
}
