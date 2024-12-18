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
func __gong__New__ContentFormCallback(
	content *models.Content,
	probe *Probe,
	formGroup *table.FormGroup,
) (contentFormCallback *ContentFormCallback) {
	contentFormCallback = new(ContentFormCallback)
	contentFormCallback.probe = probe
	contentFormCallback.content = content
	contentFormCallback.formGroup = formGroup

	contentFormCallback.CreationMode = (content == nil)

	return
}

type ContentFormCallback struct {
	content *models.Content

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (contentFormCallback *ContentFormCallback) OnSave() {

	log.Println("ContentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	contentFormCallback.probe.formStage.Checkout()

	if contentFormCallback.content == nil {
		contentFormCallback.content = new(models.Content).Stage(contentFormCallback.probe.stageOfInterest)
	}
	content_ := contentFormCallback.content
	_ = content_

	for _, formDiv := range contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(content_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(content_.Content), formDiv)
		}
	}

	// manage the suppress operation
	if contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		content_.Unstage(contentFormCallback.probe.stageOfInterest)
	}

	contentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Content](
		contentFormCallback.probe,
	)
	contentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if contentFormCallback.CreationMode || contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ContentFormCallback(
			nil,
			contentFormCallback.probe,
			newFormGroup,
		)
		content := new(models.Content)
		FillUpForm(content, newFormGroup, contentFormCallback.probe)
		contentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(contentFormCallback.probe)
}
