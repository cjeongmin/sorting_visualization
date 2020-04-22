package bubble

import (
	"image/color"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"github.com/cjeongmin/sort_visualization/node"
)

var (
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
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			data[j].Fill(GREEN)
			data[j+1].Fill(YELLOW)
			window.SetContent(newContainer(data))
			time.Sleep(time.Millisecond * 250)
			if data[j+1].Data < data[j].Data {
				data[j+1], data[j] = data[j], data[j+1]
				window.SetContent(newContainer(data))
				time.Sleep(time.Millisecond * 250)
			}
			data[j].Fill(color.White)
			data[j+1].Fill(color.White)
		}
	}
	data[0].Fill(color.White)
	data[1].Fill(color.White)
	window.SetContent(newContainer(data))
}
