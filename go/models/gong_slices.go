// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *AnotherDummyData:
		// insertion point per field

	case *Cell:
		// insertion point per field

	case *DummyData:
		// insertion point per field

	case *Element:
		// insertion point per field
		if fieldName == "SubElements" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Element) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Element)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SubElements).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SubElements = _inferedTypeInstance.SubElements[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SubElements =
								append(_inferedTypeInstance.SubElements, any(fieldInstance).(*Element))
						}
					}
				}
			}
		}
		if fieldName == "Rows" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Element) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Element)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Rows).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Rows = _inferedTypeInstance.Rows[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Rows =
								append(_inferedTypeInstance.Rows, any(fieldInstance).(*Row))
						}
					}
				}
			}
		}

	case *MarkdownContent:
		// insertion point per field

	case *Row:
		// insertion point per field
		if fieldName == "Cells" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Row) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Row)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Cells).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Cells = _inferedTypeInstance.Cells[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Cells =
								append(_inferedTypeInstance.Cells, any(fieldInstance).(*Cell))
						}
					}
				}
			}
		}

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct AnotherDummyData
	// insertion point per field

	// Compute reverse map for named struct Cell
	// insertion point per field

	// Compute reverse map for named struct DummyData
	// insertion point per field

	// Compute reverse map for named struct Element
	// insertion point per field
	clear(stage.Element_SubElements_reverseMap)
	stage.Element_SubElements_reverseMap = make(map[*Element]*Element)
	for element := range stage.Elements {
		_ = element
		for _, _element := range element.SubElements {
			stage.Element_SubElements_reverseMap[_element] = element
		}
	}
	clear(stage.Element_Rows_reverseMap)
	stage.Element_Rows_reverseMap = make(map[*Row]*Element)
	for element := range stage.Elements {
		_ = element
		for _, _row := range element.Rows {
			stage.Element_Rows_reverseMap[_row] = element
		}
	}

	// Compute reverse map for named struct MarkdownContent
	// insertion point per field

	// Compute reverse map for named struct Row
	// insertion point per field
	clear(stage.Row_Cells_reverseMap)
	stage.Row_Cells_reverseMap = make(map[*Cell]*Row)
	for row := range stage.Rows {
		_ = row
		for _, _cell := range row.Cells {
			stage.Row_Cells_reverseMap[_cell] = row
		}
	}

}
