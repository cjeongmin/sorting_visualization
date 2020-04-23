package gui

import (
	"math/rand"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/cjeongmin/sort_visualization/bubble"
	"github.com/cjeongmin/sort_visualization/insertion"
	"github.com/cjeongmin/sort_visualization/node"
	"github.com/cjeongmin/sort_visualization/selection"
	"github.com/cjeongmin/sort_visualization/shell"
)

type Visualization struct {
	fns     map[string]func(data []*node.Node, window fyne.Window)
	sorting string

	app    fyne.App
	window fyne.Window
	data   []*node.Node
}

func NewVisualization() *Visualization {
	app := app.New()
	window := app.NewWindow("Sort_Visualization")
	fns := map[string]func(data []*node.Node, window fyne.Window){
		"Bubble":    bubble.Sort,
		"Selection": selection.Sort,
		"Insertion": insertion.Sort,
		"Shell":     shell.Sort,
	}
	return &Visualization{fns, "", app, window, []*node.Node{}}
}

func (v *Visualization) shuffle() {
	v.data = make([]*node.Node, 20)
	for i := range v.data {
		v.data[i] = node.NewNode(i+1, 20)
	}
	rand.Shuffle(20, func(i, j int) {
		v.data[i], v.data[j] = v.data[j], v.data[i]
	})
}

func (v *Visualization) display() {
	container := fyne.NewContainerWithLayout(layout.NewGridLayoutWithRows(2))
	container.AddObject(widget.NewButton("Start", func() {
		if v.sorting == "" {
			return
		}
		v.shuffle()
		v.fns[v.sorting](v.data, v.window)
		v.sorting = ""
		v.display()
	}))
	container.AddObject(widget.NewSelect([]string{"Bubble", "Selection", "Insertion", "Shell"}, func(s string) {
		v.sorting = s
	}))
	v.window.Resize(fyne.NewSize(480, 360))
	v.window.SetContent(container)
}

func (v *Visualization) ShowAndRun() {
	v.window.SetFixedSize(true)
	v.display()
	v.window.ShowAndRun()
}
