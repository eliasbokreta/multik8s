package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

var tableColors = []tablewriter.Colors{
	{tablewriter.Bold, tablewriter.FgHiGreenColor},
	{tablewriter.Bold, tablewriter.FgHiBlueColor},
	{tablewriter.Bold, tablewriter.FgHiMagentaColor},
	{tablewriter.Bold, tablewriter.FgHiYellowColor},
	{tablewriter.Bold, tablewriter.FgHiRedColor},
	{tablewriter.Bold, tablewriter.FgHiCyanColor},
	{tablewriter.Bold, tablewriter.FgGreenColor},
	{tablewriter.Bold, tablewriter.FgBlueColor},
	{tablewriter.Bold, tablewriter.FgMagentaColor},
	{tablewriter.Bold, tablewriter.FgYellowColor},
	{tablewriter.Bold, tablewriter.FgRedColor},
	{tablewriter.Bold, tablewriter.FgCyanColor},
}

// Get a tablewriter color depending of a given ID
func GetTableRowColor(id int) tablewriter.Colors {
	if id >= len(tableColors) {
		id %= len(tableColors)
	}

	return tableColors[id]
}

// Generate a writer for tablewriter
func GetTableWriter(header []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(header)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t\t")
	table.SetNoWhiteSpace(true)

	return table
}
