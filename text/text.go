package text

import (
	"io"
	"strings"
	"unsafe"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

// func DataToTable(data string) {
// 	cleaned := strings.ReplaceAll(data, "\u200b", "")
// 	cleaned = strings.ReplaceAll(cleaned, "\u00a0", " ")
// }

type Text struct {
	symbols tw.Symbols
	w       io.Writer
	header  []string
	rows    [][]string
}

func NewText(w io.Writer, header bool, data any) *Text {
	var rows [][]string
	switch data := data.(type) {
	case string:
		d := strings.Split(data, "\n")
		for _, line := range d {
			rows = append(rows, strings.Split(line, "\t"))
		}
	case []string:
		for _, line := range data {
			rows = append(rows, strings.Split(line, "\t"))
		}
	case [][]string:
		rows = data
	default:
	}

	t := &Text{
		symbols: symStyles[0],
		w:       w,
	}
	if header {
		t.header = rows[0]
		rows = rows[1:]
	}
	t.rows = rows
	return t
}

func (t *Text) GetSymbolsName() []string {
	var symStyleNames []string
	for _, s := range symStyles {
		symStyleNames = append(symStyleNames, s.Name())
	}
	return symStyleNames
}

func (t *Text) SetSymbolsByName(symbolName string) {
	for _, s := range symStyles {
		if s.Name() == symbolName {
			t.symbols = s
			return
		}
	}
	t.symbols = symStyles[0]
}

type SymbolCustom struct {
	Name        string
	Center      string
	Row         string
	Column      string
	TopLeft     string
	TopMid      string
	TopRight    string
	MidLeft     string
	MidRight    string
	BottomLeft  string
	BottomMid   string
	BottomRight string
	HeaderLeft  string
	HeaderMid   string
	HeaderRight string
}

func (t *Text) SetSymbols(s *SymbolCustom) {
	t.symbols = (*tw.SymbolCustom)(unsafe.Pointer(s))
}

func (t *Text) Flush() {
	colorCfg := renderer.ColorizedConfig{
		Header: renderer.Tint{
			FG: renderer.Colors{color.FgGreen, color.Bold}, // Green bold headers
			BG: renderer.Colors{color.BgHiWhite},
		},
		Column: renderer.Tint{
			FG: renderer.Colors{color.FgCyan}, // Default cyan for rows
			Columns: []renderer.Tint{
				{FG: renderer.Colors{color.FgMagenta}}, // Magenta for column 0
				{},                                     // Inherit default (cyan)
				{FG: renderer.Colors{color.FgHiRed}},   // High-intensity red for column 2
			},
		},
		Footer: renderer.Tint{
			FG: renderer.Colors{color.FgYellow, color.Bold}, // Yellow bold footer
			Columns: []renderer.Tint{
				{},                                      // Inherit default
				{FG: renderer.Colors{color.FgHiYellow}}, // High-intensity yellow for column 1
				{},                                      // Inherit default
			},
		},
		// Border: renderer.Tint{FG: renderer.Colors{color.NoColor}}, // White borders
		// Separator: renderer.Tint{FG: renderer.Colors{color.FgWhite}}, // White separators
	}

	table := tablewriter.NewTable(t.w,
		tablewriter.WithRenderer(renderer.NewColorized(colorCfg)),
		tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{Symbols: t.symbols})),
		tablewriter.WithConfig(tablewriter.Config{
			Row: tw.CellConfig{
				Formatting: tw.CellFormatting{
					AutoWrap:  tw.WrapNormal, // Wrap long content
					Alignment: tw.AlignLeft,  // Left-align rows
				},
			},
			Header: tw.CellConfig{
				Formatting: tw.CellFormatting{Alignment: tw.AlignLeft},
			},
		}),
	)

	if len(t.header) > 0 {
		table.Header(t.header)
	}

	for _, line := range t.rows {
		table.Append(line)
	}
	table.Render()
}
