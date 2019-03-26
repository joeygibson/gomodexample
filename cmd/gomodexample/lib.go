package gomodexample

import "github.com/fatih/color"

// PrintInRed prints the given message in red
func PrintInRed(msg string) {
	color.Red(msg)
}

// PrintInGreen prints the given message in green
func PrintInGreen(msg string) {
	color.Green(msg)
}

// PrintInYellow prints the given message in yellow
func PrintInYellow(msg string) {
	color.Yellow(msg)
}

// PrintInBlue prints the given message in blue
func PrintInBlue(msg string) {
	color.Blue(msg)
}

// PrintInCyan prints the given message in cyan
func PrintInCyan(msg string) {
	color.Cyan(msg)
}
