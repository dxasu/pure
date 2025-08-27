package gui

import (
	"fmt"
	"image/color"
	"regexp"
	"strconv"
	"strings"
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

	content := TextContainer(message)
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

func parseANSI(text string) []fyne.CanvasObject {
	var parts []fyne.CanvasObject
	segments := strings.Split(text, "\x1b[") // 分割ANSI序列

	currentColor := theme.Color(theme.ColorNameBackground)
	for _, seg := range segments {
		if len(seg) < 2 {
			continue
		}

		// 解析颜色代码（简化版，支持30-37前景色）
		if seg[0] == '3' && seg[1] >= '0' && seg[1] <= '7' {
			currentColor = ansiToColor(seg[1])
			seg = seg[2:] // 移除颜色代码
		}

		if len(seg) > 0 {
			text := canvas.NewText(seg, currentColor)
			text.TextSize = 14
			parts = append(parts, text)
		}
	}
	return parts
}

func ColorFormat(f uint32, text string) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", uint8(f>>16), uint8(f>>8), uint8(f), text)
}

func TextContainer(content string) *fyne.Container {
	var parts []fyne.CanvasObject
	// 正则匹配 \033[38;2;R;G;Bm...\033[0m 格式
	re := regexp.MustCompile(`\033\[38;2;(\d{1,3});(\d{1,3});(\d{1,3})m([^\033]+)\033\[0m`)
	list := text.RegSplit(re, content)
	for i := 1; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		if list[i][0] == '\033' {
			matches := re.FindStringSubmatch(list[i])
			r, _ := strconv.ParseUint(matches[1], 10, 8)
			g, _ := strconv.ParseUint(matches[2], 10, 8)
			b, _ := strconv.ParseUint(matches[3], 10, 8)
			text := canvas.NewText(matches[4], color.RGBA{uint8(r), uint8(g), uint8(b), 255})
			text.TextSize = 14
			text.TextStyle = fyne.TextStyle{
				TabWidth:  0,
				Monospace: true,
			}
			parts = append(parts, text)
		} else {
			text := &canvas.Text{
				Color:    color.RGBA{100, 100, 100, 255},
				Text:     list[i],
				TextSize: 14,
				TextStyle: fyne.TextStyle{
					TabWidth:  0,
					Monospace: true,
				},
			}
			text.TextSize = 14
			parts = append(parts, text)
		}
	}
	c := container.NewHBox(parts...)
	return c
}

func extractRGB(text string) color.Color {
	// 假设输入格式严格为 \033[38;2;R;G;Bm...\033[0m
	start := len("\033[38;2;")
	end := strings.Index(text[start:], "m")
	parts := strings.Split(text[start:start+end], ";")
	if len(parts) != 3 {
		return color.RGBA{} // 返回默认颜色
	}
	r, _ := strconv.ParseUint(parts[0], 10, 8)
	g, _ := strconv.ParseUint(parts[1], 10, 8)
	b, _ := strconv.ParseUint(parts[2], 10, 8)
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func ansiToColor(code byte) color.Color {
	switch code {
	case '1':
		return color.RGBA{255, 0, 0, 255} // 红
	case '2':
		return color.RGBA{0, 255, 0, 255} // 绿
	case '3':
		return color.RGBA{255, 255, 0, 255} // 黄
	// 其他颜色映射...
	default:
		return theme.Color(theme.ColorNameBackground) // theme.PrimaryColor()
	}
}
