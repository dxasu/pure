package gui

import (
	"bytes"
	"image"
	"os"
)

var (
	iconCache = make(map[string]*image.Image)
)

// 加载图片资源
func loadIcon(path string) (*image.Image, error) {
	if img, ok := iconCache[path]; ok {
		return img, nil
	}
	// 实际加载逻辑...
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	iconCache[path] = &img
	return &img, nil
}
