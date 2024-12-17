// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AnotherDummyData:
		ok = stage.IsStagedAnotherDummyData(target)

	case *Cell:
		ok = stage.IsStagedCell(target)

	case *DummyData:
		ok = stage.IsStagedDummyData(target)

	case *Element:
		ok = stage.IsStagedElement(target)

	case *MarkdownContent:
		ok = stage.IsStagedMarkdownContent(target)

	case *Row:
		ok = stage.IsStagedRow(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedAnotherDummyData(anotherdummydata *AnotherDummyData) (ok bool) {

	_, ok = stage.AnotherDummyDatas[anotherdummydata]

	return
}

func (stage *StageStruct) IsStagedCell(cell *Cell) (ok bool) {

	_, ok = stage.Cells[cell]

	return
}

func (stage *StageStruct) IsStagedDummyData(dummydata *DummyData) (ok bool) {

	_, ok = stage.DummyDatas[dummydata]

	return
}

func (stage *StageStruct) IsStagedElement(element *Element) (ok bool) {

	_, ok = stage.Elements[element]

	return
}

func (stage *StageStruct) IsStagedMarkdownContent(markdowncontent *MarkdownContent) (ok bool) {

	_, ok = stage.MarkdownContents[markdowncontent]

	return
}

func (stage *StageStruct) IsStagedRow(row *Row) (ok bool) {

	_, ok = stage.Rows[row]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *AnotherDummyData:
		stage.StageBranchAnotherDummyData(target)

	case *Cell:
		stage.StageBranchCell(target)

	case *DummyData:
		stage.StageBranchDummyData(target)

	case *Element:
		stage.StageBranchElement(target)

	case *MarkdownContent:
		stage.StageBranchMarkdownContent(target)

	case *Row:
		stage.StageBranchRow(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchAnotherDummyData(anotherdummydata *AnotherDummyData) {

	// check if instance is already staged
	if IsStaged(stage, anotherdummydata) {
		return
	}

	anotherdummydata.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCell(cell *Cell) {

	// check if instance is already staged
	if IsStaged(stage, cell) {
		return
	}

	cell.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDummyData(dummydata *DummyData) {

	// check if instance is already staged
	if IsStaged(stage, dummydata) {
		return
	}

	dummydata.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dummydata.DummyPointerToGongStruct != nil {
		StageBranch(stage, dummydata.DummyPointerToGongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchElement(element *Element) {

	// check if instance is already staged
	if IsStaged(stage, element) {
		return
	}

	element.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _element := range element.SubElements {
		StageBranch(stage, _element)
	}
	for _, _row := range element.Rows {
		StageBranch(stage, _row)
	}

}

func (stage *StageStruct) StageBranchMarkdownContent(markdowncontent *MarkdownContent) {

	// check if instance is already staged
	if IsStaged(stage, markdowncontent) {
		return
	}

	markdowncontent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if markdowncontent.Root != nil {
		StageBranch(stage, markdowncontent.Root)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRow(row *Row) {

	// check if instance is already staged
	if IsStaged(stage, row) {
		return
	}

	row.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _cell := range row.Cells {
		StageBranch(stage, _cell)
	}

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AnotherDummyData:
		toT := CopyBranchAnotherDummyData(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Cell:
		toT := CopyBranchCell(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DummyData:
		toT := CopyBranchDummyData(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Element:
		toT := CopyBranchElement(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MarkdownContent:
		toT := CopyBranchMarkdownContent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Row:
		toT := CopyBranchRow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAnotherDummyData(mapOrigCopy map[any]any, anotherdummydataFrom *AnotherDummyData) (anotherdummydataTo *AnotherDummyData) {

	// anotherdummydataFrom has already been copied
	if _anotherdummydataTo, ok := mapOrigCopy[anotherdummydataFrom]; ok {
		anotherdummydataTo = _anotherdummydataTo.(*AnotherDummyData)
		return
	}

	anotherdummydataTo = new(AnotherDummyData)
	mapOrigCopy[anotherdummydataFrom] = anotherdummydataTo
	anotherdummydataFrom.CopyBasicFields(anotherdummydataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCell(mapOrigCopy map[any]any, cellFrom *Cell) (cellTo *Cell) {

	// cellFrom has already been copied
	if _cellTo, ok := mapOrigCopy[cellFrom]; ok {
		cellTo = _cellTo.(*Cell)
		return
	}

	cellTo = new(Cell)
	mapOrigCopy[cellFrom] = cellTo
	cellFrom.CopyBasicFields(cellTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDummyData(mapOrigCopy map[any]any, dummydataFrom *DummyData) (dummydataTo *DummyData) {

	// dummydataFrom has already been copied
	if _dummydataTo, ok := mapOrigCopy[dummydataFrom]; ok {
		dummydataTo = _dummydataTo.(*DummyData)
		return
	}

	dummydataTo = new(DummyData)
	mapOrigCopy[dummydataFrom] = dummydataTo
	dummydataFrom.CopyBasicFields(dummydataTo)

	//insertion point for the staging of instances referenced by pointers
	if dummydataFrom.DummyPointerToGongStruct != nil {
		dummydataTo.DummyPointerToGongStruct = CopyBranchAnotherDummyData(mapOrigCopy, dummydataFrom.DummyPointerToGongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchElement(mapOrigCopy map[any]any, elementFrom *Element) (elementTo *Element) {

	// elementFrom has already been copied
	if _elementTo, ok := mapOrigCopy[elementFrom]; ok {
		elementTo = _elementTo.(*Element)
		return
	}

	elementTo = new(Element)
	mapOrigCopy[elementFrom] = elementTo
	elementFrom.CopyBasicFields(elementTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _element := range elementFrom.SubElements {
		elementTo.SubElements = append(elementTo.SubElements, CopyBranchElement(mapOrigCopy, _element))
	}
	for _, _row := range elementFrom.Rows {
		elementTo.Rows = append(elementTo.Rows, CopyBranchRow(mapOrigCopy, _row))
	}

	return
}

func CopyBranchMarkdownContent(mapOrigCopy map[any]any, markdowncontentFrom *MarkdownContent) (markdowncontentTo *MarkdownContent) {

	// markdowncontentFrom has already been copied
	if _markdowncontentTo, ok := mapOrigCopy[markdowncontentFrom]; ok {
		markdowncontentTo = _markdowncontentTo.(*MarkdownContent)
		return
	}

	markdowncontentTo = new(MarkdownContent)
	mapOrigCopy[markdowncontentFrom] = markdowncontentTo
	markdowncontentFrom.CopyBasicFields(markdowncontentTo)

	//insertion point for the staging of instances referenced by pointers
	if markdowncontentFrom.Root != nil {
		markdowncontentTo.Root = CopyBranchElement(mapOrigCopy, markdowncontentFrom.Root)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRow(mapOrigCopy map[any]any, rowFrom *Row) (rowTo *Row) {

	// rowFrom has already been copied
	if _rowTo, ok := mapOrigCopy[rowFrom]; ok {
		rowTo = _rowTo.(*Row)
		return
	}

	rowTo = new(Row)
	mapOrigCopy[rowFrom] = rowTo
	rowFrom.CopyBasicFields(rowTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _cell := range rowFrom.Cells {
		rowTo.Cells = append(rowTo.Cells, CopyBranchCell(mapOrigCopy, _cell))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *AnotherDummyData:
		stage.UnstageBranchAnotherDummyData(target)

	case *Cell:
		stage.UnstageBranchCell(target)

	case *DummyData:
		stage.UnstageBranchDummyData(target)

	case *Element:
		stage.UnstageBranchElement(target)

	case *MarkdownContent:
		stage.UnstageBranchMarkdownContent(target)

	case *Row:
		stage.UnstageBranchRow(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchAnotherDummyData(anotherdummydata *AnotherDummyData) {

	// check if instance is already staged
	if !IsStaged(stage, anotherdummydata) {
		return
	}

	anotherdummydata.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCell(cell *Cell) {

	// check if instance is already staged
	if !IsStaged(stage, cell) {
		return
	}

	cell.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDummyData(dummydata *DummyData) {

	// check if instance is already staged
	if !IsStaged(stage, dummydata) {
		return
	}

	dummydata.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dummydata.DummyPointerToGongStruct != nil {
		UnstageBranch(stage, dummydata.DummyPointerToGongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchElement(element *Element) {

	// check if instance is already staged
	if !IsStaged(stage, element) {
		return
	}

	element.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _element := range element.SubElements {
		UnstageBranch(stage, _element)
	}
	for _, _row := range element.Rows {
		UnstageBranch(stage, _row)
	}

}

func (stage *StageStruct) UnstageBranchMarkdownContent(markdowncontent *MarkdownContent) {

	// check if instance is already staged
	if !IsStaged(stage, markdowncontent) {
		return
	}

	markdowncontent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if markdowncontent.Root != nil {
		UnstageBranch(stage, markdowncontent.Root)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRow(row *Row) {

	// check if instance is already staged
	if !IsStaged(stage, row) {
		return
	}

	row.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _cell := range row.Cells {
		UnstageBranch(stage, _cell)
	}

}
