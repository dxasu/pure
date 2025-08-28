package main

import (
	"fmt"

	"github.com/dxasu/pure/gui"
)

func main() {
	app := gui.New()
	w := app.NewWin("2025", 300, 200)
	w.AddLable("welcome to here")
	input := w.AddInput("please input", false)
	w.AddCheck("check")
	o := []string{"选项1", "选项2", "3"}
	w.AddRadio(o, func(s string) {
		fmt.Println(s)
	})

	w.AddSelect(o, func(s string) {
		fmt.Println(s)
	})

	// w.AddSelectSearch(o, func(s string) {
	// 	fmt.Println("search:", s)
	// })

	// w.AddFileTree([]string{".go", ".mod"}, func(s string) {
	// 	fmt.Println(s)
	// })

	w.AddButton("calc", func(string) {
		input.Text = "good"
		input.Refresh()
		// txt := w.Show("question", "# Hello Fyne\n\nThis is **Markdown** support!", "cancel", "??", "ok")
		// w.Show("result", txt)
		// w.ShowModel("hello world")

		w.ShowToast("hi,#ff0000hello#world.#00ff00ni好#to swiming#0000ffgo哈哈go😊#normal#toast")
		// w.ShowCalendar(func(t time.Time) {
		// 	fmt.Println("select date:", t)

		// w.ShowFileTree([]string{".go", ".mod"}, func(s string) {
		// 	fmt.Println("select file:", s)
		// })
		// }()
	})

	// d, ok := w.ShowForm([]*widget.FormItem{
	// 	{Text: "1", Widget: widget.NewEntry()},
	// 	{Text: "2", Widget: widget.NewEntry()},
	// 	{Text: "3", Widget: widget.NewMultiLineEntry()},
	// })
	// if ok {
	// 	fmt.Println(d)
	// }
	w.Run()
}
