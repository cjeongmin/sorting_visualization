package selection

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
	for i := 0; i < len(data)-1; i++ {
		minIdx := i
		data[i].Fill(GREEN)
		for j := i + 1; j < len(data); j++ {
			data[j].Fill(RED)
			if data[minIdx].Data > data[j].Data {
				if minIdx != i {
					data[minIdx].Fill(color.White)
				}
				minIdx = j
				data[minIdx].Fill(YELLOW)
			}
			window.SetContent(newContainer(data))
			time.Sleep(time.Millisecond * 250)
			if j != minIdx {
				data[j].Fill(color.White)
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
		time.Sleep(time.Millisecond * 250)
		data[i].Fill(color.White)
		data[minIdx].Fill(color.White)
	}
	for _, node := range data {
		node.Fill(color.White)
	}
	window.SetContent(newContainer(data))
	time.Sleep(time.Millisecond * 250)
}
