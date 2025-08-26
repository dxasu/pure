package gui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cast"
)

type Mover interface {
	Move(pos fyne.Position)
}

type Option struct {
	Tag  string
	X, Y int
	Fn   func(string)
}

func Move(m Mover, _o ...Option) {
	var o Option
	if len(_o) > 0 {
		o = _o[0]
	}
	if o.X == 0 && o.Y == 0 {
		return
	}
	m.Move(fyne.NewPos(float32(o.X), float32(o.Y)))
}

type Gui struct {
	fyne.Theme
	app fyne.App
}

type Win struct {
	Title string
	w, h  int
	objs  []fyne.CanvasObject
	Win   fyne.Window
}

func New() *Gui {
	g := new(Gui)
	g.app = app.New()
	g.Theme = myTheme{}
	if g.Theme != nil {
		g.app.Settings().SetTheme(g.Theme)
	}
	return g
}

func (g *Gui) GetApp() fyne.App {
	return g.app
}

func (g *Gui) NewWin(title string, w, h int) Win {
	window := Win{Title: title, w: w, h: h}
	myWindow := g.app.NewWindow(title)
	myWindow.Resize(fyne.NewSize(float32(w), float32(h)))
	window.Win = myWindow
	return window
}

func (w *Win) Run() {
	w.Win.SetContent(container.NewVBox(w.objs...))
	w.Win.ShowAndRun()
}

func (w *Win) AddButton(text string, _o ...Option) *widget.Button {
	var o Option
	if len(_o) > 0 {
		o = _o[0]
	}
	button := widget.NewButton(text, func() { o.Fn("") })
	Move(button, o)
	w.objs = append(w.objs, button)
	return button
}

func (w *Win) AddLable(text string, _o ...Option) *widget.Label {
	var o Option
	if len(_o) > 0 {
		o = _o[0]
	}
	label := widget.NewLabel(text)
	Move(label, o)
	w.objs = append(w.objs, label)
	return label
}

func (w *Win) AddInput(placeHolder string, multiLine bool, _o ...Option) *widget.Entry {
	var o Option
	if len(_o) > 0 {
		o = _o[0]
	}
	entry := widget.NewEntry()
	Move(entry, o)
	entry.MultiLine = multiLine
	if placeHolder != "" {
		entry.SetPlaceHolder(placeHolder)
	}
	entry.OnChanged = o.Fn
	w.objs = append(w.objs, entry)
	return entry
}

func (w *Win) AddRadio(namelist []string, _o ...Option) *widget.RadioGroup {
	var o Option
	if len(_o) > 0 {
		o = _o[0]
	}
	radio := widget.NewRadioGroup(namelist, o.Fn)
	Move(radio, o)
	w.objs = append(w.objs, radio)
	return radio
}

func (w *Win) AddCheck(text string, _o ...Option) *widget.Check {
	var o Option
	if len(_o) > 0 {
		o = _o[0]
	}
	check := widget.NewCheck(text, func(b bool) {
		if o.Fn != nil {
			o.Fn(cast.ToString(b))
		}
	})
	Move(check, o)
	w.objs = append(w.objs, check)
	return check
}

func (w *Win) Merge(count, col int) *fyne.Container {
	if count > len(w.objs) {
		panic(fmt.Errorf("too big for count: %d", count))
	}
	c := container.New(layout.NewGridLayout(col), w.objs[len(w.objs)-count:]...)
	nObjs := make([]fyne.CanvasObject, len(w.objs)-count+1)
	copy(nObjs, w.objs)
	nObjs[len(w.objs)-count] = c
	w.objs = nObjs
	return c
}

// Show mkdownContent as:
// "# Hello Fyne\n\nThis is **Markdown** support!"
func (w *Win) Show(title, mkdownContent string, buttonName ...string) string {
	renderer := widget.NewRichTextFromMarkdown(mkdownContent)
	l := make([]fyne.CanvasObject, 0, 2)
	l = append(l, renderer)
	selected := make(chan string)
	if len(buttonName) > 0 {
		buttons := make([]fyne.CanvasObject, 0)
		for _, v := range buttonName {
			buttons = append(buttons, widget.NewButton(v, func() {
				selected <- v
				w.Win.Close()
			}))
		}
		c := container.New(layout.NewGridLayout(len(buttonName)), buttons...)
		l = append(l, c)
	} else {
		close(selected)
	}
	w.Win.SetContent(container.NewVBox(l...))
	w.Win.CenterOnScreen()
	w.Win.Show()
	return <-selected
}

func (w *Win) ShowModel(text string) {
	content := widget.NewLabel(text)
	popup := widget.NewModalPopUp(content, w.Win.Canvas())
	popup.Show()
}

func (w *Win) ShowCalendar(d *time.Time) {
	w.Win.SetContent(widget.NewCalendar(time.Now(), func(date time.Time) {
		*d = date
	}))
}

func (w *Win) ShowForm(list []*widget.FormItem) ([]string, bool) {
	vals := make([]string, 0)
	form := &widget.Form{
		Items: list,
		OnSubmit: func() {
			for _, entry := range list {
				if e, ok := entry.Widget.(*widget.Entry); ok {
					vals = append(vals, e.Text)
				}
			}
			w.Win.Close()
		},
		OnCancel: func() {
			w.Win.Close()
		},
		SubmitText: "Ok",
		CancelText: "Cancel",
	}
	w.Win.SetContent(form)
	w.Win.ShowAndRun()
	return vals, len(vals) > 0
}
