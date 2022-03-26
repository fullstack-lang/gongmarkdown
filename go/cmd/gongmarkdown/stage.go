package main

import (
	"time"

	"gongmarkdown/go/models"
)

func init() {
	var __Dummy_time_variable time.Time
	_ = __Dummy_time_variable
	InjectionGateway["stage"] = stageInjection
}

// stageInjection will stage objects of database "stage"
func stageInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Element
	__Element__000000_Chapter_1 := (&models.Element{Name: "Chapter 1"}).Stage()
	__Element__000001_Chapter_1_1_ := (&models.Element{Name: "Chapter 1 1 "}).Stage()
	__Element__000002_Chapter_2 := (&models.Element{Name: "Chapter 2"}).Stage()
	__Element__000003_Content_of_chapter_1_1_ := (&models.Element{Name: "Content of chapter 1 1 "}).Stage()
	__Element__000004_Root_Element := (&models.Element{Name: "Root Element"}).Stage()

	// Declarations of staged instances of MarkdownContent
	__MarkdownContent__000000_Singloton := (&models.MarkdownContent{Name: "Singloton"}).Stage()

	// Setup of values

	// Element Chapter 1 values setup
	__Element__000000_Chapter_1.Name = `Chapter 1`
	__Element__000000_Chapter_1.Content = `Title of chapter 1`
	__Element__000000_Chapter_1.IsTitle = true

	// Element Chapter 1 1  values setup
	__Element__000001_Chapter_1_1_.Name = `Chapter 1 1 `
	__Element__000001_Chapter_1_1_.Content = `Title of chapter 1.1`
	__Element__000001_Chapter_1_1_.IsTitle = true

	// Element Chapter 2 values setup
	__Element__000002_Chapter_2.Name = `Chapter 2`
	__Element__000002_Chapter_2.Content = `Title of chapter 2`
	__Element__000002_Chapter_2.IsTitle = true

	// Element Content of chapter 1 1  values setup
	__Element__000003_Content_of_chapter_1_1_.Name = `Content of chapter 1 1 `
	__Element__000003_Content_of_chapter_1_1_.Content = `Content of chapter 1 1`
	__Element__000003_Content_of_chapter_1_1_.IsTitle = false

	// Element Root Element values setup
	__Element__000004_Root_Element.Name = `Root Element`
	__Element__000004_Root_Element.Content = ``
	__Element__000004_Root_Element.IsTitle = false

	// MarkdownContent Singloton values setup
	__MarkdownContent__000000_Singloton.Name = `Singloton`
	__MarkdownContent__000000_Singloton.Content = ``

	// Setup of pointers
	__Element__000000_Chapter_1.SubElements = append(__Element__000000_Chapter_1.SubElements, __Element__000001_Chapter_1_1_)
	__Element__000000_Chapter_1.SubElements = append(__Element__000000_Chapter_1.SubElements, __Element__000003_Content_of_chapter_1_1_)
	__Element__000004_Root_Element.SubElements = append(__Element__000004_Root_Element.SubElements, __Element__000000_Chapter_1)
	__Element__000004_Root_Element.SubElements = append(__Element__000004_Root_Element.SubElements, __Element__000002_Chapter_2)
	__MarkdownContent__000000_Singloton.Root = __Element__000004_Root_Element
}


