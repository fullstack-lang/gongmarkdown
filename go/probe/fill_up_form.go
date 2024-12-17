// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongmarkdown/go/models"
	"github.com/fullstack-lang/gongmarkdown/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.AnotherDummyData:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Cell:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Row"
			rf.Fieldname = "Cells"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Row),
					"Cells",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Row, *models.Cell](
					nil,
					"Cells",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DummyData:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DummyString", instanceWithInferedType.DummyString, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DummyInt", instanceWithInferedType.DummyInt, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DummyFloat", instanceWithInferedType.DummyFloat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DummyBool", instanceWithInferedType.DummyBool, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("DummyEnumString", instanceWithInferedType.DummyEnumString, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeIntToForm("DummyEnumInt", instanceWithInferedType.DummyEnumInt, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("DummyTime", instanceWithInferedType.DummyTime, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DummyDuration", instanceWithInferedType.DummyDuration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DummyPointerToGongStruct", instanceWithInferedType.DummyPointerToGongStruct, formGroup, probe)

	case *models.Element:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("SubElements", instanceWithInferedType, &instanceWithInferedType.SubElements, formGroup, probe)
		AssociationSliceToForm("Rows", instanceWithInferedType, &instanceWithInferedType.Rows, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Element"
			rf.Fieldname = "SubElements"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Element),
					"SubElements",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Element, *models.Element](
					nil,
					"SubElements",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.MarkdownContent:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		AssociationFieldToForm("Root", instanceWithInferedType.Root, formGroup, probe)

	case *models.Row:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Cells", instanceWithInferedType, &instanceWithInferedType.Cells, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Element"
			rf.Fieldname = "Rows"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Element),
					"Rows",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Element, *models.Row](
					nil,
					"Rows",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
