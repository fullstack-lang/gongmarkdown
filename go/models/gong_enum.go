// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for DummnyTypeInt
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (dummnytypeint DummnyTypeInt) ToInt() (res int) {

	// migration of former implementation of enum
	switch dummnytypeint {
	// insertion code per enum code
	case ONE:
		res = 0
	case TWO:
		res = 1
	}
	return
}

func (dummnytypeint *DummnyTypeInt) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*dummnytypeint = ONE
		return
	case 1:
		*dummnytypeint = TWO
		return
	default:
		return errUnkownEnum
	}
}

func (dummnytypeint *DummnyTypeInt) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ONE":
		*dummnytypeint = ONE
	case "TWO":
		*dummnytypeint = TWO
	default:
		return errUnkownEnum
	}
	return
}

func (dummnytypeint *DummnyTypeInt) ToCodeString() (res string) {

	switch *dummnytypeint {
	// insertion code per enum code
	case ONE:
		res = "ONE"
	case TWO:
		res = "TWO"
	}
	return
}

func (dummnytypeint DummnyTypeInt) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ONE")
	res = append(res, "TWO")

	return
}

func (dummnytypeint DummnyTypeInt) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)

	return
}

// Utility function for ElementType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (elementtype ElementType) ToString() (res string) {

	// migration of former implementation of enum
	switch elementtype {
	// insertion code per enum code
	case PARAGRAPH:
		res = "Paragraph"
	case TITLE:
		res = "Title"
	case TABLE:
		res = "Table"
	}
	return
}

func (elementtype *ElementType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Paragraph":
		*elementtype = PARAGRAPH
		return
	case "Title":
		*elementtype = TITLE
		return
	case "Table":
		*elementtype = TABLE
		return
	default:
		return errUnkownEnum
	}
}

func (elementtype *ElementType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PARAGRAPH":
		*elementtype = PARAGRAPH
	case "TITLE":
		*elementtype = TITLE
	case "TABLE":
		*elementtype = TABLE
	default:
		return errUnkownEnum
	}
	return
}

func (elementtype *ElementType) ToCodeString() (res string) {

	switch *elementtype {
	// insertion code per enum code
	case PARAGRAPH:
		res = "PARAGRAPH"
	case TITLE:
		res = "TITLE"
	case TABLE:
		res = "TABLE"
	}
	return
}

func (elementtype ElementType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PARAGRAPH")
	res = append(res, "TITLE")
	res = append(res, "TABLE")

	return
}

func (elementtype ElementType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Paragraph")
	res = append(res, "Title")
	res = append(res, "Table")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type PointerToGongstructEnumStringField interface {
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int | DummnyTypeInt
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	*DummnyTypeInt
	FromCodeString(input string) (err error)
}

// Last line of the template
