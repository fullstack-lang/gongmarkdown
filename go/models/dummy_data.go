package models

import (
	"fmt"
	"time"
)

type DummyData struct {
	Name        string
	DummyString string
	DummyInt    int
	DummyFloat  float64
	DummyBool   bool

	DummyEnumString ElementType

	DummyEnumInt DummnyTypeInt

	DummyTime                time.Time
	DummyDuration            time.Duration
	DummyPointerToGongStruct *AnotherDummyData
}

type DummnyTypeInt int

// values for EnumType
const (
	ONE DummnyTypeInt = 1
	TWO DummnyTypeInt = 2
)

func GenerateTableOfDummnies(stage *StageStruct) (element *Element) {
	element = (&Element{Name: "TableOfDummnies "}).Stage(stage)
	element.Type = TABLE

	// fill up title cells
	titleRow := (&Row{Name: "Title Row "}).Stage(stage)
	element.Rows = append(element.Rows, titleRow)

	{
		cell := (&Cell{Name: "Name "}).Stage(stage)
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyString "}).Stage(stage)
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyFloat "}).Stage(stage)
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyInt "}).Stage(stage)
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyBool "}).Stage(stage)
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyTime "}).Stage(stage)
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyDuration "}).Stage(stage)
		titleRow.Cells = append(titleRow.Cells, cell)
	}
	{
		cell := (&Cell{Name: "DummyPointerToGongStruct "}).Stage(stage)
		titleRow.Cells = append(titleRow.Cells, cell)
	}

	for dummy := range stage.DummyDatas {
		// fill up title cells
		row := (&Row{Name: "Date Row of " + dummy.Name}).Stage(stage)
		element.Rows = append(element.Rows, row)

		{
			cell := (&Cell{Name: dummy.Name}).Stage(stage)
			row.Cells = append(row.Cells, cell)
		}
		{
			cell := (&Cell{Name: dummy.DummyString}).Stage(stage)
			row.Cells = append(row.Cells, cell)
		}
		{
			cell := (&Cell{Name: fmt.Sprintf("%f", dummy.DummyFloat)}).Stage(stage)
			row.Cells = append(row.Cells, cell)
		}
		{
			cell := (&Cell{Name: fmt.Sprintf("%d", dummy.DummyInt)}).Stage(stage)
			row.Cells = append(row.Cells, cell)
		}
		{
			cell := (&Cell{Name: fmt.Sprintf("%t", dummy.DummyBool)}).Stage(stage)
			row.Cells = append(row.Cells, cell)
		}
		{
			cell := (&Cell{Name: dummy.DummyTime.String()}).Stage(stage)
			row.Cells = append(row.Cells, cell)
		}
		{
			cell := (&Cell{Name: dummy.DummyDuration.String()}).Stage(stage)
			row.Cells = append(row.Cells, cell)
		}
		{
			if dummy.DummyPointerToGongStruct != nil {
				cell := (&Cell{Name: dummy.DummyPointerToGongStruct.Name}).Stage(stage)
				row.Cells = append(row.Cells, cell)
			}
		}
	}
	return
}
