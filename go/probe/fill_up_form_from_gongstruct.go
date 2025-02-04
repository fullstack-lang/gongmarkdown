// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongmarkdown/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Content:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Content Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ContentFormCallback(
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
