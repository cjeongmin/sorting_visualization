package gui

import (
	"math/rand"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/cjeongmin/sort_visualization/bubble"
	"github.com/cjeongmin/sort_visualization/node"
	"github.com/cjeongmin/sort_visualization/selection"
)

type Visualization struct {
	sorting string

	app    fyne.App
	window fyne.Window
	data   []*node.Node
}

func NewVisualization() *Visualization {
	app := app.New()
	window := app.NewWindow("Sort_Visualization")
	window.SetFixedSize(false)
	return &Visualization{"", app, window, []*node.Node{}}
}

func (v *Visualization) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	length := int(rand.Int31n(20-15) + 15)
	v.data = make([]*node.Node, length)
	for i := range v.data {
		v.data[i] = node.NewNode(i+1, length)
	}
	rand.Shuffle(length, func(i, j int) {
		v.data[i], v.data[j] = v.data[j], v.data[i]
	})
}

func (v *Visualization) Display() {
	container := fyne.NewContainerWithLayout(layout.NewGridLayoutWithRows(2))
	container.AddObject(widget.NewButton("Start", func() {
		v.Shuffle()
		switch v.sorting {
		case "Bubble":
			bubble.Sort(v.data, v.window)
			v.sorting = ""
			v.Display()
		case "Select":
			selection.Sort(v.data, v.window)
			v.sorting = ""
			v.Display()
		case "Insert":

			v.sorting = ""
			v.Display()
		}
	}))
	container.AddObject(widget.NewSelect([]string{"Bubble", "Select", "Insert"}, func(s string) {
		v.sorting = s
	}))
	v.window.Resize(fyne.NewSize(480, 360))
	v.window.SetContent(container)
}

func (v *Visualization) ShowAndRun() {
	v.Display()
	v.window.ShowAndRun()
}
