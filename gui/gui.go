package gui

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	widgetX "fyne.io/x/fyne/widget"
	"github.com/spf13/cast"
)

type Gui struct {
	fyne.Theme
	app fyne.App
}

type Win struct {
	Title    string
	w, h     int
	objs     []fyne.CanvasObject
	Win      fyne.Window
	startPos int // for merge start pos
	gui      *Gui
}

type Callback func(string)

type callPara []Callback

func (_c callPara) fn() Callback {
	var c Callback = func(string) {}
	if len(_c) > 0 {
		c = _c[0]
	}
	return c
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
	window.gui = g
	return window
}

func (w *Win) Run() {
	w.Win.SetContent(container.NewVBox(w.objs...))
	w.Win.ShowAndRun()
}

func (w *Win) AddButton(text string, _c ...Callback) *widget.Button {
	button := widget.NewButton(text, func() { callPara(_c).fn()("") })
	w.objs = append(w.objs, button)
	return button
}

func (w *Win) AddSelect(options []string, _c ...Callback) *widget.Select {
	selectW := widget.NewSelect(options, callPara(_c).fn())
	w.objs = append(w.objs, selectW)
	return selectW
}

// func (w *Win) AddSelectSearch(options []string, _c ...Callback) *fyne.Container {
// 	combo := widgetX.NewNumericalEntry().SelectedText(items, func(s string) {})
// 	w.objs = append(w.objs, c)
// 	return c
// }

func (w *Win) AddSelectMulti(options []string, _c ...Callback) *widget.CheckGroup {
	multiSelect := widget.NewCheckGroup(options, func(selected []string) {
		callPara(_c).fn()(strings.Join(selected, ","))
	})
	w.objs = append(w.objs, multiSelect)
	return multiSelect
}

func (w *Win) AddLable(text string) *widget.Label {
	label := widget.NewLabel(text)
	w.objs = append(w.objs, label)
	return label
}

func (w *Win) AddInput(placeHolder string, multiLine bool, _c ...Callback) *widget.Entry {
	entry := widget.NewEntry()
	entry.MultiLine = multiLine
	if placeHolder != "" {
		entry.SetPlaceHolder(placeHolder)
	}
	entry.OnChanged = callPara(_c).fn()
	w.objs = append(w.objs, entry)
	return entry
}

func (w *Win) AddRadio(namelist []string, _c ...Callback) *widget.RadioGroup {
	c := func(string) {}
	if len(_c) > 0 {
		c = _c[0]
	}
	radio := widget.NewRadioGroup(namelist, c)
	w.objs = append(w.objs, radio)
	return radio
}

func (w *Win) AddCheck(text string, _c ...Callback) *widget.Check {
	check := widget.NewCheck(text, func(b bool) {
		callPara(_c).fn()(cast.ToString(b))
	})
	w.objs = append(w.objs, check)
	return check
}

func (w *Win) AddSlider(min, max float64, _c ...Callback) *widget.Slider {
	slider := widget.NewSlider(min, max)
	slider.OnChanged = func(f float64) {
		callPara(_c).fn()(cast.ToString(f))
	}
	w.objs = append(w.objs, slider)
	return slider
}

func (w *Win) AddTimePicker(t time.Time, _c ...Callback) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("2006-01-02 15:04:05")
	entry.Text = t.Format("2006-01-02 15:04:05")
	entry.OnChanged = callPara(_c).fn()
	w.objs = append(w.objs, entry)
	return entry
}

func (w *Win) AddDatePicker(t time.Time, _c ...Callback) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("2006-01-02")
	entry.Text = t.Format("2006-01-02")
	entry.OnChanged = callPara(_c).fn()
	w.objs = append(w.objs, entry)
	return entry
}

func (w *Win) AddProgressBar() *widget.ProgressBar {
	bar := widget.NewProgressBar()
	w.objs = append(w.objs, bar)
	return bar
}

func (w *Win) AddProgressBarInfinite() *widget.ProgressBarInfinite {
	bar := widget.NewProgressBarInfinite()
	w.objs = append(w.objs, bar)
	return bar
}

func (w *Win) AddSeparator() *widget.Separator {
	sep := widget.NewSeparator()
	w.objs = append(w.objs, sep)
	return sep
}

func (w *Win) AddNewLine() *widget.Label {
	l := widget.NewLabel(" ")
	w.objs = append(w.objs, l)
	return l
}

func (w *Win) AddCalendar(f func(time.Time)) *widgetX.Calendar {
	// var c *widget.Calendar
	// c = widget.NewCalendar(time.Now(), func(t time.Time) {
	// 	f(t)
	// 	c.Hide()
	// })
	var c *widgetX.Calendar
	c = widgetX.NewCalendar(time.Now(), func(t time.Time) {
		f(t)
		c.Hide()
		c.Refresh()
	})
	w.objs = append(w.objs, c)
	return c
}

func (w *Win) MergeMark() {
	w.startPos = len(w.objs)
}

func (w *Win) Merge(col int) *fyne.Container {
	count := len(w.objs) - w.startPos
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
