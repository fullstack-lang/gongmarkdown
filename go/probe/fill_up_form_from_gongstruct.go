// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongmarkdown/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.AnotherDummyData:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "AnotherDummyData Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnotherDummyDataFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Cell:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Cell Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CellFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DummyData:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DummyData Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DummyDataFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Element:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Element Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ElementFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MarkdownContent:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "MarkdownContent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MarkdownContentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Row:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Row Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RowFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
