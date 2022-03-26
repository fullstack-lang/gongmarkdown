package models

type Element struct {
	Name    string
	Content string

	IsTitle bool

	SubElements []*Element
}
