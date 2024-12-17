package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongmarkdown/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__Field__000000_Content := (&models.Field{}).Stage(stage)
	__Field__000001_Content := (&models.Field{}).Stage(stage)
	__Field__000002_Type := (&models.Field{}).Stage(stage)

	__GongEnumShape__000000_Default_ElementType := (&models.GongEnumShape{}).Stage(stage)

	__GongEnumValueEntry__000000_PARAGRAPH := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000001_TABLE := (&models.GongEnumValueEntry{}).Stage(stage)
	__GongEnumValueEntry__000002_TITLE := (&models.GongEnumValueEntry{}).Stage(stage)

	__GongStructShape__000000_Default_Cell := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Element := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_MarkdownContent := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_Row := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_Cells := (&models.Link{}).Stage(stage)
	__Link__000001_Root := (&models.Link{}).Stage(stage)
	__Link__000002_Rows := (&models.Link{}).Stage(stage)
	__Link__000003_SubElements := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_Default_Cell := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_Element := (&models.Position{}).Stage(stage)
	__Position__000002_Pos_Default_ElementType := (&models.Position{}).Stage(stage)
	__Position__000003_Pos_Default_MarkdownContent := (&models.Position{}).Stage(stage)
	__Position__000004_Pos_Default_Row := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_Element := (&models.Vertice{}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_Row := (&models.Vertice{}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_MarkdownContent_and_Default_Element := (&models.Vertice{}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Content.Name = `Content`

	//gong:ident [ref_models.MarkdownContent.Content] comment added to overcome the problem with the comment map association
	__Field__000000_Content.Identifier = `ref_models.MarkdownContent.Content`
	__Field__000000_Content.FieldTypeAsString = ``
	__Field__000000_Content.Structname = `MarkdownContent`
	__Field__000000_Content.Fieldtypename = `string`

	__Field__000001_Content.Name = `Content`

	//gong:ident [ref_models.Element.Content] comment added to overcome the problem with the comment map association
	__Field__000001_Content.Identifier = `ref_models.Element.Content`
	__Field__000001_Content.FieldTypeAsString = ``
	__Field__000001_Content.Structname = `Element`
	__Field__000001_Content.Fieldtypename = `string`

	__Field__000002_Type.Name = `Type`

	//gong:ident [ref_models.Element.Type] comment added to overcome the problem with the comment map association
	__Field__000002_Type.Identifier = `ref_models.Element.Type`
	__Field__000002_Type.FieldTypeAsString = ``
	__Field__000002_Type.Structname = `Element`
	__Field__000002_Type.Fieldtypename = `ElementType`

	__GongEnumShape__000000_Default_ElementType.Name = `Default-ElementType`

	//gong:ident [ref_models.ElementType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000000_Default_ElementType.Identifier = `ref_models.ElementType`
	__GongEnumShape__000000_Default_ElementType.Width = 240.000000
	__GongEnumShape__000000_Default_ElementType.Height = 108.000000

	__GongEnumValueEntry__000000_PARAGRAPH.Name = `PARAGRAPH`

	//gong:ident [ref_models.ElementType.PARAGRAPH] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000000_PARAGRAPH.Identifier = `ref_models.ElementType.PARAGRAPH`

	__GongEnumValueEntry__000001_TABLE.Name = `TABLE`

	//gong:ident [ref_models.ElementType.TABLE] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000001_TABLE.Identifier = `ref_models.ElementType.TABLE`

	__GongEnumValueEntry__000002_TITLE.Name = `TITLE`

	//gong:ident [ref_models.ElementType.TITLE] comment added to overcome the problem with the comment map association
	__GongEnumValueEntry__000002_TITLE.Identifier = `ref_models.ElementType.TITLE`

	__GongStructShape__000000_Default_Cell.Name = `Default-Cell`

	//gong:ident [ref_models.Cell] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Cell.Identifier = `ref_models.Cell`
	__GongStructShape__000000_Default_Cell.ShowNbInstances = false
	__GongStructShape__000000_Default_Cell.NbInstances = 0
	__GongStructShape__000000_Default_Cell.Width = 240.000000
	__GongStructShape__000000_Default_Cell.Height = 63.000000
	__GongStructShape__000000_Default_Cell.IsSelected = false

	__GongStructShape__000001_Default_Element.Name = `Default-Element`

	//gong:ident [ref_models.Element] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Element.Identifier = `ref_models.Element`
	__GongStructShape__000001_Default_Element.ShowNbInstances = false
	__GongStructShape__000001_Default_Element.NbInstances = 0
	__GongStructShape__000001_Default_Element.Width = 240.000000
	__GongStructShape__000001_Default_Element.Height = 93.000000
	__GongStructShape__000001_Default_Element.IsSelected = false

	__GongStructShape__000002_Default_MarkdownContent.Name = `Default-MarkdownContent`

	//gong:ident [ref_models.MarkdownContent] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_MarkdownContent.Identifier = `ref_models.MarkdownContent`
	__GongStructShape__000002_Default_MarkdownContent.ShowNbInstances = false
	__GongStructShape__000002_Default_MarkdownContent.NbInstances = 0
	__GongStructShape__000002_Default_MarkdownContent.Width = 240.000000
	__GongStructShape__000002_Default_MarkdownContent.Height = 78.000000
	__GongStructShape__000002_Default_MarkdownContent.IsSelected = false

	__GongStructShape__000003_Default_Row.Name = `Default-Row`

	//gong:ident [ref_models.Row] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_Row.Identifier = `ref_models.Row`
	__GongStructShape__000003_Default_Row.ShowNbInstances = false
	__GongStructShape__000003_Default_Row.NbInstances = 0
	__GongStructShape__000003_Default_Row.Width = 240.000000
	__GongStructShape__000003_Default_Row.Height = 63.000000
	__GongStructShape__000003_Default_Row.IsSelected = false

	__Link__000000_Cells.Name = `Cells`

	//gong:ident [ref_models.Row.Cells] comment added to overcome the problem with the comment map association
	__Link__000000_Cells.Identifier = `ref_models.Row.Cells`

	//gong:ident [ref_models.Cell] comment added to overcome the problem with the comment map association
	__Link__000000_Cells.Fieldtypename = `ref_models.Cell`
	__Link__000000_Cells.FieldOffsetX = -50.000000
	__Link__000000_Cells.FieldOffsetY = -16.000000
	__Link__000000_Cells.TargetMultiplicity = models.MANY
	__Link__000000_Cells.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_Cells.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_Cells.SourceMultiplicity = models.MANY
	__Link__000000_Cells.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_Cells.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_Cells.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Cells.StartRatio = 0.500000
	__Link__000000_Cells.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Cells.EndRatio = 0.500000
	__Link__000000_Cells.CornerOffsetRatio = 1.380000

	__Link__000001_Root.Name = `Root`

	//gong:ident [ref_models.MarkdownContent.Root] comment added to overcome the problem with the comment map association
	__Link__000001_Root.Identifier = `ref_models.MarkdownContent.Root`

	//gong:ident [ref_models.Element] comment added to overcome the problem with the comment map association
	__Link__000001_Root.Fieldtypename = `ref_models.Element`
	__Link__000001_Root.FieldOffsetX = -50.000000
	__Link__000001_Root.FieldOffsetY = -16.000000
	__Link__000001_Root.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_Root.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_Root.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_Root.SourceMultiplicity = models.MANY
	__Link__000001_Root.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_Root.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_Root.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Root.StartRatio = 0.500000
	__Link__000001_Root.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Root.EndRatio = 0.500000
	__Link__000001_Root.CornerOffsetRatio = 1.380000

	__Link__000002_Rows.Name = `Rows`

	//gong:ident [ref_models.Element.Rows] comment added to overcome the problem with the comment map association
	__Link__000002_Rows.Identifier = `ref_models.Element.Rows`

	//gong:ident [ref_models.Row] comment added to overcome the problem with the comment map association
	__Link__000002_Rows.Fieldtypename = `ref_models.Row`
	__Link__000002_Rows.FieldOffsetX = -50.000000
	__Link__000002_Rows.FieldOffsetY = -16.000000
	__Link__000002_Rows.TargetMultiplicity = models.MANY
	__Link__000002_Rows.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_Rows.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_Rows.SourceMultiplicity = models.MANY
	__Link__000002_Rows.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_Rows.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_Rows.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_Rows.StartRatio = 0.418055
	__Link__000002_Rows.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Rows.EndRatio = 0.500000
	__Link__000002_Rows.CornerOffsetRatio = 2.071684

	__Link__000003_SubElements.Name = `SubElements`

	//gong:ident [ref_models.Element.SubElements] comment added to overcome the problem with the comment map association
	__Link__000003_SubElements.Identifier = `ref_models.Element.SubElements`

	//gong:ident [ref_models.Element] comment added to overcome the problem with the comment map association
	__Link__000003_SubElements.Fieldtypename = `ref_models.Element`
	__Link__000003_SubElements.FieldOffsetX = -50.000000
	__Link__000003_SubElements.FieldOffsetY = -16.000000
	__Link__000003_SubElements.TargetMultiplicity = models.MANY
	__Link__000003_SubElements.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_SubElements.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_SubElements.SourceMultiplicity = models.MANY
	__Link__000003_SubElements.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_SubElements.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_SubElements.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_SubElements.StartRatio = 0.125448
	__Link__000003_SubElements.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_SubElements.EndRatio = 1.000000
	__Link__000003_SubElements.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_Cell.X = 1070.999878
	__Position__000000_Pos_Default_Cell.Y = 273.000000
	__Position__000000_Pos_Default_Cell.Name = `Pos-Default-Cell`

	__Position__000001_Pos_Default_Element.X = 465.000000
	__Position__000001_Pos_Default_Element.Y = 29.000000
	__Position__000001_Pos_Default_Element.Name = `Pos-Default-Element`

	__Position__000002_Pos_Default_ElementType.X = 887.999939
	__Position__000002_Pos_Default_ElementType.Y = 53.000000
	__Position__000002_Pos_Default_ElementType.Name = `Pos-Default-ElementType`

	__Position__000003_Pos_Default_MarkdownContent.X = 17.000000
	__Position__000003_Pos_Default_MarkdownContent.Y = 26.000000
	__Position__000003_Pos_Default_MarkdownContent.Name = `Pos-Default-MarkdownContent`

	__Position__000004_Pos_Default_Row.X = 676.000000
	__Position__000004_Pos_Default_Row.Y = 276.999969
	__Position__000004_Pos_Default_Row.Name = `Pos-Default-Row`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_Element.X = 825.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_Element.Y = 75.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_Element.Name = `Verticle in class diagram Default in middle between Default-Element and Default-Element`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_Row.X = 798.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_Row.Y = 293.999985
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_Row.Name = `Verticle in class diagram Default in middle between Default-Element and Default-Row`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_MarkdownContent_and_Default_Element.X = 377.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_MarkdownContent_and_Default_Element.Y = 62.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_MarkdownContent_and_Default_Element.Name = `Verticle in class diagram Default in middle between Default-MarkdownContent and Default-Element`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell.X = 1233.499939
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell.Y = 306.499985
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell.Name = `Verticle in class diagram Default in middle between Default-Row and Default-Cell`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_MarkdownContent)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Element)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Row)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Cell)
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000000_Default_ElementType)
	__GongEnumShape__000000_Default_ElementType.Position = __Position__000002_Pos_Default_ElementType
	__GongEnumShape__000000_Default_ElementType.GongEnumValueEntrys = append(__GongEnumShape__000000_Default_ElementType.GongEnumValueEntrys, __GongEnumValueEntry__000000_PARAGRAPH)
	__GongEnumShape__000000_Default_ElementType.GongEnumValueEntrys = append(__GongEnumShape__000000_Default_ElementType.GongEnumValueEntrys, __GongEnumValueEntry__000002_TITLE)
	__GongEnumShape__000000_Default_ElementType.GongEnumValueEntrys = append(__GongEnumShape__000000_Default_ElementType.GongEnumValueEntrys, __GongEnumValueEntry__000001_TABLE)
	__GongStructShape__000000_Default_Cell.Position = __Position__000000_Pos_Default_Cell
	__GongStructShape__000001_Default_Element.Position = __Position__000001_Pos_Default_Element
	__GongStructShape__000001_Default_Element.Fields = append(__GongStructShape__000001_Default_Element.Fields, __Field__000001_Content)
	__GongStructShape__000001_Default_Element.Fields = append(__GongStructShape__000001_Default_Element.Fields, __Field__000002_Type)
	__GongStructShape__000001_Default_Element.Links = append(__GongStructShape__000001_Default_Element.Links, __Link__000003_SubElements)
	__GongStructShape__000001_Default_Element.Links = append(__GongStructShape__000001_Default_Element.Links, __Link__000002_Rows)
	__GongStructShape__000002_Default_MarkdownContent.Position = __Position__000003_Pos_Default_MarkdownContent
	__GongStructShape__000002_Default_MarkdownContent.Fields = append(__GongStructShape__000002_Default_MarkdownContent.Fields, __Field__000000_Content)
	__GongStructShape__000002_Default_MarkdownContent.Links = append(__GongStructShape__000002_Default_MarkdownContent.Links, __Link__000001_Root)
	__GongStructShape__000003_Default_Row.Position = __Position__000004_Pos_Default_Row
	__GongStructShape__000003_Default_Row.Links = append(__GongStructShape__000003_Default_Row.Links, __Link__000000_Cells)
	__Link__000000_Cells.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Row_and_Default_Cell
	__Link__000001_Root.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_MarkdownContent_and_Default_Element
	__Link__000002_Rows.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_Row
	__Link__000003_SubElements.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_Element
}
