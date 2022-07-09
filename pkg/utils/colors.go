package utils

import (
	"github.com/fatih/color"
)

var colors = [][]*color.Color{
	{color.New(color.FgHiMagenta).Add(color.Bold), color.New(color.FgMagenta)},
	{color.New(color.FgHiYellow).Add(color.Bold), color.New(color.FgYellow)},
	{color.New(color.FgHiGreen).Add(color.Bold), color.New(color.FgGreen)},
	{color.New(color.FgHiBlue).Add(color.Bold), color.New(color.FgBlue)},
	{color.New(color.FgHiCyan).Add(color.Bold), color.New(color.FgCyan)},
	{color.New(color.FgHiRed).Add(color.Bold), color.New(color.FgRed)},
}

// Get two color functions
func GetColorsFn(id int) (func(...interface{}) string, func(...interface{}) string) {
	if id >= len(colors) {
		id %= len(colors)
	}

	return colors[id][0].SprintFunc(), colors[id][1].SprintFunc()
}

var colorsSimple = []*color.Color{
	color.New(color.FgMagenta).Add(color.Bold),
	color.New(color.FgYellow).Add(color.Bold),
	color.New(color.FgGreen).Add(color.Bold),
	color.New(color.FgBlue).Add(color.Bold),
	color.New(color.FgCyan).Add(color.Bold),
	color.New(color.FgRed).Add(color.Bold),
	color.New(color.FgHiMagenta).Add(color.Bold),
	color.New(color.FgHiYellow).Add(color.Bold),
	color.New(color.FgHiGreen).Add(color.Bold),
	color.New(color.FgHiBlue).Add(color.Bold),
	color.New(color.FgHiCyan).Add(color.Bold),
	color.New(color.FgHiRed).Add(color.Bold),
}

// Get a color function
func GetColorFn(id int) func(...interface{}) string {
	if id >= len(colorsSimple) {
		id %= len(colorsSimple)
	}
	return colorsSimple[id].SprintFunc()
}
