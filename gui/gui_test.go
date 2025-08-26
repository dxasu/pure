package gui

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func TestWin_task(td *testing.T) {
	app := New()
	w := app.NewWin("2025", 300, 200)
	w.AddLable("welcome to here")
	input := w.AddInput("please input", false)
	w.AddCheck("check")
	w.AddRadio([]string{"选项1", "选项2"}, Option{Fn: func(s string) {
		fmt.Println(s)
	}})

	w.AddButton("calc", Option{
		Fn: func(string) {
			input.Text = "good"
			input.Refresh()
			go func() {
				txt := w.Show("question", "# Hello Fyne\n\nThis is **Markdown** support!", "cancel", "??", "ok")
				// w.Show("result", txt)
				w.ShowModel(txt)
			}()
		},
	})
	w.Merge(3, 2)

	// d, ok := w.ShowForm([]*widget.FormItem{
	// 	{Text: "1", Widget: widget.NewEntry()},
	// 	{Text: "2", Widget: widget.NewEntry()},
	// 	{Text: "3", Widget: widget.NewMultiLineEntry()},
	// })
	// if ok {
	// 	fmt.Println(d)
	// }
	t := time.Now()
	w.ShowCalendar(&t)
	fmt.Println(t)
	w.Run()
}

func TestWin_ShowForm(t *testing.T) {
	type fields struct {
		Title string
		w     int
		h     int
		objs  []fyne.CanvasObject
		Win   fyne.Window
	}
	type args struct {
		list []*widget.FormItem
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Win{
				Title: tt.fields.Title,
				w:     tt.fields.w,
				h:     tt.fields.h,
				objs:  tt.fields.objs,
				Win:   tt.fields.Win,
			}
			got, got1 := w.ShowForm(tt.args.list)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Win.ShowForm() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Win.ShowForm() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
