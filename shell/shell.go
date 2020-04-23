package shell

import (
	"fmt"
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
	for _, v := range data {
		container.AddObject(v.Container)
	}
	return container
}

func Sort(data []*node.Node, window fyne.Window) {
	window.SetContent(newContainer(data))
	for gap := len(data) / 2; gap >= 1; gap /= 2 {
		if gap%2 == 0 {
			gap++
		}
		window.SetTitle(fmt.Sprintf("gap : %d", gap))
		for i := gap; i < len(data); i++ {
			key, j := data[i], i
			key.Fill(GREEN)
			window.SetContent(newContainer(data))
			time.Sleep(time.Millisecond * 150)
			for j >= gap && data[j-gap].Data > key.Data {
				if j != i {
					data[j].Fill(YELLOW)
				}
				data[j-gap].Fill(YELLOW)
				window.SetContent(newContainer(data))
				time.Sleep(time.Millisecond * 250)
				data[j] = data[j-gap]
				j -= gap
			}
			data[j] = key
			window.SetContent(newContainer(data))
			time.Sleep(time.Millisecond * 150)
			for _, v := range data {
				v.Fill(color.White)
			}
			window.SetContent(newContainer(data))
		}
	}
	window.SetTitle("Sort_Visualization")
	time.Sleep(time.Millisecond * 250)
}
