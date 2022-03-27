package main

import (
	"time"

	"github.com/fullstack-lang/gongmarkdown/go/models"
)

func init() {
	var __Dummy_time_variable time.Time
	_ = __Dummy_time_variable
	InjectionGateway["stage"] = stageInjection
}

// stageInjection will stage objects of database "stage"
func stageInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Cell
	__Cell__000000_C11 := (&models.Cell{Name: "C11"}).Stage()
	__Cell__000001_C12 := (&models.Cell{Name: "C12"}).Stage()
	__Cell__000002_C13 := (&models.Cell{Name: "C13"}).Stage()
	__Cell__000003_C21 := (&models.Cell{Name: "C21"}).Stage()
	__Cell__000004_C22 := (&models.Cell{Name: "C22"}).Stage()
	__Cell__000005_C23 := (&models.Cell{Name: "C23"}).Stage()
	__Cell__000006_C31 := (&models.Cell{Name: "C31"}).Stage()
	__Cell__000007_C32 := (&models.Cell{Name: "C32"}).Stage()
	__Cell__000008_C33 := (&models.Cell{Name: "C33"}).Stage()

	// Declarations of staged instances of Element
	__Element__000000_Chapter_1 := (&models.Element{Name: "Chapter 1"}).Stage()
	__Element__000001_Chapter_1_1_ := (&models.Element{Name: "Chapter 1 1 "}).Stage()
	__Element__000002_Chapter_2 := (&models.Element{Name: "Chapter 2"}).Stage()
	__Element__000003_Content_of_chapter_1_1_ := (&models.Element{Name: "Content of chapter 1 1 "}).Stage()
	__Element__000004_List_in_Chapter_1 := (&models.Element{Name: "List in Chapter 1"}).Stage()
	__Element__000005_Root_Element := (&models.Element{Name: "Root Element"}).Stage()
	__Element__000006_Table_in_Chapter_2 := (&models.Element{Name: "Table in Chapter 2"}).Stage()

	// Declarations of staged instances of MarkdownContent
	__MarkdownContent__000000_Singloton := (&models.MarkdownContent{Name: "Singloton"}).Stage()

	// Declarations of staged instances of Row
	__Row__000000_R1 := (&models.Row{Name: "R1"}).Stage()
	__Row__000001_R2 := (&models.Row{Name: "R2"}).Stage()
	__Row__000002_R3 := (&models.Row{Name: "R3"}).Stage()

	// Setup of values

	// Cell C11 values setup
	__Cell__000000_C11.Name = `C11`

	// Cell C12 values setup
	__Cell__000001_C12.Name = `C12`

	// Cell C13 values setup
	__Cell__000002_C13.Name = `C13`

	// Cell C21 values setup
	__Cell__000003_C21.Name = `C21`

	// Cell C22 values setup
	__Cell__000004_C22.Name = `C22`

	// Cell C23 values setup
	__Cell__000005_C23.Name = `C23`

	// Cell C31 values setup
	__Cell__000006_C31.Name = `C31`

	// Cell C32 values setup
	__Cell__000007_C32.Name = `C32`

	// Cell C33 values setup
	__Cell__000008_C33.Name = `C33`

	// Element Chapter 1 values setup
	__Element__000000_Chapter_1.Name = `Chapter 1`
	__Element__000000_Chapter_1.Content = `Title of chapter 1`
	__Element__000000_Chapter_1.Type = models.TITLE

	// Element Chapter 1 1  values setup
	__Element__000001_Chapter_1_1_.Name = `Chapter 1 1 `
	__Element__000001_Chapter_1_1_.Content = `Title of chapter 1.1`
	__Element__000001_Chapter_1_1_.Type = models.TITLE

	// Element Chapter 2 values setup
	__Element__000002_Chapter_2.Name = `Chapter 2`
	__Element__000002_Chapter_2.Content = `Title of chapter 2`
	__Element__000002_Chapter_2.Type = models.TITLE

	// Element Content of chapter 1 1  values setup
	__Element__000003_Content_of_chapter_1_1_.Name = `Content of chapter 1 1 `
	__Element__000003_Content_of_chapter_1_1_.Content = `Content of chapter 1 1`
	__Element__000003_Content_of_chapter_1_1_.Type = models.PARAGRAPH

	// Element List in Chapter 1 values setup
	__Element__000004_List_in_Chapter_1.Name = `List in Chapter 1`
	__Element__000004_List_in_Chapter_1.Content = `- Item 1
- Item 2
- Item 3
  - Item 3.1
  - Item 3.2


`
	__Element__000004_List_in_Chapter_1.Type = models.PARAGRAPH

	// Element Root Element values setup
	__Element__000005_Root_Element.Name = `Root Element`
	__Element__000005_Root_Element.Content = ``
	__Element__000005_Root_Element.Type = models.PARAGRAPH

	// Element Table in Chapter 2 values setup
	__Element__000006_Table_in_Chapter_2.Name = `Table in Chapter 2`
	__Element__000006_Table_in_Chapter_2.Content = `Legend of table`
	__Element__000006_Table_in_Chapter_2.Type = models.TABLE

	// MarkdownContent Singloton values setup
	__MarkdownContent__000000_Singloton.Name = `Singloton`
	__MarkdownContent__000000_Singloton.Content = `

## Title of chapter 1

- Item 1
- Item 2
- Item 3
  - Item 3.1
  - Item 3.2




### Title of chapter 1.1

Content of chapter 1 1

## Title of chapter 2

Legend of table

`

	// Row R1 values setup
	__Row__000000_R1.Name = `R1`

	// Row R2 values setup
	__Row__000001_R2.Name = `R2`

	// Row R3 values setup
	__Row__000002_R3.Name = `R3`

	// Setup of pointers
	__Element__000000_Chapter_1.SubElements = append(__Element__000000_Chapter_1.SubElements, __Element__000004_List_in_Chapter_1)
	__Element__000000_Chapter_1.SubElements = append(__Element__000000_Chapter_1.SubElements, __Element__000001_Chapter_1_1_)
	__Element__000001_Chapter_1_1_.SubElements = append(__Element__000001_Chapter_1_1_.SubElements, __Element__000003_Content_of_chapter_1_1_)
	__Element__000002_Chapter_2.SubElements = append(__Element__000002_Chapter_2.SubElements, __Element__000006_Table_in_Chapter_2)
	__Element__000005_Root_Element.SubElements = append(__Element__000005_Root_Element.SubElements, __Element__000000_Chapter_1)
	__Element__000005_Root_Element.SubElements = append(__Element__000005_Root_Element.SubElements, __Element__000002_Chapter_2)
	__Element__000006_Table_in_Chapter_2.Rows = append(__Element__000006_Table_in_Chapter_2.Rows, __Row__000000_R1)
	__Element__000006_Table_in_Chapter_2.Rows = append(__Element__000006_Table_in_Chapter_2.Rows, __Row__000001_R2)
	__Element__000006_Table_in_Chapter_2.Rows = append(__Element__000006_Table_in_Chapter_2.Rows, __Row__000002_R3)
	__MarkdownContent__000000_Singloton.Root = __Element__000005_Root_Element
	__Row__000000_R1.Cells = append(__Row__000000_R1.Cells, __Cell__000000_C11)
	__Row__000000_R1.Cells = append(__Row__000000_R1.Cells, __Cell__000001_C12)
	__Row__000000_R1.Cells = append(__Row__000000_R1.Cells, __Cell__000002_C13)
	__Row__000001_R2.Cells = append(__Row__000001_R2.Cells, __Cell__000003_C21)
	__Row__000001_R2.Cells = append(__Row__000001_R2.Cells, __Cell__000004_C22)
	__Row__000001_R2.Cells = append(__Row__000001_R2.Cells, __Cell__000005_C23)
	__Row__000002_R3.Cells = append(__Row__000002_R3.Cells, __Cell__000006_C31)
	__Row__000002_R3.Cells = append(__Row__000002_R3.Cells, __Cell__000007_C32)
	__Row__000002_R3.Cells = append(__Row__000002_R3.Cells, __Cell__000008_C33)
}


