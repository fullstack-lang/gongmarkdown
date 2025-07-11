// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gong/lib/doc2/go/db"
	"github.com/fullstack-lang/gong/lib/doc2/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_GongEnumShape_sql sql.NullBool
var dummy_GongEnumShape_time time.Duration
var dummy_GongEnumShape_sort sort.Float64Slice

// GongEnumShapeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model gongenumshapeAPI
type GongEnumShapeAPI struct {
	gorm.Model

	models.GongEnumShape_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	GongEnumShapePointersEncoding GongEnumShapePointersEncoding
}

// GongEnumShapePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GongEnumShapePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field GongEnumValueShapes is a slice of pointers to another Struct (optional or 0..1)
	GongEnumValueShapes IntSlice `gorm:"type:TEXT"`
}

// GongEnumShapeDB describes a gongenumshape in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model gongenumshapeDB
type GongEnumShapeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field gongenumshapeDB.Name
	Name_Data sql.NullString

	// Declation for basic field gongenumshapeDB.X
	X_Data sql.NullFloat64

	// Declation for basic field gongenumshapeDB.Y
	Y_Data sql.NullFloat64

	// Declation for basic field gongenumshapeDB.Width
	Width_Data sql.NullFloat64

	// Declation for basic field gongenumshapeDB.Height
	Height_Data sql.NullFloat64

	// Declation for basic field gongenumshapeDB.IsExpanded
	// provide the sql storage for the boolan
	IsExpanded_Data sql.NullBool

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	GongEnumShapePointersEncoding
}

// GongEnumShapeDBs arrays gongenumshapeDBs
// swagger:response gongenumshapeDBsResponse
type GongEnumShapeDBs []GongEnumShapeDB

// GongEnumShapeDBResponse provides response
// swagger:response gongenumshapeDBResponse
type GongEnumShapeDBResponse struct {
	GongEnumShapeDB
}

// GongEnumShapeWOP is a GongEnumShape without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GongEnumShapeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	X float64 `xlsx:"2"`

	Y float64 `xlsx:"3"`

	IdentifierMeta any `xlsx:"4"`

	Width float64 `xlsx:"5"`

	Height float64 `xlsx:"6"`

	IsExpanded bool `xlsx:"7"`
	// insertion for WOP pointer fields
}

var GongEnumShape_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"X",
	"Y",
	"IdentifierMeta",
	"Width",
	"Height",
	"IsExpanded",
}

type BackRepoGongEnumShapeStruct struct {
	// stores GongEnumShapeDB according to their gorm ID
	Map_GongEnumShapeDBID_GongEnumShapeDB map[uint]*GongEnumShapeDB

	// stores GongEnumShapeDB ID according to GongEnumShape address
	Map_GongEnumShapePtr_GongEnumShapeDBID map[*models.GongEnumShape]uint

	// stores GongEnumShape according to their gorm ID
	Map_GongEnumShapeDBID_GongEnumShapePtr map[uint]*models.GongEnumShape

	db db.DBInterface

	stage *models.Stage
}

func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) GetStage() (stage *models.Stage) {
	stage = backRepoGongEnumShape.stage
	return
}

func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) GetDB() db.DBInterface {
	return backRepoGongEnumShape.db
}

// GetGongEnumShapeDBFromGongEnumShapePtr is a handy function to access the back repo instance from the stage instance
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) GetGongEnumShapeDBFromGongEnumShapePtr(gongenumshape *models.GongEnumShape) (gongenumshapeDB *GongEnumShapeDB) {
	id := backRepoGongEnumShape.Map_GongEnumShapePtr_GongEnumShapeDBID[gongenumshape]
	gongenumshapeDB = backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB[id]
	return
}

// BackRepoGongEnumShape.CommitPhaseOne commits all staged instances of GongEnumShape to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var gongenumshapes []*models.GongEnumShape
	for gongenumshape := range stage.GongEnumShapes {
		gongenumshapes = append(gongenumshapes, gongenumshape)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(gongenumshapes, func(i, j int) bool {
		return stage.GongEnumShapeMap_Staged_Order[gongenumshapes[i]] < stage.GongEnumShapeMap_Staged_Order[gongenumshapes[j]]
	})

	for _, gongenumshape := range gongenumshapes {
		backRepoGongEnumShape.CommitPhaseOneInstance(gongenumshape)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, gongenumshape := range backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr {
		if _, ok := stage.GongEnumShapes[gongenumshape]; !ok {
			backRepoGongEnumShape.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGongEnumShape.CommitDeleteInstance commits deletion of GongEnumShape to the BackRepo
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) CommitDeleteInstance(id uint) (Error error) {

	gongenumshape := backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[id]

	// gongenumshape is not staged anymore, remove gongenumshapeDB
	gongenumshapeDB := backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB[id]
	db, _ := backRepoGongEnumShape.db.Unscoped()
	_, err := db.Delete(gongenumshapeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoGongEnumShape.Map_GongEnumShapePtr_GongEnumShapeDBID, gongenumshape)
	delete(backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr, id)
	delete(backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB, id)

	return
}

// BackRepoGongEnumShape.CommitPhaseOneInstance commits gongenumshape staged instances of GongEnumShape to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) CommitPhaseOneInstance(gongenumshape *models.GongEnumShape) (Error error) {

	// check if the gongenumshape is not commited yet
	if _, ok := backRepoGongEnumShape.Map_GongEnumShapePtr_GongEnumShapeDBID[gongenumshape]; ok {
		return
	}

	// initiate gongenumshape
	var gongenumshapeDB GongEnumShapeDB
	gongenumshapeDB.CopyBasicFieldsFromGongEnumShape(gongenumshape)

	_, err := backRepoGongEnumShape.db.Create(&gongenumshapeDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoGongEnumShape.Map_GongEnumShapePtr_GongEnumShapeDBID[gongenumshape] = gongenumshapeDB.ID
	backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[gongenumshapeDB.ID] = gongenumshape
	backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB[gongenumshapeDB.ID] = &gongenumshapeDB

	return
}

// BackRepoGongEnumShape.CommitPhaseTwo commits all staged instances of GongEnumShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, gongenumshape := range backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr {
		backRepoGongEnumShape.CommitPhaseTwoInstance(backRepo, idx, gongenumshape)
	}

	return
}

// BackRepoGongEnumShape.CommitPhaseTwoInstance commits {{structname }} of models.GongEnumShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, gongenumshape *models.GongEnumShape) (Error error) {

	// fetch matching gongenumshapeDB
	if gongenumshapeDB, ok := backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB[idx]; ok {

		gongenumshapeDB.CopyBasicFieldsFromGongEnumShape(gongenumshape)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		gongenumshapeDB.GongEnumShapePointersEncoding.GongEnumValueShapes = make([]int, 0)
		// 2. encode
		for _, gongenumvalueshapeAssocEnd := range gongenumshape.GongEnumValueShapes {
			gongenumvalueshapeAssocEnd_DB :=
				backRepo.BackRepoGongEnumValueShape.GetGongEnumValueShapeDBFromGongEnumValueShapePtr(gongenumvalueshapeAssocEnd)
			
			// the stage might be inconsistant, meaning that the gongenumvalueshapeAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if gongenumvalueshapeAssocEnd_DB == nil {
				continue
			}
			
			gongenumshapeDB.GongEnumShapePointersEncoding.GongEnumValueShapes =
				append(gongenumshapeDB.GongEnumShapePointersEncoding.GongEnumValueShapes, int(gongenumvalueshapeAssocEnd_DB.ID))
		}

		_, err := backRepoGongEnumShape.db.Save(gongenumshapeDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown GongEnumShape intance %s", gongenumshape.Name))
		return err
	}

	return
}

// BackRepoGongEnumShape.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) CheckoutPhaseOne() (Error error) {

	gongenumshapeDBArray := make([]GongEnumShapeDB, 0)
	_, err := backRepoGongEnumShape.db.Find(&gongenumshapeDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	gongenumshapeInstancesToBeRemovedFromTheStage := make(map[*models.GongEnumShape]any)
	for key, value := range backRepoGongEnumShape.stage.GongEnumShapes {
		gongenumshapeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, gongenumshapeDB := range gongenumshapeDBArray {
		backRepoGongEnumShape.CheckoutPhaseOneInstance(&gongenumshapeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		gongenumshape, ok := backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[gongenumshapeDB.ID]
		if ok {
			delete(gongenumshapeInstancesToBeRemovedFromTheStage, gongenumshape)
		}
	}

	// remove from stage and back repo's 3 maps all gongenumshapes that are not in the checkout
	for gongenumshape := range gongenumshapeInstancesToBeRemovedFromTheStage {
		gongenumshape.Unstage(backRepoGongEnumShape.GetStage())

		// remove instance from the back repo 3 maps
		gongenumshapeID := backRepoGongEnumShape.Map_GongEnumShapePtr_GongEnumShapeDBID[gongenumshape]
		delete(backRepoGongEnumShape.Map_GongEnumShapePtr_GongEnumShapeDBID, gongenumshape)
		delete(backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB, gongenumshapeID)
		delete(backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr, gongenumshapeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a gongenumshapeDB that has been found in the DB, updates the backRepo and stages the
// models version of the gongenumshapeDB
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) CheckoutPhaseOneInstance(gongenumshapeDB *GongEnumShapeDB) (Error error) {

	gongenumshape, ok := backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[gongenumshapeDB.ID]
	if !ok {
		gongenumshape = new(models.GongEnumShape)

		backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[gongenumshapeDB.ID] = gongenumshape
		backRepoGongEnumShape.Map_GongEnumShapePtr_GongEnumShapeDBID[gongenumshape] = gongenumshapeDB.ID

		// append model store with the new element
		gongenumshape.Name = gongenumshapeDB.Name_Data.String
		gongenumshape.Stage(backRepoGongEnumShape.GetStage())
	}
	gongenumshapeDB.CopyBasicFieldsToGongEnumShape(gongenumshape)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	gongenumshape.Stage(backRepoGongEnumShape.GetStage())

	// preserve pointer to gongenumshapeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GongEnumShapeDBID_GongEnumShapeDB)[gongenumshapeDB hold variable pointers
	gongenumshapeDB_Data := *gongenumshapeDB
	preservedPtrToGongEnumShape := &gongenumshapeDB_Data
	backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB[gongenumshapeDB.ID] = preservedPtrToGongEnumShape

	return
}

// BackRepoGongEnumShape.CheckoutPhaseTwo Checkouts all staged instances of GongEnumShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, gongenumshapeDB := range backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB {
		backRepoGongEnumShape.CheckoutPhaseTwoInstance(backRepo, gongenumshapeDB)
	}
	return
}

// BackRepoGongEnumShape.CheckoutPhaseTwoInstance Checkouts staged instances of GongEnumShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, gongenumshapeDB *GongEnumShapeDB) (Error error) {

	gongenumshape := backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr[gongenumshapeDB.ID]

	gongenumshapeDB.DecodePointers(backRepo, gongenumshape)

	return
}

func (gongenumshapeDB *GongEnumShapeDB) DecodePointers(backRepo *BackRepoStruct, gongenumshape *models.GongEnumShape) {

	// insertion point for checkout of pointer encoding
	// This loop redeem gongenumshape.GongEnumValueShapes in the stage from the encode in the back repo
	// It parses all GongEnumValueShapeDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	gongenumshape.GongEnumValueShapes = gongenumshape.GongEnumValueShapes[:0]
	for _, _GongEnumValueShapeid := range gongenumshapeDB.GongEnumShapePointersEncoding.GongEnumValueShapes {
		gongenumshape.GongEnumValueShapes = append(gongenumshape.GongEnumValueShapes, backRepo.BackRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapePtr[uint(_GongEnumValueShapeid)])
	}

	return
}

// CommitGongEnumShape allows commit of a single gongenumshape (if already staged)
func (backRepo *BackRepoStruct) CommitGongEnumShape(gongenumshape *models.GongEnumShape) {
	backRepo.BackRepoGongEnumShape.CommitPhaseOneInstance(gongenumshape)
	if id, ok := backRepo.BackRepoGongEnumShape.Map_GongEnumShapePtr_GongEnumShapeDBID[gongenumshape]; ok {
		backRepo.BackRepoGongEnumShape.CommitPhaseTwoInstance(backRepo, id, gongenumshape)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitGongEnumShape allows checkout of a single gongenumshape (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGongEnumShape(gongenumshape *models.GongEnumShape) {
	// check if the gongenumshape is staged
	if _, ok := backRepo.BackRepoGongEnumShape.Map_GongEnumShapePtr_GongEnumShapeDBID[gongenumshape]; ok {

		if id, ok := backRepo.BackRepoGongEnumShape.Map_GongEnumShapePtr_GongEnumShapeDBID[gongenumshape]; ok {
			var gongenumshapeDB GongEnumShapeDB
			gongenumshapeDB.ID = id

			if _, err := backRepo.BackRepoGongEnumShape.db.First(&gongenumshapeDB, id); err != nil {
				log.Fatalln("CheckoutGongEnumShape : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGongEnumShape.CheckoutPhaseOneInstance(&gongenumshapeDB)
			backRepo.BackRepoGongEnumShape.CheckoutPhaseTwoInstance(backRepo, &gongenumshapeDB)
		}
	}
}

// CopyBasicFieldsFromGongEnumShape
func (gongenumshapeDB *GongEnumShapeDB) CopyBasicFieldsFromGongEnumShape(gongenumshape *models.GongEnumShape) {
	// insertion point for fields commit

	gongenumshapeDB.Name_Data.String = gongenumshape.Name
	gongenumshapeDB.Name_Data.Valid = true

	gongenumshapeDB.X_Data.Float64 = gongenumshape.X
	gongenumshapeDB.X_Data.Valid = true

	gongenumshapeDB.Y_Data.Float64 = gongenumshape.Y
	gongenumshapeDB.Y_Data.Valid = true

	gongenumshapeDB.Width_Data.Float64 = gongenumshape.Width
	gongenumshapeDB.Width_Data.Valid = true

	gongenumshapeDB.Height_Data.Float64 = gongenumshape.Height
	gongenumshapeDB.Height_Data.Valid = true

	gongenumshapeDB.IsExpanded_Data.Bool = gongenumshape.IsExpanded
	gongenumshapeDB.IsExpanded_Data.Valid = true
}

// CopyBasicFieldsFromGongEnumShape_WOP
func (gongenumshapeDB *GongEnumShapeDB) CopyBasicFieldsFromGongEnumShape_WOP(gongenumshape *models.GongEnumShape_WOP) {
	// insertion point for fields commit

	gongenumshapeDB.Name_Data.String = gongenumshape.Name
	gongenumshapeDB.Name_Data.Valid = true

	gongenumshapeDB.X_Data.Float64 = gongenumshape.X
	gongenumshapeDB.X_Data.Valid = true

	gongenumshapeDB.Y_Data.Float64 = gongenumshape.Y
	gongenumshapeDB.Y_Data.Valid = true

	gongenumshapeDB.Width_Data.Float64 = gongenumshape.Width
	gongenumshapeDB.Width_Data.Valid = true

	gongenumshapeDB.Height_Data.Float64 = gongenumshape.Height
	gongenumshapeDB.Height_Data.Valid = true

	gongenumshapeDB.IsExpanded_Data.Bool = gongenumshape.IsExpanded
	gongenumshapeDB.IsExpanded_Data.Valid = true
}

// CopyBasicFieldsFromGongEnumShapeWOP
func (gongenumshapeDB *GongEnumShapeDB) CopyBasicFieldsFromGongEnumShapeWOP(gongenumshape *GongEnumShapeWOP) {
	// insertion point for fields commit

	gongenumshapeDB.Name_Data.String = gongenumshape.Name
	gongenumshapeDB.Name_Data.Valid = true

	gongenumshapeDB.X_Data.Float64 = gongenumshape.X
	gongenumshapeDB.X_Data.Valid = true

	gongenumshapeDB.Y_Data.Float64 = gongenumshape.Y
	gongenumshapeDB.Y_Data.Valid = true

	gongenumshapeDB.Width_Data.Float64 = gongenumshape.Width
	gongenumshapeDB.Width_Data.Valid = true

	gongenumshapeDB.Height_Data.Float64 = gongenumshape.Height
	gongenumshapeDB.Height_Data.Valid = true

	gongenumshapeDB.IsExpanded_Data.Bool = gongenumshape.IsExpanded
	gongenumshapeDB.IsExpanded_Data.Valid = true
}

// CopyBasicFieldsToGongEnumShape
func (gongenumshapeDB *GongEnumShapeDB) CopyBasicFieldsToGongEnumShape(gongenumshape *models.GongEnumShape) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongenumshape.Name = gongenumshapeDB.Name_Data.String
	gongenumshape.X = gongenumshapeDB.X_Data.Float64
	gongenumshape.Y = gongenumshapeDB.Y_Data.Float64
	gongenumshape.Width = gongenumshapeDB.Width_Data.Float64
	gongenumshape.Height = gongenumshapeDB.Height_Data.Float64
	gongenumshape.IsExpanded = gongenumshapeDB.IsExpanded_Data.Bool
}

// CopyBasicFieldsToGongEnumShape_WOP
func (gongenumshapeDB *GongEnumShapeDB) CopyBasicFieldsToGongEnumShape_WOP(gongenumshape *models.GongEnumShape_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongenumshape.Name = gongenumshapeDB.Name_Data.String
	gongenumshape.X = gongenumshapeDB.X_Data.Float64
	gongenumshape.Y = gongenumshapeDB.Y_Data.Float64
	gongenumshape.Width = gongenumshapeDB.Width_Data.Float64
	gongenumshape.Height = gongenumshapeDB.Height_Data.Float64
	gongenumshape.IsExpanded = gongenumshapeDB.IsExpanded_Data.Bool
}

// CopyBasicFieldsToGongEnumShapeWOP
func (gongenumshapeDB *GongEnumShapeDB) CopyBasicFieldsToGongEnumShapeWOP(gongenumshape *GongEnumShapeWOP) {
	gongenumshape.ID = int(gongenumshapeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	gongenumshape.Name = gongenumshapeDB.Name_Data.String
	gongenumshape.X = gongenumshapeDB.X_Data.Float64
	gongenumshape.Y = gongenumshapeDB.Y_Data.Float64
	gongenumshape.Width = gongenumshapeDB.Width_Data.Float64
	gongenumshape.Height = gongenumshapeDB.Height_Data.Float64
	gongenumshape.IsExpanded = gongenumshapeDB.IsExpanded_Data.Bool
}

// Backup generates a json file from a slice of all GongEnumShapeDB instances in the backrepo
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GongEnumShapeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongEnumShapeDB, 0)
	for _, gongenumshapeDB := range backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB {
		forBackup = append(forBackup, gongenumshapeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json GongEnumShape ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json GongEnumShape file", err.Error())
	}
}

// Backup generates a json file from a slice of all GongEnumShapeDB instances in the backrepo
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongEnumShapeDB, 0)
	for _, gongenumshapeDB := range backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB {
		forBackup = append(forBackup, gongenumshapeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("GongEnumShape")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&GongEnumShape_Fields, -1)
	for _, gongenumshapeDB := range forBackup {

		var gongenumshapeWOP GongEnumShapeWOP
		gongenumshapeDB.CopyBasicFieldsToGongEnumShapeWOP(&gongenumshapeWOP)

		row := sh.AddRow()
		row.WriteStruct(&gongenumshapeWOP, -1)
	}
}

// RestoreXL from the "GongEnumShape" sheet all GongEnumShapeDB instances
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGongEnumShapeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["GongEnumShape"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGongEnumShape.rowVisitorGongEnumShape)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) rowVisitorGongEnumShape(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var gongenumshapeWOP GongEnumShapeWOP
		row.ReadStruct(&gongenumshapeWOP)

		// add the unmarshalled struct to the stage
		gongenumshapeDB := new(GongEnumShapeDB)
		gongenumshapeDB.CopyBasicFieldsFromGongEnumShapeWOP(&gongenumshapeWOP)

		gongenumshapeDB_ID_atBackupTime := gongenumshapeDB.ID
		gongenumshapeDB.ID = 0
		_, err := backRepoGongEnumShape.db.Create(gongenumshapeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB[gongenumshapeDB.ID] = gongenumshapeDB
		BackRepoGongEnumShapeid_atBckpTime_newID[gongenumshapeDB_ID_atBackupTime] = gongenumshapeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GongEnumShapeDB.json" in dirPath that stores an array
// of GongEnumShapeDB and stores it in the database
// the map BackRepoGongEnumShapeid_atBckpTime_newID is updated accordingly
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGongEnumShapeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GongEnumShapeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json GongEnumShape file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GongEnumShapeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GongEnumShapeDBID_GongEnumShapeDB
	for _, gongenumshapeDB := range forRestore {

		gongenumshapeDB_ID_atBackupTime := gongenumshapeDB.ID
		gongenumshapeDB.ID = 0
		_, err := backRepoGongEnumShape.db.Create(gongenumshapeDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB[gongenumshapeDB.ID] = gongenumshapeDB
		BackRepoGongEnumShapeid_atBckpTime_newID[gongenumshapeDB_ID_atBackupTime] = gongenumshapeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json GongEnumShape file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<GongEnumShape>id_atBckpTime_newID
// to compute new index
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) RestorePhaseTwo() {

	for _, gongenumshapeDB := range backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB {

		// next line of code is to avert unused variable compilation error
		_ = gongenumshapeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoGongEnumShape.db.Model(gongenumshapeDB)
		_, err := db.Updates(*gongenumshapeDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoGongEnumShape.ResetReversePointers commits all staged instances of GongEnumShape to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, gongenumshape := range backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapePtr {
		backRepoGongEnumShape.ResetReversePointersInstance(backRepo, idx, gongenumshape)
	}

	return
}

func (backRepoGongEnumShape *BackRepoGongEnumShapeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, gongenumshape *models.GongEnumShape) (Error error) {

	// fetch matching gongenumshapeDB
	if gongenumshapeDB, ok := backRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB[idx]; ok {
		_ = gongenumshapeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGongEnumShapeid_atBckpTime_newID map[uint]uint
