// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_AnotherDummyData_Identifiers := make(map[*AnotherDummyData]string)
	_ = map_AnotherDummyData_Identifiers

	anotherdummydataOrdered := []*AnotherDummyData{}
	for anotherdummydata := range stage.AnotherDummyDatas {
		anotherdummydataOrdered = append(anotherdummydataOrdered, anotherdummydata)
	}
	sort.Slice(anotherdummydataOrdered[:], func(i, j int) bool {
		return anotherdummydataOrdered[i].Name < anotherdummydataOrdered[j].Name
	})
	if len(anotherdummydataOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, anotherdummydata := range anotherdummydataOrdered {

		id = generatesIdentifier("AnotherDummyData", idx, anotherdummydata.Name)
		map_AnotherDummyData_Identifiers[anotherdummydata] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AnotherDummyData")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", anotherdummydata.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(anotherdummydata.Name))
		initializerStatements += setValueField

	}

	map_Cell_Identifiers := make(map[*Cell]string)
	_ = map_Cell_Identifiers

	cellOrdered := []*Cell{}
	for cell := range stage.Cells {
		cellOrdered = append(cellOrdered, cell)
	}
	sort.Slice(cellOrdered[:], func(i, j int) bool {
		return cellOrdered[i].Name < cellOrdered[j].Name
	})
	if len(cellOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, cell := range cellOrdered {

		id = generatesIdentifier("Cell", idx, cell.Name)
		map_Cell_Identifiers[cell] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Cell")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", cell.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cell.Name))
		initializerStatements += setValueField

	}

	map_DummyData_Identifiers := make(map[*DummyData]string)
	_ = map_DummyData_Identifiers

	dummydataOrdered := []*DummyData{}
	for dummydata := range stage.DummyDatas {
		dummydataOrdered = append(dummydataOrdered, dummydata)
	}
	sort.Slice(dummydataOrdered[:], func(i, j int) bool {
		return dummydataOrdered[i].Name < dummydataOrdered[j].Name
	})
	if len(dummydataOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, dummydata := range dummydataOrdered {

		id = generatesIdentifier("DummyData", idx, dummydata.Name)
		map_DummyData_Identifiers[dummydata] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DummyData")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", dummydata.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dummydata.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DummyString")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dummydata.DummyString))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DummyInt")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", dummydata.DummyInt))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DummyFloat")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", dummydata.DummyFloat))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DummyBool")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", dummydata.DummyBool))
		initializerStatements += setValueField

		if dummydata.DummyEnumString != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DummyEnumString")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+dummydata.DummyEnumString.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DummyEnumInt")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+dummydata.DummyEnumInt.ToCodeString())
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DummyTime")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", dummydata.DummyTime.String())
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DummyDuration")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", dummydata.DummyDuration))
		initializerStatements += setValueField

	}

	map_Element_Identifiers := make(map[*Element]string)
	_ = map_Element_Identifiers

	elementOrdered := []*Element{}
	for element := range stage.Elements {
		elementOrdered = append(elementOrdered, element)
	}
	sort.Slice(elementOrdered[:], func(i, j int) bool {
		return elementOrdered[i].Name < elementOrdered[j].Name
	})
	if len(elementOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, element := range elementOrdered {

		id = generatesIdentifier("Element", idx, element.Name)
		map_Element_Identifiers[element] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Element")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", element.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Content))
		initializerStatements += setValueField

		if element.Type != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+element.Type.ToCodeString())
			initializerStatements += setValueField
		}

	}

	map_MarkdownContent_Identifiers := make(map[*MarkdownContent]string)
	_ = map_MarkdownContent_Identifiers

	markdowncontentOrdered := []*MarkdownContent{}
	for markdowncontent := range stage.MarkdownContents {
		markdowncontentOrdered = append(markdowncontentOrdered, markdowncontent)
	}
	sort.Slice(markdowncontentOrdered[:], func(i, j int) bool {
		return markdowncontentOrdered[i].Name < markdowncontentOrdered[j].Name
	})
	if len(markdowncontentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, markdowncontent := range markdowncontentOrdered {

		id = generatesIdentifier("MarkdownContent", idx, markdowncontent.Name)
		map_MarkdownContent_Identifiers[markdowncontent] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MarkdownContent")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", markdowncontent.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(markdowncontent.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(markdowncontent.Content))
		initializerStatements += setValueField

	}

	map_Row_Identifiers := make(map[*Row]string)
	_ = map_Row_Identifiers

	rowOrdered := []*Row{}
	for row := range stage.Rows {
		rowOrdered = append(rowOrdered, row)
	}
	sort.Slice(rowOrdered[:], func(i, j int) bool {
		return rowOrdered[i].Name < rowOrdered[j].Name
	})
	if len(rowOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, row := range rowOrdered {

		id = generatesIdentifier("Row", idx, row.Name)
		map_Row_Identifiers[row] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Row")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", row.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(row.Name))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, anotherdummydata := range anotherdummydataOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AnotherDummyData", idx, anotherdummydata.Name)
		map_AnotherDummyData_Identifiers[anotherdummydata] = id

		// Initialisation of values
	}

	for idx, cell := range cellOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Cell", idx, cell.Name)
		map_Cell_Identifiers[cell] = id

		// Initialisation of values
	}

	for idx, dummydata := range dummydataOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DummyData", idx, dummydata.Name)
		map_DummyData_Identifiers[dummydata] = id

		// Initialisation of values
		if dummydata.DummyPointerToGongStruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DummyPointerToGongStruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AnotherDummyData_Identifiers[dummydata.DummyPointerToGongStruct])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, element := range elementOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Element", idx, element.Name)
		map_Element_Identifiers[element] = id

		// Initialisation of values
		for _, _element := range element.SubElements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SubElements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

		for _, _row := range element.Rows {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Rows")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Row_Identifiers[_row])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, markdowncontent := range markdowncontentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("MarkdownContent", idx, markdowncontent.Name)
		map_MarkdownContent_Identifiers[markdowncontent] = id

		// Initialisation of values
		if markdowncontent.Root != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Root")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[markdowncontent.Root])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, row := range rowOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Row", idx, row.Name)
		map_Row_Identifiers[row] = id

		// Initialisation of values
		for _, _cell := range row.Cells {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Cells")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Cell_Identifiers[_cell])
			pointersInitializesStatements += setPointerField
		}

	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.StageStruct",
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
