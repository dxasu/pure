package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct{ fyne.Theme }

// 确保实现了 fyne.Theme 接口
var _ fyne.Theme = (*myTheme)(nil)

// 自定义颜色
func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.NRGBA{R: 0xf0, G: 0xf0, B: 0xff, A: 0xff} // 浅蓝色背景
		}
		return color.NRGBA{R: 0x30, G: 0x30, B: 0x40, A: 0xff} // 深蓝色背景
	}

	if name == theme.ColorNamePrimary {
		return color.NRGBA{R: 0x80, G: 0x80, B: 0xff, A: 0xff} // 紫蓝色主色调
	}

	// 其他颜色使用默认主题
	return theme.DefaultTheme().Color(name, variant)
}

// 使用默认字体
func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

// 使用默认图标
func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// 自定义尺寸
func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNamePadding {
		return 10 // 自定义内边距
	}
	return theme.DefaultTheme().Size(name)
}
