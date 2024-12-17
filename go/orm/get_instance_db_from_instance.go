// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongmarkdown/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.AnotherDummyData:
		anotherdummydataInstance := any(concreteInstance).(*models.AnotherDummyData)
		ret2 := backRepo.BackRepoAnotherDummyData.GetAnotherDummyDataDBFromAnotherDummyDataPtr(anotherdummydataInstance)
		ret = any(ret2).(*T2)
	case *models.Cell:
		cellInstance := any(concreteInstance).(*models.Cell)
		ret2 := backRepo.BackRepoCell.GetCellDBFromCellPtr(cellInstance)
		ret = any(ret2).(*T2)
	case *models.DummyData:
		dummydataInstance := any(concreteInstance).(*models.DummyData)
		ret2 := backRepo.BackRepoDummyData.GetDummyDataDBFromDummyDataPtr(dummydataInstance)
		ret = any(ret2).(*T2)
	case *models.Element:
		elementInstance := any(concreteInstance).(*models.Element)
		ret2 := backRepo.BackRepoElement.GetElementDBFromElementPtr(elementInstance)
		ret = any(ret2).(*T2)
	case *models.MarkdownContent:
		markdowncontentInstance := any(concreteInstance).(*models.MarkdownContent)
		ret2 := backRepo.BackRepoMarkdownContent.GetMarkdownContentDBFromMarkdownContentPtr(markdowncontentInstance)
		ret = any(ret2).(*T2)
	case *models.Row:
		rowInstance := any(concreteInstance).(*models.Row)
		ret2 := backRepo.BackRepoRow.GetRowDBFromRowPtr(rowInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.AnotherDummyData:
		tmp := GetInstanceDBFromInstance[models.AnotherDummyData, AnotherDummyDataDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Cell:
		tmp := GetInstanceDBFromInstance[models.Cell, CellDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DummyData:
		tmp := GetInstanceDBFromInstance[models.DummyData, DummyDataDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Element:
		tmp := GetInstanceDBFromInstance[models.Element, ElementDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MarkdownContent:
		tmp := GetInstanceDBFromInstance[models.MarkdownContent, MarkdownContentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Row:
		tmp := GetInstanceDBFromInstance[models.Row, RowDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.AnotherDummyData:
		tmp := GetInstanceDBFromInstance[models.AnotherDummyData, AnotherDummyDataDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Cell:
		tmp := GetInstanceDBFromInstance[models.Cell, CellDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DummyData:
		tmp := GetInstanceDBFromInstance[models.DummyData, DummyDataDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Element:
		tmp := GetInstanceDBFromInstance[models.Element, ElementDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MarkdownContent:
		tmp := GetInstanceDBFromInstance[models.MarkdownContent, MarkdownContentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Row:
		tmp := GetInstanceDBFromInstance[models.Row, RowDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
