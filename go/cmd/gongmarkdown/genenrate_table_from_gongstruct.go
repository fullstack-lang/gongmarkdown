package main

import (
	"github.com/fullstack-lang/gongmarkdown/go/models"
	gongmarkdown_models "github.com/fullstack-lang/gongmarkdown/go/models"
)

func GenerateTableFromDummyDatas(gongStructInstances []models.GetFieldsInterface) (element *gongmarkdown_models.Element) {

	if len(gongStructInstances) == 0 {
		return
	}

	element = (&gongmarkdown_models.Element{Name: "TableOfDummnies "}).Stage()
	element.Type = gongmarkdown_models.TABLE

	// fill up title cells
	titleRow := (&gongmarkdown_models.Row{Name: "Title Row "}).Stage()
	element.Rows = append(element.Rows, titleRow)

	fields := gongStructInstances[0].GetFields()
	for _, field := range fields {
		titleRow.Cells = append(titleRow.Cells, (&gongmarkdown_models.Cell{Name: field}).Stage())
	}

	for _, gongStructInstance := range gongStructInstances {
		// fill up title cells
		row := (&gongmarkdown_models.Row{Name: "Date Row of " + gongStructInstance.GetFieldStringValue("Name")}).Stage()
		element.Rows = append(element.Rows, row)

		for _, field := range fields {
			cell := (&gongmarkdown_models.Cell{Name: gongStructInstance.GetFieldStringValue(field)}).Stage()
			row.Cells = append(row.Cells, cell)
		}
	}

	return
}

func SimpleInterfaceCall(gongStructInstances models.GetFieldsInterface) {

}
