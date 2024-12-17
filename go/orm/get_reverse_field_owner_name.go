// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongmarkdown/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.AnotherDummyData:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Cell:
		switch reverseField.GongstructName {
		// insertion point
		case "Row":
			switch reverseField.Fieldname {
			case "Cells":
				if _row, ok := stage.Row_Cells_reverseMap[inst]; ok {
					res = _row.Name
				}
			}
		}

	case *models.DummyData:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Element:
		switch reverseField.GongstructName {
		// insertion point
		case "Element":
			switch reverseField.Fieldname {
			case "SubElements":
				if _element, ok := stage.Element_SubElements_reverseMap[inst]; ok {
					res = _element.Name
				}
			}
		}

	case *models.MarkdownContent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Row:
		switch reverseField.GongstructName {
		// insertion point
		case "Element":
			switch reverseField.Fieldname {
			case "Rows":
				if _element, ok := stage.Element_Rows_reverseMap[inst]; ok {
					res = _element.Name
				}
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.AnotherDummyData:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Cell:
		switch reverseField.GongstructName {
		// insertion point
		case "Row":
			switch reverseField.Fieldname {
			case "Cells":
				res = stage.Row_Cells_reverseMap[inst]
			}
		}

	case *models.DummyData:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Element:
		switch reverseField.GongstructName {
		// insertion point
		case "Element":
			switch reverseField.Fieldname {
			case "SubElements":
				res = stage.Element_SubElements_reverseMap[inst]
			}
		}

	case *models.MarkdownContent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Row:
		switch reverseField.GongstructName {
		// insertion point
		case "Element":
			switch reverseField.Fieldname {
			case "Rows":
				res = stage.Element_Rows_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
