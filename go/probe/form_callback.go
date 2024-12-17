// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongmarkdown/go/models"
	"github.com/fullstack-lang/gongmarkdown/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__AnotherDummyDataFormCallback(
	anotherdummydata *models.AnotherDummyData,
	probe *Probe,
	formGroup *table.FormGroup,
) (anotherdummydataFormCallback *AnotherDummyDataFormCallback) {
	anotherdummydataFormCallback = new(AnotherDummyDataFormCallback)
	anotherdummydataFormCallback.probe = probe
	anotherdummydataFormCallback.anotherdummydata = anotherdummydata
	anotherdummydataFormCallback.formGroup = formGroup

	anotherdummydataFormCallback.CreationMode = (anotherdummydata == nil)

	return
}

type AnotherDummyDataFormCallback struct {
	anotherdummydata *models.AnotherDummyData

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (anotherdummydataFormCallback *AnotherDummyDataFormCallback) OnSave() {

	log.Println("AnotherDummyDataFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	anotherdummydataFormCallback.probe.formStage.Checkout()

	if anotherdummydataFormCallback.anotherdummydata == nil {
		anotherdummydataFormCallback.anotherdummydata = new(models.AnotherDummyData).Stage(anotherdummydataFormCallback.probe.stageOfInterest)
	}
	anotherdummydata_ := anotherdummydataFormCallback.anotherdummydata
	_ = anotherdummydata_

	for _, formDiv := range anotherdummydataFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(anotherdummydata_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if anotherdummydataFormCallback.formGroup.HasSuppressButtonBeenPressed {
		anotherdummydata_.Unstage(anotherdummydataFormCallback.probe.stageOfInterest)
	}

	anotherdummydataFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.AnotherDummyData](
		anotherdummydataFormCallback.probe,
	)
	anotherdummydataFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if anotherdummydataFormCallback.CreationMode || anotherdummydataFormCallback.formGroup.HasSuppressButtonBeenPressed {
		anotherdummydataFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(anotherdummydataFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AnotherDummyDataFormCallback(
			nil,
			anotherdummydataFormCallback.probe,
			newFormGroup,
		)
		anotherdummydata := new(models.AnotherDummyData)
		FillUpForm(anotherdummydata, newFormGroup, anotherdummydataFormCallback.probe)
		anotherdummydataFormCallback.probe.formStage.Commit()
	}

	fillUpTree(anotherdummydataFormCallback.probe)
}
func __gong__New__CellFormCallback(
	cell *models.Cell,
	probe *Probe,
	formGroup *table.FormGroup,
) (cellFormCallback *CellFormCallback) {
	cellFormCallback = new(CellFormCallback)
	cellFormCallback.probe = probe
	cellFormCallback.cell = cell
	cellFormCallback.formGroup = formGroup

	cellFormCallback.CreationMode = (cell == nil)

	return
}

type CellFormCallback struct {
	cell *models.Cell

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (cellFormCallback *CellFormCallback) OnSave() {

	log.Println("CellFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cellFormCallback.probe.formStage.Checkout()

	if cellFormCallback.cell == nil {
		cellFormCallback.cell = new(models.Cell).Stage(cellFormCallback.probe.stageOfInterest)
	}
	cell_ := cellFormCallback.cell
	_ = cell_

	for _, formDiv := range cellFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cell_.Name), formDiv)
		case "Row:Cells":
			// we need to retrieve the field owner before the change
			var pastRowOwner *models.Row
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Row"
			rf.Fieldname = "Cells"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				cellFormCallback.probe.stageOfInterest,
				cellFormCallback.probe.backRepoOfInterest,
				cell_,
				&rf)

			if reverseFieldOwner != nil {
				pastRowOwner = reverseFieldOwner.(*models.Row)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRowOwner != nil {
					idx := slices.Index(pastRowOwner.Cells, cell_)
					pastRowOwner.Cells = slices.Delete(pastRowOwner.Cells, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _row := range *models.GetGongstructInstancesSet[models.Row](cellFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _row.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRowOwner := _row // we have a match
						if pastRowOwner != nil {
							if newRowOwner != pastRowOwner {
								idx := slices.Index(pastRowOwner.Cells, cell_)
								pastRowOwner.Cells = slices.Delete(pastRowOwner.Cells, idx, idx+1)
								newRowOwner.Cells = append(newRowOwner.Cells, cell_)
							}
						} else {
							newRowOwner.Cells = append(newRowOwner.Cells, cell_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if cellFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cell_.Unstage(cellFormCallback.probe.stageOfInterest)
	}

	cellFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Cell](
		cellFormCallback.probe,
	)
	cellFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if cellFormCallback.CreationMode || cellFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cellFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(cellFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CellFormCallback(
			nil,
			cellFormCallback.probe,
			newFormGroup,
		)
		cell := new(models.Cell)
		FillUpForm(cell, newFormGroup, cellFormCallback.probe)
		cellFormCallback.probe.formStage.Commit()
	}

	fillUpTree(cellFormCallback.probe)
}
func __gong__New__DummyDataFormCallback(
	dummydata *models.DummyData,
	probe *Probe,
	formGroup *table.FormGroup,
) (dummydataFormCallback *DummyDataFormCallback) {
	dummydataFormCallback = new(DummyDataFormCallback)
	dummydataFormCallback.probe = probe
	dummydataFormCallback.dummydata = dummydata
	dummydataFormCallback.formGroup = formGroup

	dummydataFormCallback.CreationMode = (dummydata == nil)

	return
}

type DummyDataFormCallback struct {
	dummydata *models.DummyData

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (dummydataFormCallback *DummyDataFormCallback) OnSave() {

	log.Println("DummyDataFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dummydataFormCallback.probe.formStage.Checkout()

	if dummydataFormCallback.dummydata == nil {
		dummydataFormCallback.dummydata = new(models.DummyData).Stage(dummydataFormCallback.probe.stageOfInterest)
	}
	dummydata_ := dummydataFormCallback.dummydata
	_ = dummydata_

	for _, formDiv := range dummydataFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(dummydata_.Name), formDiv)
		case "DummyString":
			FormDivBasicFieldToField(&(dummydata_.DummyString), formDiv)
		case "DummyInt":
			FormDivBasicFieldToField(&(dummydata_.DummyInt), formDiv)
		case "DummyFloat":
			FormDivBasicFieldToField(&(dummydata_.DummyFloat), formDiv)
		case "DummyBool":
			FormDivBasicFieldToField(&(dummydata_.DummyBool), formDiv)
		case "DummyEnumString":
			FormDivEnumStringFieldToField(&(dummydata_.DummyEnumString), formDiv)
		case "DummyEnumInt":
			FormDivEnumIntFieldToField(&(dummydata_.DummyEnumInt), formDiv)
		case "DummyTime":
			FormDivBasicFieldToField(&(dummydata_.DummyTime), formDiv)
		case "DummyDuration":
			FormDivBasicFieldToField(&(dummydata_.DummyDuration), formDiv)
		case "DummyPointerToGongStruct":
			FormDivSelectFieldToField(&(dummydata_.DummyPointerToGongStruct), dummydataFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if dummydataFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dummydata_.Unstage(dummydataFormCallback.probe.stageOfInterest)
	}

	dummydataFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DummyData](
		dummydataFormCallback.probe,
	)
	dummydataFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if dummydataFormCallback.CreationMode || dummydataFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dummydataFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(dummydataFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DummyDataFormCallback(
			nil,
			dummydataFormCallback.probe,
			newFormGroup,
		)
		dummydata := new(models.DummyData)
		FillUpForm(dummydata, newFormGroup, dummydataFormCallback.probe)
		dummydataFormCallback.probe.formStage.Commit()
	}

	fillUpTree(dummydataFormCallback.probe)
}
func __gong__New__ElementFormCallback(
	element *models.Element,
	probe *Probe,
	formGroup *table.FormGroup,
) (elementFormCallback *ElementFormCallback) {
	elementFormCallback = new(ElementFormCallback)
	elementFormCallback.probe = probe
	elementFormCallback.element = element
	elementFormCallback.formGroup = formGroup

	elementFormCallback.CreationMode = (element == nil)

	return
}

type ElementFormCallback struct {
	element *models.Element

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (elementFormCallback *ElementFormCallback) OnSave() {

	log.Println("ElementFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	elementFormCallback.probe.formStage.Checkout()

	if elementFormCallback.element == nil {
		elementFormCallback.element = new(models.Element).Stage(elementFormCallback.probe.stageOfInterest)
	}
	element_ := elementFormCallback.element
	_ = element_

	for _, formDiv := range elementFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(element_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(element_.Content), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(element_.Type), formDiv)
		case "Element:SubElements":
			// we need to retrieve the field owner before the change
			var pastElementOwner *models.Element
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Element"
			rf.Fieldname = "SubElements"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				elementFormCallback.probe.stageOfInterest,
				elementFormCallback.probe.backRepoOfInterest,
				element_,
				&rf)

			if reverseFieldOwner != nil {
				pastElementOwner = reverseFieldOwner.(*models.Element)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastElementOwner != nil {
					idx := slices.Index(pastElementOwner.SubElements, element_)
					pastElementOwner.SubElements = slices.Delete(pastElementOwner.SubElements, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _element := range *models.GetGongstructInstancesSet[models.Element](elementFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _element.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newElementOwner := _element // we have a match
						if pastElementOwner != nil {
							if newElementOwner != pastElementOwner {
								idx := slices.Index(pastElementOwner.SubElements, element_)
								pastElementOwner.SubElements = slices.Delete(pastElementOwner.SubElements, idx, idx+1)
								newElementOwner.SubElements = append(newElementOwner.SubElements, element_)
							}
						} else {
							newElementOwner.SubElements = append(newElementOwner.SubElements, element_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if elementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		element_.Unstage(elementFormCallback.probe.stageOfInterest)
	}

	elementFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Element](
		elementFormCallback.probe,
	)
	elementFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if elementFormCallback.CreationMode || elementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		elementFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(elementFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ElementFormCallback(
			nil,
			elementFormCallback.probe,
			newFormGroup,
		)
		element := new(models.Element)
		FillUpForm(element, newFormGroup, elementFormCallback.probe)
		elementFormCallback.probe.formStage.Commit()
	}

	fillUpTree(elementFormCallback.probe)
}
func __gong__New__MarkdownContentFormCallback(
	markdowncontent *models.MarkdownContent,
	probe *Probe,
	formGroup *table.FormGroup,
) (markdowncontentFormCallback *MarkdownContentFormCallback) {
	markdowncontentFormCallback = new(MarkdownContentFormCallback)
	markdowncontentFormCallback.probe = probe
	markdowncontentFormCallback.markdowncontent = markdowncontent
	markdowncontentFormCallback.formGroup = formGroup

	markdowncontentFormCallback.CreationMode = (markdowncontent == nil)

	return
}

type MarkdownContentFormCallback struct {
	markdowncontent *models.MarkdownContent

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (markdowncontentFormCallback *MarkdownContentFormCallback) OnSave() {

	log.Println("MarkdownContentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	markdowncontentFormCallback.probe.formStage.Checkout()

	if markdowncontentFormCallback.markdowncontent == nil {
		markdowncontentFormCallback.markdowncontent = new(models.MarkdownContent).Stage(markdowncontentFormCallback.probe.stageOfInterest)
	}
	markdowncontent_ := markdowncontentFormCallback.markdowncontent
	_ = markdowncontent_

	for _, formDiv := range markdowncontentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(markdowncontent_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(markdowncontent_.Content), formDiv)
		case "Root":
			FormDivSelectFieldToField(&(markdowncontent_.Root), markdowncontentFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if markdowncontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		markdowncontent_.Unstage(markdowncontentFormCallback.probe.stageOfInterest)
	}

	markdowncontentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.MarkdownContent](
		markdowncontentFormCallback.probe,
	)
	markdowncontentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if markdowncontentFormCallback.CreationMode || markdowncontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		markdowncontentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(markdowncontentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MarkdownContentFormCallback(
			nil,
			markdowncontentFormCallback.probe,
			newFormGroup,
		)
		markdowncontent := new(models.MarkdownContent)
		FillUpForm(markdowncontent, newFormGroup, markdowncontentFormCallback.probe)
		markdowncontentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(markdowncontentFormCallback.probe)
}
func __gong__New__RowFormCallback(
	row *models.Row,
	probe *Probe,
	formGroup *table.FormGroup,
) (rowFormCallback *RowFormCallback) {
	rowFormCallback = new(RowFormCallback)
	rowFormCallback.probe = probe
	rowFormCallback.row = row
	rowFormCallback.formGroup = formGroup

	rowFormCallback.CreationMode = (row == nil)

	return
}

type RowFormCallback struct {
	row *models.Row

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (rowFormCallback *RowFormCallback) OnSave() {

	log.Println("RowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rowFormCallback.probe.formStage.Checkout()

	if rowFormCallback.row == nil {
		rowFormCallback.row = new(models.Row).Stage(rowFormCallback.probe.stageOfInterest)
	}
	row_ := rowFormCallback.row
	_ = row_

	for _, formDiv := range rowFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(row_.Name), formDiv)
		case "Element:Rows":
			// we need to retrieve the field owner before the change
			var pastElementOwner *models.Element
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Element"
			rf.Fieldname = "Rows"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				rowFormCallback.probe.stageOfInterest,
				rowFormCallback.probe.backRepoOfInterest,
				row_,
				&rf)

			if reverseFieldOwner != nil {
				pastElementOwner = reverseFieldOwner.(*models.Element)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastElementOwner != nil {
					idx := slices.Index(pastElementOwner.Rows, row_)
					pastElementOwner.Rows = slices.Delete(pastElementOwner.Rows, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _element := range *models.GetGongstructInstancesSet[models.Element](rowFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _element.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newElementOwner := _element // we have a match
						if pastElementOwner != nil {
							if newElementOwner != pastElementOwner {
								idx := slices.Index(pastElementOwner.Rows, row_)
								pastElementOwner.Rows = slices.Delete(pastElementOwner.Rows, idx, idx+1)
								newElementOwner.Rows = append(newElementOwner.Rows, row_)
							}
						} else {
							newElementOwner.Rows = append(newElementOwner.Rows, row_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if rowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		row_.Unstage(rowFormCallback.probe.stageOfInterest)
	}

	rowFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Row](
		rowFormCallback.probe,
	)
	rowFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if rowFormCallback.CreationMode || rowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rowFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(rowFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RowFormCallback(
			nil,
			rowFormCallback.probe,
			newFormGroup,
		)
		row := new(models.Row)
		FillUpForm(row, newFormGroup, rowFormCallback.probe)
		rowFormCallback.probe.formStage.Commit()
	}

	fillUpTree(rowFormCallback.probe)
}
