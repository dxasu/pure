package gui

import (
	"image/color"
	"strconv"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dxasu/pure/text"

	"time"

	"fyne.io/fyne/v2/container"
)

type Toast struct {
	toastQueue   []string
	isShowing    bool
	parent       fyne.Window
	size         fyne.Size
	intervalTime time.Duration
	animateTime  time.Duration
	cr           text.ColorRegex
}

var (
	toastInstance *Toast = nil
	toastOnce            = sync.Once{}
)

func NewToast(parentWindow fyne.Window) *Toast {
	toastOnce.Do(func() {
		toastInstance = &Toast{
			toastQueue:   make([]string, 0),
			isShowing:    false,
			size:         fyne.NewSize(300, 50),
			intervalTime: 2000 * time.Millisecond,
			animateTime:  200 * time.Millisecond,
			parent:       parentWindow,
			cr:           text.ColorRegexSharp,
		}
	})

	return toastInstance
}

func (t *Toast) QueueToast(msg string) {
	t.toastQueue = append(t.toastQueue, msg)
	if !t.isShowing {
		t.showNextToast()
	}
}

func (t *Toast) showNextToast() {
	if len(t.toastQueue) == 0 {
		t.isShowing = false
		return
	}

	msg := t.toastQueue[0]
	t.toastQueue = t.toastQueue[1:]
	t.isShowing = true

	t.showAnimatedToast(msg)
	time.AfterFunc(t.intervalTime, func() {
		t.showNextToast()
	})
}

func (t *Toast) showAnimatedToast(message string) {
	// 解析ANSI彩色文本
	// textParts := parseANSI(message)
	// content := container.NewHBox(textParts...)

	content := t.TextContainer(message)
	// 创建背景容器
	bg := canvas.NewRectangle(theme.Color(theme.ColorNameBackground))
	toast := container.NewStack(bg, content)
	toast.Resize(t.size)

	// 初始位置（屏幕底部外侧）
	startPos := fyne.NewPos(
		(t.parent.Canvas().Size().Width-300)/2,
		t.parent.Canvas().Size().Height,
	)
	targetPos := fyne.NewPos(
		(t.parent.Canvas().Size().Width-300)/2,
		t.parent.Canvas().Size().Height-100,
	)

	// 创建PopUp并设置初始位置
	popUp := widget.NewPopUp(toast, t.parent.Canvas())
	popUp.Move(startPos)
	popUp.Show()

	// 淡入动画
	fadeIn := canvas.NewPositionAnimation(
		startPos, targetPos,
		t.animateTime, // 动画时长
		func(p fyne.Position) { popUp.Move(p) },
	)
	fadeIn.Curve = fyne.AnimationEaseOut // 缓出效果更自然
	fadeIn.Start()

	// 淡出
	time.AfterFunc(t.intervalTime*7/8, func() {
		fadeOut := canvas.NewPositionAnimation(
			targetPos, startPos,
			t.animateTime,
			func(p fyne.Position) {
				popUp.Move(p)
				if p.Y >= t.parent.Canvas().Size().Height {
					popUp.Hide() // 完全移出后隐藏
				}
			},
		)
		fadeOut.Curve = fyne.AnimationEaseIn
		fadeOut.Start()
	})
}

func (t *Toast) TextContainer(content string) *fyne.Container {
	var parts []fyne.CanvasObject
	list := t.cr.RegSubSplit(content)
	for i := 1; i < len(list); i++ {
		if len(list[i]) == 4 {
			r, _ := strconv.ParseUint(list[i][0], 16, 8)
			g, _ := strconv.ParseUint(list[i][1], 16, 8)
			b, _ := strconv.ParseUint(list[i][2], 16, 8)
			text := canvas.NewText(list[i][3], color.RGBA{uint8(r), uint8(g), uint8(b), 255})
			text.TextSize = 14
			text.TextStyle = fyne.TextStyle{
				TabWidth:  0,
				Monospace: false,
			}
			parts = append(parts, text)
		} else {
			text := &canvas.Text{
				Color:    color.RGBA{100, 100, 100, 255},
				Text:     list[i][0],
				TextSize: 14,
				TextStyle: fyne.TextStyle{
					TabWidth:  0,
					Monospace: false,
				},
			}

			text.TextSize = 14
			parts = append(parts, text)
		}
	}
	return container.New(&zeroSpaceHBoxLayout{}, parts...)
}
