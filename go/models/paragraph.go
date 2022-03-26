package models

type Element struct {
	Name    string
	Content string

	SubElements []*Element
}
