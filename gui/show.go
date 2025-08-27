package gui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	widgetX "fyne.io/x/fyne/widget"
)

func (w *Win) ShowCalendar(onSelected func(time.Time)) {
	w.ShowSubWindow("select date", func(sub fyne.Window) fyne.CanvasObject {
		sub.Resize(w.Win.Canvas().Size())
		return widget.NewCalendar(time.Now(), func(date time.Time) {
			onSelected(date)
			sub.Hide()
		})
	})
}

// ShowFileTree adds a file tree to the window with the specified file extensions filter.
// For example, to show only .txt and .md files, use filter []string{".txt", ".md"}.
func (w *Win) ShowFileTree(filter []string, _c ...Callback) {
	w.ShowSubWindow("select file", func(sub fyne.Window) fyne.CanvasObject {
		sub.Resize(w.Win.Canvas().Size())
		tree := widgetX.NewFileTree(storage.NewFileURI("."))
		tree.Filter = storage.NewExtensionFileFilter(filter) // Filter files
		tree.Sorter = func(u1, u2 fyne.URI) bool {
			return u1.String() < u2.String()
		}
		tree.OnSelected = func(uid widget.TreeNodeID) {
			callPara(_c).fn()(uid)
			sub.Close()
		}
		return tree
	})
}

func (w *Win) ShowModel(text string) {
	// 显示模态弹窗（自动禁用父窗口）
	modal := dialog.NewCustom("tip", "close", widget.NewLabel(text), w.Win)
	s := w.Win.Canvas().Size()
	modal.Resize(fyne.NewSize(s.Width*3/4, s.Height*3/4))
	modal.Show()
}

func (w *Win) ShowSubWindow(title string, fn func(fyne.Window) fyne.CanvasObject) {
	parent := w.Win
	subWin := w.gui.app.NewWindow(title)
	if parent != nil {
		subWin.SetCloseIntercept(func() {
			subWin.Hide()
		})
	}
	subWin.SetContent(fn(subWin))
	if parent != nil {
		parentPos := parent.Content().Position()
		subWin.Content().Move(parentPos.Add(fyne.NewPos(
			(parent.Canvas().Size().Width)/2,  // 水平居中
			(parent.Canvas().Size().Height)/2, // 垂直居中
		)))
	}
	subWin.Show()
}

func (w *Win) ShowError(err error) {
	dialog.ShowError(err, w.Win)
}

func (w *Win) ShowToast(message string) {
	NewToast(w.Win).QueueToast(message)
}
