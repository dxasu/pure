package text

import "github.com/olekukonko/tablewriter/tw"

var symStyles = []tw.Symbols{
	tw.NewSymbolCustom("Dotted").
		WithRow("·").
		WithColumn(":").
		WithTopLeft(".").
		WithTopMid("·").
		WithTopRight(".").
		WithMidLeft(":").
		WithCenter("+").
		WithMidRight(":").
		WithBottomLeft("'").
		WithBottomMid("·").
		WithBottomRight("'"),

	// arrow style
	tw.NewSymbolCustom("Arrow").
		WithRow("→").
		WithColumn("↓").
		WithTopLeft("↗").
		WithTopMid("↑").
		WithTopRight("↖").
		WithMidLeft("→").
		WithCenter("↔").
		WithMidRight("←").
		WithBottomLeft("↘").
		WithBottomMid("↓").
		WithBottomRight("↙"),

	// start style
	tw.NewSymbolCustom("Starry").
		WithRow("★").
		WithColumn("☆").
		WithTopLeft("✧").
		WithTopMid("✯").
		WithTopRight("✧").
		WithMidLeft("✦").
		WithCenter("✶").
		WithMidRight("✦").
		WithBottomLeft("✧").
		WithBottomMid("✯").
		WithBottomRight("✧"),

	tw.NewSymbolCustom("Hearts").
		WithRow("♥").
		WithColumn("❤").
		WithTopLeft("❥").
		WithTopMid("♡").
		WithTopRight("❥").
		WithMidLeft("❣").
		WithCenter("✚").
		WithMidRight("❣").
		WithBottomLeft("❦").
		WithBottomMid("♡").
		WithBottomRight("❦"),

	tw.NewSymbolCustom("Tech").
		WithRow("=").
		WithColumn("||").
		WithTopLeft("/*").
		WithTopMid("##").
		WithTopRight("*/").
		WithMidLeft("//").
		WithCenter("<>").
		WithMidRight("\\").
		WithBottomLeft("\\*").
		WithBottomMid("##").
		WithBottomRight("*/"),

	tw.NewSymbolCustom("Nature").
		WithRow("~").
		WithColumn("|").
		WithTopLeft("🌱").
		WithTopMid("🌿").
		WithTopRight("🌱").
		WithMidLeft("🍃").
		WithCenter("❀").
		WithMidRight("🍃").
		WithBottomLeft("🌻").
		WithBottomMid("🌾").
		WithBottomRight("🌻"),

	tw.NewSymbolCustom("Artistic").
		WithRow("▬").
		WithColumn("▐").
		WithTopLeft("◈").
		WithTopMid("◊").
		WithTopRight("◈").
		WithMidLeft("◀").
		WithCenter("⬔").
		WithMidRight("▶").
		WithBottomLeft("◭").
		WithBottomMid("▣").
		WithBottomRight("◮"),

	tw.NewSymbolCustom("8-Bit").
		WithRow("■").
		WithColumn("█").
		WithTopLeft("╔").
		WithTopMid("▲").
		WithTopRight("╗").
		WithMidLeft("◄").
		WithCenter("♦").
		WithMidRight("►").
		WithBottomLeft("╚").
		WithBottomMid("▼").
		WithBottomRight("╝"),

	tw.NewSymbolCustom("Chaos").
		WithRow("≈").
		WithColumn("§").
		WithTopLeft("⌘").
		WithTopMid("∞").
		WithTopRight("⌥").
		WithMidLeft("⚡").
		WithCenter("☯").
		WithMidRight("♞").
		WithBottomLeft("⌂").
		WithBottomMid("∆").
		WithBottomRight("◊"),

	tw.NewSymbolCustom("Dots").
		WithRow("·").
		WithColumn(" "). // Invisible column lines
		WithTopLeft("·").
		WithTopMid("·").
		WithTopRight("·").
		WithMidLeft(" ").
		WithCenter("·").
		WithMidRight(" ").
		WithBottomLeft("·").
		WithBottomMid("·").
		WithBottomRight("·"),

	tw.NewSymbolCustom("Blocks").
		WithRow("▀").
		WithColumn("█").
		WithTopLeft("▛").
		WithTopMid("▀").
		WithTopRight("▜").
		WithMidLeft("▌").
		WithCenter("█").
		WithMidRight("▐").
		WithBottomLeft("▙").
		WithBottomMid("▄").
		WithBottomRight("▟"),

	tw.NewSymbolCustom("Zen").
		WithRow("~").
		WithColumn(" ").
		WithTopLeft(" ").
		WithTopMid("♨").
		WithTopRight(" ").
		WithMidLeft(" ").
		WithCenter("☯").
		WithMidRight(" ").
		WithBottomLeft(" ").
		WithBottomMid("♨").
		WithBottomRight(" "),
}
