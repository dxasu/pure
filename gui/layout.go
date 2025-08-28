package gui

import "fyne.io/fyne/v2"

type zeroSpaceHBoxLayout struct{}

func (z *zeroSpaceHBoxLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	x := float32(0) // 起始位置
	for _, obj := range objects {
		obj.Move(fyne.NewPos(x, 0))
		obj.Resize(fyne.NewSize(obj.MinSize().Width, size.Height))
		x += obj.MinSize().Width // 无间距累加
	}
}

func (z *zeroSpaceHBoxLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	width, height := float32(0), float32(0)
	for _, obj := range objects {
		width += obj.MinSize().Width
		height = fyne.Max(height, obj.MinSize().Height)
	}
	return fyne.NewSize(width, height)
}
