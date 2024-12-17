// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	AnotherDummyDatas           map[*AnotherDummyData]any
	AnotherDummyDatas_mapString map[string]*AnotherDummyData

	// insertion point for slice of pointers maps
	OnAfterAnotherDummyDataCreateCallback OnAfterCreateInterface[AnotherDummyData]
	OnAfterAnotherDummyDataUpdateCallback OnAfterUpdateInterface[AnotherDummyData]
	OnAfterAnotherDummyDataDeleteCallback OnAfterDeleteInterface[AnotherDummyData]
	OnAfterAnotherDummyDataReadCallback   OnAfterReadInterface[AnotherDummyData]

	Cells           map[*Cell]any
	Cells_mapString map[string]*Cell

	// insertion point for slice of pointers maps
	OnAfterCellCreateCallback OnAfterCreateInterface[Cell]
	OnAfterCellUpdateCallback OnAfterUpdateInterface[Cell]
	OnAfterCellDeleteCallback OnAfterDeleteInterface[Cell]
	OnAfterCellReadCallback   OnAfterReadInterface[Cell]

	DummyDatas           map[*DummyData]any
	DummyDatas_mapString map[string]*DummyData

	// insertion point for slice of pointers maps
	OnAfterDummyDataCreateCallback OnAfterCreateInterface[DummyData]
	OnAfterDummyDataUpdateCallback OnAfterUpdateInterface[DummyData]
	OnAfterDummyDataDeleteCallback OnAfterDeleteInterface[DummyData]
	OnAfterDummyDataReadCallback   OnAfterReadInterface[DummyData]

	Elements           map[*Element]any
	Elements_mapString map[string]*Element

	// insertion point for slice of pointers maps
	Element_SubElements_reverseMap map[*Element]*Element

	Element_Rows_reverseMap map[*Row]*Element

	OnAfterElementCreateCallback OnAfterCreateInterface[Element]
	OnAfterElementUpdateCallback OnAfterUpdateInterface[Element]
	OnAfterElementDeleteCallback OnAfterDeleteInterface[Element]
	OnAfterElementReadCallback   OnAfterReadInterface[Element]

	MarkdownContents           map[*MarkdownContent]any
	MarkdownContents_mapString map[string]*MarkdownContent

	// insertion point for slice of pointers maps
	OnAfterMarkdownContentCreateCallback OnAfterCreateInterface[MarkdownContent]
	OnAfterMarkdownContentUpdateCallback OnAfterUpdateInterface[MarkdownContent]
	OnAfterMarkdownContentDeleteCallback OnAfterDeleteInterface[MarkdownContent]
	OnAfterMarkdownContentReadCallback   OnAfterReadInterface[MarkdownContent]

	Rows           map[*Row]any
	Rows_mapString map[string]*Row

	// insertion point for slice of pointers maps
	Row_Cells_reverseMap map[*Cell]*Row

	OnAfterRowCreateCallback OnAfterCreateInterface[Row]
	OnAfterRowUpdateCallback OnAfterUpdateInterface[Row]
	OnAfterRowDeleteCallback OnAfterDeleteInterface[Row]
	OnAfterRowReadCallback   OnAfterReadInterface[Row]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gongmarkdown/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitAnotherDummyData(anotherdummydata *AnotherDummyData)
	CheckoutAnotherDummyData(anotherdummydata *AnotherDummyData)
	CommitCell(cell *Cell)
	CheckoutCell(cell *Cell)
	CommitDummyData(dummydata *DummyData)
	CheckoutDummyData(dummydata *DummyData)
	CommitElement(element *Element)
	CheckoutElement(element *Element)
	CommitMarkdownContent(markdowncontent *MarkdownContent)
	CheckoutMarkdownContent(markdowncontent *MarkdownContent)
	CommitRow(row *Row)
	CheckoutRow(row *Row)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		AnotherDummyDatas:           make(map[*AnotherDummyData]any),
		AnotherDummyDatas_mapString: make(map[string]*AnotherDummyData),

		Cells:           make(map[*Cell]any),
		Cells_mapString: make(map[string]*Cell),

		DummyDatas:           make(map[*DummyData]any),
		DummyDatas_mapString: make(map[string]*DummyData),

		Elements:           make(map[*Element]any),
		Elements_mapString: make(map[string]*Element),

		MarkdownContents:           make(map[*MarkdownContent]any),
		MarkdownContents_mapString: make(map[string]*MarkdownContent),

		Rows:           make(map[*Row]any),
		Rows_mapString: make(map[string]*Row),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["AnotherDummyData"] = len(stage.AnotherDummyDatas)
	stage.Map_GongStructName_InstancesNb["Cell"] = len(stage.Cells)
	stage.Map_GongStructName_InstancesNb["DummyData"] = len(stage.DummyDatas)
	stage.Map_GongStructName_InstancesNb["Element"] = len(stage.Elements)
	stage.Map_GongStructName_InstancesNb["MarkdownContent"] = len(stage.MarkdownContents)
	stage.Map_GongStructName_InstancesNb["Row"] = len(stage.Rows)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["AnotherDummyData"] = len(stage.AnotherDummyDatas)
	stage.Map_GongStructName_InstancesNb["Cell"] = len(stage.Cells)
	stage.Map_GongStructName_InstancesNb["DummyData"] = len(stage.DummyDatas)
	stage.Map_GongStructName_InstancesNb["Element"] = len(stage.Elements)
	stage.Map_GongStructName_InstancesNb["MarkdownContent"] = len(stage.MarkdownContents)
	stage.Map_GongStructName_InstancesNb["Row"] = len(stage.Rows)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts anotherdummydata to the model stage
func (anotherdummydata *AnotherDummyData) Stage(stage *StageStruct) *AnotherDummyData {
	stage.AnotherDummyDatas[anotherdummydata] = __member
	stage.AnotherDummyDatas_mapString[anotherdummydata.Name] = anotherdummydata

	return anotherdummydata
}

// Unstage removes anotherdummydata off the model stage
func (anotherdummydata *AnotherDummyData) Unstage(stage *StageStruct) *AnotherDummyData {
	delete(stage.AnotherDummyDatas, anotherdummydata)
	delete(stage.AnotherDummyDatas_mapString, anotherdummydata.Name)
	return anotherdummydata
}

// UnstageVoid removes anotherdummydata off the model stage
func (anotherdummydata *AnotherDummyData) UnstageVoid(stage *StageStruct) {
	delete(stage.AnotherDummyDatas, anotherdummydata)
	delete(stage.AnotherDummyDatas_mapString, anotherdummydata.Name)
}

// commit anotherdummydata to the back repo (if it is already staged)
func (anotherdummydata *AnotherDummyData) Commit(stage *StageStruct) *AnotherDummyData {
	if _, ok := stage.AnotherDummyDatas[anotherdummydata]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAnotherDummyData(anotherdummydata)
		}
	}
	return anotherdummydata
}

func (anotherdummydata *AnotherDummyData) CommitVoid(stage *StageStruct) {
	anotherdummydata.Commit(stage)
}

// Checkout anotherdummydata to the back repo (if it is already staged)
func (anotherdummydata *AnotherDummyData) Checkout(stage *StageStruct) *AnotherDummyData {
	if _, ok := stage.AnotherDummyDatas[anotherdummydata]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAnotherDummyData(anotherdummydata)
		}
	}
	return anotherdummydata
}

// for satisfaction of GongStruct interface
func (anotherdummydata *AnotherDummyData) GetName() (res string) {
	return anotherdummydata.Name
}

// Stage puts cell to the model stage
func (cell *Cell) Stage(stage *StageStruct) *Cell {
	stage.Cells[cell] = __member
	stage.Cells_mapString[cell.Name] = cell

	return cell
}

// Unstage removes cell off the model stage
func (cell *Cell) Unstage(stage *StageStruct) *Cell {
	delete(stage.Cells, cell)
	delete(stage.Cells_mapString, cell.Name)
	return cell
}

// UnstageVoid removes cell off the model stage
func (cell *Cell) UnstageVoid(stage *StageStruct) {
	delete(stage.Cells, cell)
	delete(stage.Cells_mapString, cell.Name)
}

// commit cell to the back repo (if it is already staged)
func (cell *Cell) Commit(stage *StageStruct) *Cell {
	if _, ok := stage.Cells[cell]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCell(cell)
		}
	}
	return cell
}

func (cell *Cell) CommitVoid(stage *StageStruct) {
	cell.Commit(stage)
}

// Checkout cell to the back repo (if it is already staged)
func (cell *Cell) Checkout(stage *StageStruct) *Cell {
	if _, ok := stage.Cells[cell]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCell(cell)
		}
	}
	return cell
}

// for satisfaction of GongStruct interface
func (cell *Cell) GetName() (res string) {
	return cell.Name
}

// Stage puts dummydata to the model stage
func (dummydata *DummyData) Stage(stage *StageStruct) *DummyData {
	stage.DummyDatas[dummydata] = __member
	stage.DummyDatas_mapString[dummydata.Name] = dummydata

	return dummydata
}

// Unstage removes dummydata off the model stage
func (dummydata *DummyData) Unstage(stage *StageStruct) *DummyData {
	delete(stage.DummyDatas, dummydata)
	delete(stage.DummyDatas_mapString, dummydata.Name)
	return dummydata
}

// UnstageVoid removes dummydata off the model stage
func (dummydata *DummyData) UnstageVoid(stage *StageStruct) {
	delete(stage.DummyDatas, dummydata)
	delete(stage.DummyDatas_mapString, dummydata.Name)
}

// commit dummydata to the back repo (if it is already staged)
func (dummydata *DummyData) Commit(stage *StageStruct) *DummyData {
	if _, ok := stage.DummyDatas[dummydata]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDummyData(dummydata)
		}
	}
	return dummydata
}

func (dummydata *DummyData) CommitVoid(stage *StageStruct) {
	dummydata.Commit(stage)
}

// Checkout dummydata to the back repo (if it is already staged)
func (dummydata *DummyData) Checkout(stage *StageStruct) *DummyData {
	if _, ok := stage.DummyDatas[dummydata]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDummyData(dummydata)
		}
	}
	return dummydata
}

// for satisfaction of GongStruct interface
func (dummydata *DummyData) GetName() (res string) {
	return dummydata.Name
}

// Stage puts element to the model stage
func (element *Element) Stage(stage *StageStruct) *Element {
	stage.Elements[element] = __member
	stage.Elements_mapString[element.Name] = element

	return element
}

// Unstage removes element off the model stage
func (element *Element) Unstage(stage *StageStruct) *Element {
	delete(stage.Elements, element)
	delete(stage.Elements_mapString, element.Name)
	return element
}

// UnstageVoid removes element off the model stage
func (element *Element) UnstageVoid(stage *StageStruct) {
	delete(stage.Elements, element)
	delete(stage.Elements_mapString, element.Name)
}

// commit element to the back repo (if it is already staged)
func (element *Element) Commit(stage *StageStruct) *Element {
	if _, ok := stage.Elements[element]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitElement(element)
		}
	}
	return element
}

func (element *Element) CommitVoid(stage *StageStruct) {
	element.Commit(stage)
}

// Checkout element to the back repo (if it is already staged)
func (element *Element) Checkout(stage *StageStruct) *Element {
	if _, ok := stage.Elements[element]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutElement(element)
		}
	}
	return element
}

// for satisfaction of GongStruct interface
func (element *Element) GetName() (res string) {
	return element.Name
}

// Stage puts markdowncontent to the model stage
func (markdowncontent *MarkdownContent) Stage(stage *StageStruct) *MarkdownContent {
	stage.MarkdownContents[markdowncontent] = __member
	stage.MarkdownContents_mapString[markdowncontent.Name] = markdowncontent

	return markdowncontent
}

// Unstage removes markdowncontent off the model stage
func (markdowncontent *MarkdownContent) Unstage(stage *StageStruct) *MarkdownContent {
	delete(stage.MarkdownContents, markdowncontent)
	delete(stage.MarkdownContents_mapString, markdowncontent.Name)
	return markdowncontent
}

// UnstageVoid removes markdowncontent off the model stage
func (markdowncontent *MarkdownContent) UnstageVoid(stage *StageStruct) {
	delete(stage.MarkdownContents, markdowncontent)
	delete(stage.MarkdownContents_mapString, markdowncontent.Name)
}

// commit markdowncontent to the back repo (if it is already staged)
func (markdowncontent *MarkdownContent) Commit(stage *StageStruct) *MarkdownContent {
	if _, ok := stage.MarkdownContents[markdowncontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMarkdownContent(markdowncontent)
		}
	}
	return markdowncontent
}

func (markdowncontent *MarkdownContent) CommitVoid(stage *StageStruct) {
	markdowncontent.Commit(stage)
}

// Checkout markdowncontent to the back repo (if it is already staged)
func (markdowncontent *MarkdownContent) Checkout(stage *StageStruct) *MarkdownContent {
	if _, ok := stage.MarkdownContents[markdowncontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMarkdownContent(markdowncontent)
		}
	}
	return markdowncontent
}

// for satisfaction of GongStruct interface
func (markdowncontent *MarkdownContent) GetName() (res string) {
	return markdowncontent.Name
}

// Stage puts row to the model stage
func (row *Row) Stage(stage *StageStruct) *Row {
	stage.Rows[row] = __member
	stage.Rows_mapString[row.Name] = row

	return row
}

// Unstage removes row off the model stage
func (row *Row) Unstage(stage *StageStruct) *Row {
	delete(stage.Rows, row)
	delete(stage.Rows_mapString, row.Name)
	return row
}

// UnstageVoid removes row off the model stage
func (row *Row) UnstageVoid(stage *StageStruct) {
	delete(stage.Rows, row)
	delete(stage.Rows_mapString, row.Name)
}

// commit row to the back repo (if it is already staged)
func (row *Row) Commit(stage *StageStruct) *Row {
	if _, ok := stage.Rows[row]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRow(row)
		}
	}
	return row
}

func (row *Row) CommitVoid(stage *StageStruct) {
	row.Commit(stage)
}

// Checkout row to the back repo (if it is already staged)
func (row *Row) Checkout(stage *StageStruct) *Row {
	if _, ok := stage.Rows[row]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRow(row)
		}
	}
	return row
}

// for satisfaction of GongStruct interface
func (row *Row) GetName() (res string) {
	return row.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAnotherDummyData(AnotherDummyData *AnotherDummyData)
	CreateORMCell(Cell *Cell)
	CreateORMDummyData(DummyData *DummyData)
	CreateORMElement(Element *Element)
	CreateORMMarkdownContent(MarkdownContent *MarkdownContent)
	CreateORMRow(Row *Row)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAnotherDummyData(AnotherDummyData *AnotherDummyData)
	DeleteORMCell(Cell *Cell)
	DeleteORMDummyData(DummyData *DummyData)
	DeleteORMElement(Element *Element)
	DeleteORMMarkdownContent(MarkdownContent *MarkdownContent)
	DeleteORMRow(Row *Row)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.AnotherDummyDatas = make(map[*AnotherDummyData]any)
	stage.AnotherDummyDatas_mapString = make(map[string]*AnotherDummyData)

	stage.Cells = make(map[*Cell]any)
	stage.Cells_mapString = make(map[string]*Cell)

	stage.DummyDatas = make(map[*DummyData]any)
	stage.DummyDatas_mapString = make(map[string]*DummyData)

	stage.Elements = make(map[*Element]any)
	stage.Elements_mapString = make(map[string]*Element)

	stage.MarkdownContents = make(map[*MarkdownContent]any)
	stage.MarkdownContents_mapString = make(map[string]*MarkdownContent)

	stage.Rows = make(map[*Row]any)
	stage.Rows_mapString = make(map[string]*Row)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.AnotherDummyDatas = nil
	stage.AnotherDummyDatas_mapString = nil

	stage.Cells = nil
	stage.Cells_mapString = nil

	stage.DummyDatas = nil
	stage.DummyDatas_mapString = nil

	stage.Elements = nil
	stage.Elements_mapString = nil

	stage.MarkdownContents = nil
	stage.MarkdownContents_mapString = nil

	stage.Rows = nil
	stage.Rows_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for anotherdummydata := range stage.AnotherDummyDatas {
		anotherdummydata.Unstage(stage)
	}

	for cell := range stage.Cells {
		cell.Unstage(stage)
	}

	for dummydata := range stage.DummyDatas {
		dummydata.Unstage(stage)
	}

	for element := range stage.Elements {
		element.Unstage(stage)
	}

	for markdowncontent := range stage.MarkdownContents {
		markdowncontent.Unstage(stage)
	}

	for row := range stage.Rows {
		row.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*AnotherDummyData]any:
		return any(&stage.AnotherDummyDatas).(*Type)
	case map[*Cell]any:
		return any(&stage.Cells).(*Type)
	case map[*DummyData]any:
		return any(&stage.DummyDatas).(*Type)
	case map[*Element]any:
		return any(&stage.Elements).(*Type)
	case map[*MarkdownContent]any:
		return any(&stage.MarkdownContents).(*Type)
	case map[*Row]any:
		return any(&stage.Rows).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*AnotherDummyData:
		return any(&stage.AnotherDummyDatas_mapString).(*Type)
	case map[string]*Cell:
		return any(&stage.Cells_mapString).(*Type)
	case map[string]*DummyData:
		return any(&stage.DummyDatas_mapString).(*Type)
	case map[string]*Element:
		return any(&stage.Elements_mapString).(*Type)
	case map[string]*MarkdownContent:
		return any(&stage.MarkdownContents_mapString).(*Type)
	case map[string]*Row:
		return any(&stage.Rows_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AnotherDummyData:
		return any(&stage.AnotherDummyDatas).(*map[*Type]any)
	case Cell:
		return any(&stage.Cells).(*map[*Type]any)
	case DummyData:
		return any(&stage.DummyDatas).(*map[*Type]any)
	case Element:
		return any(&stage.Elements).(*map[*Type]any)
	case MarkdownContent:
		return any(&stage.MarkdownContents).(*map[*Type]any)
	case Row:
		return any(&stage.Rows).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *AnotherDummyData:
		return any(&stage.AnotherDummyDatas).(*map[Type]any)
	case *Cell:
		return any(&stage.Cells).(*map[Type]any)
	case *DummyData:
		return any(&stage.DummyDatas).(*map[Type]any)
	case *Element:
		return any(&stage.Elements).(*map[Type]any)
	case *MarkdownContent:
		return any(&stage.MarkdownContents).(*map[Type]any)
	case *Row:
		return any(&stage.Rows).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AnotherDummyData:
		return any(&stage.AnotherDummyDatas_mapString).(*map[string]*Type)
	case Cell:
		return any(&stage.Cells_mapString).(*map[string]*Type)
	case DummyData:
		return any(&stage.DummyDatas_mapString).(*map[string]*Type)
	case Element:
		return any(&stage.Elements_mapString).(*map[string]*Type)
	case MarkdownContent:
		return any(&stage.MarkdownContents_mapString).(*map[string]*Type)
	case Row:
		return any(&stage.Rows_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case AnotherDummyData:
		return any(&AnotherDummyData{
			// Initialisation of associations
		}).(*Type)
	case Cell:
		return any(&Cell{
			// Initialisation of associations
		}).(*Type)
	case DummyData:
		return any(&DummyData{
			// Initialisation of associations
			// field is initialized with an instance of AnotherDummyData with the name of the field
			DummyPointerToGongStruct: &AnotherDummyData{Name: "DummyPointerToGongStruct"},
		}).(*Type)
	case Element:
		return any(&Element{
			// Initialisation of associations
			// field is initialized with an instance of Element with the name of the field
			SubElements: []*Element{{Name: "SubElements"}},
			// field is initialized with an instance of Row with the name of the field
			Rows: []*Row{{Name: "Rows"}},
		}).(*Type)
	case MarkdownContent:
		return any(&MarkdownContent{
			// Initialisation of associations
			// field is initialized with an instance of Element with the name of the field
			Root: &Element{Name: "Root"},
		}).(*Type)
	case Row:
		return any(&Row{
			// Initialisation of associations
			// field is initialized with an instance of Cell with the name of the field
			Cells: []*Cell{{Name: "Cells"}},
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of AnotherDummyData
	case AnotherDummyData:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Cell
	case Cell:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DummyData
	case DummyData:
		switch fieldname {
		// insertion point for per direct association field
		case "DummyPointerToGongStruct":
			res := make(map[*AnotherDummyData][]*DummyData)
			for dummydata := range stage.DummyDatas {
				if dummydata.DummyPointerToGongStruct != nil {
					anotherdummydata_ := dummydata.DummyPointerToGongStruct
					var dummydatas []*DummyData
					_, ok := res[anotherdummydata_]
					if ok {
						dummydatas = res[anotherdummydata_]
					} else {
						dummydatas = make([]*DummyData, 0)
					}
					dummydatas = append(dummydatas, dummydata)
					res[anotherdummydata_] = dummydatas
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Element
	case Element:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MarkdownContent
	case MarkdownContent:
		switch fieldname {
		// insertion point for per direct association field
		case "Root":
			res := make(map[*Element][]*MarkdownContent)
			for markdowncontent := range stage.MarkdownContents {
				if markdowncontent.Root != nil {
					element_ := markdowncontent.Root
					var markdowncontents []*MarkdownContent
					_, ok := res[element_]
					if ok {
						markdowncontents = res[element_]
					} else {
						markdowncontents = make([]*MarkdownContent, 0)
					}
					markdowncontents = append(markdowncontents, markdowncontent)
					res[element_] = markdowncontents
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Row
	case Row:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of AnotherDummyData
	case AnotherDummyData:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Cell
	case Cell:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DummyData
	case DummyData:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Element
	case Element:
		switch fieldname {
		// insertion point for per direct association field
		case "SubElements":
			res := make(map[*Element]*Element)
			for element := range stage.Elements {
				for _, element_ := range element.SubElements {
					res[element_] = element
				}
			}
			return any(res).(map[*End]*Start)
		case "Rows":
			res := make(map[*Row]*Element)
			for element := range stage.Elements {
				for _, row_ := range element.Rows {
					res[row_] = element
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of MarkdownContent
	case MarkdownContent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Row
	case Row:
		switch fieldname {
		// insertion point for per direct association field
		case "Cells":
			res := make(map[*Cell]*Row)
			for row := range stage.Rows {
				for _, cell_ := range row.Cells {
					res[cell_] = row
				}
			}
			return any(res).(map[*End]*Start)
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case AnotherDummyData:
		res = "AnotherDummyData"
	case Cell:
		res = "Cell"
	case DummyData:
		res = "DummyData"
	case Element:
		res = "Element"
	case MarkdownContent:
		res = "MarkdownContent"
	case Row:
		res = "Row"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *AnotherDummyData:
		res = "AnotherDummyData"
	case *Cell:
		res = "Cell"
	case *DummyData:
		res = "DummyData"
	case *Element:
		res = "Element"
	case *MarkdownContent:
		res = "MarkdownContent"
	case *Row:
		res = "Row"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case AnotherDummyData:
		res = []string{"Name"}
	case Cell:
		res = []string{"Name"}
	case DummyData:
		res = []string{"Name", "DummyString", "DummyInt", "DummyFloat", "DummyBool", "DummyEnumString", "DummyEnumInt", "DummyTime", "DummyDuration", "DummyPointerToGongStruct"}
	case Element:
		res = []string{"Name", "Content", "Type", "SubElements", "Rows"}
	case MarkdownContent:
		res = []string{"Name", "Content", "Root"}
	case Row:
		res = []string{"Name", "Cells"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case AnotherDummyData:
		var rf ReverseField
		_ = rf
	case Cell:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Row"
		rf.Fieldname = "Cells"
		res = append(res, rf)
	case DummyData:
		var rf ReverseField
		_ = rf
	case Element:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Element"
		rf.Fieldname = "SubElements"
		res = append(res, rf)
	case MarkdownContent:
		var rf ReverseField
		_ = rf
	case Row:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Element"
		rf.Fieldname = "Rows"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *AnotherDummyData:
		res = []string{"Name"}
	case *Cell:
		res = []string{"Name"}
	case *DummyData:
		res = []string{"Name", "DummyString", "DummyInt", "DummyFloat", "DummyBool", "DummyEnumString", "DummyEnumInt", "DummyTime", "DummyDuration", "DummyPointerToGongStruct"}
	case *Element:
		res = []string{"Name", "Content", "Type", "SubElements", "Rows"}
	case *MarkdownContent:
		res = []string{"Name", "Content", "Root"}
	case *Row:
		res = []string{"Name", "Cells"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *AnotherDummyData:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *Cell:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *DummyData:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DummyString":
			res = inferedInstance.DummyString
		case "DummyInt":
			res = fmt.Sprintf("%d", inferedInstance.DummyInt)
		case "DummyFloat":
			res = fmt.Sprintf("%f", inferedInstance.DummyFloat)
		case "DummyBool":
			res = fmt.Sprintf("%t", inferedInstance.DummyBool)
		case "DummyEnumString":
			enum := inferedInstance.DummyEnumString
			res = enum.ToCodeString()
		case "DummyEnumInt":
			enum := inferedInstance.DummyEnumInt
			res = enum.ToCodeString()
		case "DummyTime":
			res = inferedInstance.DummyTime.String()
		case "DummyDuration":
			if math.Abs(inferedInstance.DummyDuration.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.DummyDuration.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.DummyDuration.Hours()) % 24
				remainingMinutes := int(inferedInstance.DummyDuration.Minutes()) % 60
				remainingSeconds := int(inferedInstance.DummyDuration.Seconds()) % 60

				if inferedInstance.DummyDuration.Hours() < 0 {
					res = "- "
				}

				if months > 0 {
					if months > 1 {
						res = res + fmt.Sprintf("%d months", months)
					} else {
						res = res + fmt.Sprintf("%d month", months)
					}
				}
				if days > 0 {
					if months != 0 {
						res = res + ", "
					}
					if days > 1 {
						res = res + fmt.Sprintf("%d days", days)
					} else {
						res = res + fmt.Sprintf("%d day", days)
					}

				}
				if remainingHours != 0 || remainingMinutes != 0 || remainingSeconds != 0 {
					if days != 0 || (days == 0 && months != 0) {
						res = res + ", "
					}
					res = res + fmt.Sprintf("%d hours, %d minutes, %d seconds\n", remainingHours, remainingMinutes, remainingSeconds)
				}
			} else {
				res = fmt.Sprintf("%s\n", inferedInstance.DummyDuration.String())
			}
		case "DummyPointerToGongStruct":
			if inferedInstance.DummyPointerToGongStruct != nil {
				res = inferedInstance.DummyPointerToGongStruct.Name
			}
		}
	case *Element:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Content":
			res = inferedInstance.Content
		case "Type":
			enum := inferedInstance.Type
			res = enum.ToCodeString()
		case "SubElements":
			for idx, __instance__ := range inferedInstance.SubElements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Rows":
			for idx, __instance__ := range inferedInstance.Rows {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *MarkdownContent:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Content":
			res = inferedInstance.Content
		case "Root":
			if inferedInstance.Root != nil {
				res = inferedInstance.Root.Name
			}
		}
	case *Row:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Cells":
			for idx, __instance__ := range inferedInstance.Cells {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case AnotherDummyData:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case Cell:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case DummyData:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DummyString":
			res = inferedInstance.DummyString
		case "DummyInt":
			res = fmt.Sprintf("%d", inferedInstance.DummyInt)
		case "DummyFloat":
			res = fmt.Sprintf("%f", inferedInstance.DummyFloat)
		case "DummyBool":
			res = fmt.Sprintf("%t", inferedInstance.DummyBool)
		case "DummyEnumString":
			enum := inferedInstance.DummyEnumString
			res = enum.ToCodeString()
		case "DummyEnumInt":
			enum := inferedInstance.DummyEnumInt
			res = enum.ToCodeString()
		case "DummyTime":
			res = inferedInstance.DummyTime.String()
		case "DummyDuration":
			if math.Abs(inferedInstance.DummyDuration.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.DummyDuration.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.DummyDuration.Hours()) % 24
				remainingMinutes := int(inferedInstance.DummyDuration.Minutes()) % 60
				remainingSeconds := int(inferedInstance.DummyDuration.Seconds()) % 60

				if inferedInstance.DummyDuration.Hours() < 0 {
					res = "- "
				}

				if months > 0 {
					if months > 1 {
						res = res + fmt.Sprintf("%d months", months)
					} else {
						res = res + fmt.Sprintf("%d month", months)
					}
				}
				if days > 0 {
					if months != 0 {
						res = res + ", "
					}
					if days > 1 {
						res = res + fmt.Sprintf("%d days", days)
					} else {
						res = res + fmt.Sprintf("%d day", days)
					}

				}
				if remainingHours != 0 || remainingMinutes != 0 || remainingSeconds != 0 {
					if days != 0 || (days == 0 && months != 0) {
						res = res + ", "
					}
					res = res + fmt.Sprintf("%d hours, %d minutes, %d seconds\n", remainingHours, remainingMinutes, remainingSeconds)
				}
			} else {
				res = fmt.Sprintf("%s\n", inferedInstance.DummyDuration.String())
			}
		case "DummyPointerToGongStruct":
			if inferedInstance.DummyPointerToGongStruct != nil {
				res = inferedInstance.DummyPointerToGongStruct.Name
			}
		}
	case Element:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Content":
			res = inferedInstance.Content
		case "Type":
			enum := inferedInstance.Type
			res = enum.ToCodeString()
		case "SubElements":
			for idx, __instance__ := range inferedInstance.SubElements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Rows":
			for idx, __instance__ := range inferedInstance.Rows {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case MarkdownContent:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Content":
			res = inferedInstance.Content
		case "Root":
			if inferedInstance.Root != nil {
				res = inferedInstance.Root.Name
			}
		}
	case Row:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Cells":
			for idx, __instance__ := range inferedInstance.Cells {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
