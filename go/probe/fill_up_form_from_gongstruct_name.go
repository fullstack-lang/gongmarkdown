// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongmarkdown/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "AnotherDummyData":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "AnotherDummyData Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnotherDummyDataFormCallback(
			nil,
			probe,
			formGroup,
		)
		anotherdummydata := new(models.AnotherDummyData)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(anotherdummydata, formGroup, probe)
	case "Cell":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Cell Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellFormCallback(
			nil,
			probe,
			formGroup,
		)
		cell := new(models.Cell)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cell, formGroup, probe)
	case "DummyData":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DummyData Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DummyDataFormCallback(
			nil,
			probe,
			formGroup,
		)
		dummydata := new(models.DummyData)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(dummydata, formGroup, probe)
	case "Element":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Element Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ElementFormCallback(
			nil,
			probe,
			formGroup,
		)
		element := new(models.Element)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(element, formGroup, probe)
	case "MarkdownContent":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "MarkdownContent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MarkdownContentFormCallback(
			nil,
			probe,
			formGroup,
		)
		markdowncontent := new(models.MarkdownContent)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(markdowncontent, formGroup, probe)
	case "Row":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Row Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RowFormCallback(
			nil,
			probe,
			formGroup,
		)
		row := new(models.Row)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(row, formGroup, probe)
	}
	formStage.Commit()
}
