package node

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type Node struct {
	Data int

	Rects     []*canvas.Rectangle
	Container *fyne.Container
}

func NewNode(data int, length int) *Node {
	node := &Node{}
	rects := []*canvas.Rectangle{}
	container := fyne.NewContainerWithLayout(layout.NewGridLayoutWithRows(length + 2))
	for i := 0; i < length-data; i++ {
		container.AddObject(layout.NewSpacer())
	}
	for i := 0; i < data; i++ {
		rect := canvas.NewRectangle(color.White)
		container.AddObject(rect)
		rects = append(rects, rect)
	}
	label := widget.NewLabel(strconv.Itoa(data))
	label.Alignment = fyne.TextAlignCenter
	container.AddObject(label)
	node.Data = data
	node.Rects = rects
	node.Container = container
	return node
}

func (node *Node) Fill(c color.Color) {
	for _, rect := range node.Rects {
		rect.FillColor = c
	}
}

func (node *Node) String() string {
	return fmt.Sprint(node.Data)
}
