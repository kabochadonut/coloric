package coloric

type Color struct {
	Params []Attribute
}

type Attribute int

const escape = "\x1b"

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkFast
	Reverse
	Concealed
	CrossedOut
)

const (
	ResetBold Attribute = iota + 22
	ResetItalic
	ResetUnderline
	ResetBlinking
	_
	ResetReversed
	ResetConcealed
	ResetCrossedOut
)

var mapResetAttributes map[Attribute]Attribute = map[Attribute]Attribute{
	Bold:       ResetBold,
	Faint:      ResetBold,
	Italic:     ResetItalic,
	Underline:  ResetUnderline,
	BlinkSlow:  ResetBlinking,
	BlinkFast:  ResetBlinking,
	Reverse:    ResetReversed,
	Concealed:  ResetConcealed,
	CrossedOut: ResetCrossedOut,
}

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite

	// used internally for 256 and 24-bit coloring
	foreground
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite

	// used internally for 256 and 24-bit coloring
	background
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)
