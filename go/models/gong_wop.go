// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type AnotherDummyData_WOP struct {
	// insertion point
	Name string
}

func (from *AnotherDummyData) CopyBasicFields(to *AnotherDummyData) {
	// insertion point
	to.Name = from.Name
}

type Cell_WOP struct {
	// insertion point
	Name string
}

func (from *Cell) CopyBasicFields(to *Cell) {
	// insertion point
	to.Name = from.Name
}

type DummyData_WOP struct {
	// insertion point
	Name string
	DummyString string
	DummyInt int
	DummyFloat float64
	DummyBool bool
	DummyEnumString ElementType
	DummyEnumInt DummnyTypeInt
	DummyTime time.Time
	DummyDuration time.Duration
}

func (from *DummyData) CopyBasicFields(to *DummyData) {
	// insertion point
	to.Name = from.Name
	to.DummyString = from.DummyString
	to.DummyInt = from.DummyInt
	to.DummyFloat = from.DummyFloat
	to.DummyBool = from.DummyBool
	to.DummyEnumString = from.DummyEnumString
	to.DummyEnumInt = from.DummyEnumInt
	to.DummyTime = from.DummyTime
	to.DummyDuration = from.DummyDuration
}

type Element_WOP struct {
	// insertion point
	Name string
	Content string
	Type ElementType
}

func (from *Element) CopyBasicFields(to *Element) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
	to.Type = from.Type
}

type MarkdownContent_WOP struct {
	// insertion point
	Name string
	Content string
}

func (from *MarkdownContent) CopyBasicFields(to *MarkdownContent) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

type Row_WOP struct {
	// insertion point
	Name string
}

func (from *Row) CopyBasicFields(to *Row) {
	// insertion point
	to.Name = from.Name
}

