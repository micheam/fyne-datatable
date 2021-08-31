package datatable

import (
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func New(tmplObj interface{}, data interface{}) *widget.Table {
	var v = reflect.ValueOf(data)
	var headers = tagValues(tmplObj)
	if v.Kind() != reflect.Array {
		panic(fmt.Sprintf("need %v but got %v", reflect.Array, v.Kind()))
	}

	table := widget.NewTable(

		func() (int, int) {
			rows := reflect.ValueOf(data).Len() + 1 // data rows with header
			cols := len(headers)
			return rows, cols
		},

		// callback fn for Create each cell.
		func() fyne.CanvasObject {
			l := widget.NewLabel("placeholder")
			l.Wrapping = fyne.TextTruncate
			return l
		},

		// callback fn for Update each cell.
		// This may trigger on initial rendering process.
		// override result of second param in NewTable()
		func(id widget.TableCellID, c fyne.CanvasObject) {
			label := c.(*widget.Label)
			col, row := id.Col, id.Row
			if row == 0 { // Header Row
				label.Alignment = fyne.TextAlignCenter
				label.TextStyle = fyne.TextStyle{Bold: true}
				label.Text = headers[col]
				return
			}
			// Data row
			acc := reflect.ValueOf(data).Index(row - 1).Interface()
			label.SetText(getFieldValue(acc, col).(string))
		})

	// NOTE: Set width for each columns...
	//
	// Columns for widget.Table is automatically determined from the template object
	// specified in CreateCell (second arg of function NewTable) by default.
	// Here, the size of each column is determined separately from a tmpl of the data.
	for i := range headers {
		text := getFieldValue(tmplObj, i).(string)
		table.SetColumnWidth(i, widget.NewLabel(text).MinSize().Width)
	}

	return table
}
