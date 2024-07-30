package pkg

import (
	"fmt"

	color "github.com/fatih/color"
)

var (
	RED     = color.New(color.FgRed, color.Bold)
	BLUE    = color.New(color.FgBlue, color.Bold)
	YELLOW  = color.New(color.FgHiYellow, color.Bold)
	GREEN   = color.New(color.FgGreen, color.Bold)
	WHITE   = color.New(color.FgWhite, color.Bold)
	CYAN    = color.New(color.FgCyan, color.Bold)
	MAGENTA = color.New(color.FgMagenta, color.Bold)
)

const (
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
	ColorReset   = "\033[0m"
)

func Print(Text string) string {
	fmt.Print(Text)
	return Text
}

func Int(Text string) int {
	intValue := 0
	fmt.Sscan(Text, &intValue)
	return intValue
}

func PPrint(COLORT *color.Color, Text string, LineUnder bool) string {
	var Line string
	if LineUnder {
		Line = "\n"
	}
	WHITE.Print("[")
	RED.Print("+")
	WHITE.Print("] ")
	COLORT.Print(Text, Line)
	return Text
}

func Input(Output string) string {
	PPrint(WHITE, Output, false)
	fmt.Scanln(&Output)

	return Output
}
func setColor(color string) {
	fmt.Print(color)
}

func resetColor() {
	fmt.Print(ColorReset)
}
