package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// Generate a writer for table
func GetTableWriter(header []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)

	colors := []tablewriter.Colors{}
	colors = append(colors, tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor})

	for i := 1; i < len(header); i++ {
		colors = append(colors, tablewriter.Colors{tablewriter.Normal, tablewriter.Normal})
	}

	table.SetHeader(header)
	table.SetColumnColor(colors...)
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
