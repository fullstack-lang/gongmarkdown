// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __NoteShape__dummysDeclaration__ models.NoteShape
var __NoteShape_time__dummyDeclaration time.Duration

// An NoteShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNoteShape updateNoteShape deleteNoteShape
type NoteShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NoteShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNoteShape updateNoteShape
type NoteShapeInput struct {
	// The NoteShape to submit or modify
	// in: body
	NoteShape *orm.NoteShapeAPI
}

// GetNoteShapes
//
// swagger:route GET /noteshapes noteshapes getNoteShapes
//
// # Get all noteshapes
//
// Responses:
// default: genericError
//
//	200: noteshapeDBResponse
func GetNoteShapes(c *gin.Context) {
	db := orm.BackRepo.BackRepoNoteShape.GetDB()

	// source slice
	var noteshapeDBs []orm.NoteShapeDB
	query := db.Find(&noteshapeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	noteshapeAPIs := make([]orm.NoteShapeAPI, 0)

	// for each noteshape, update fields from the database nullable fields
	for idx := range noteshapeDBs {
		noteshapeDB := &noteshapeDBs[idx]
		_ = noteshapeDB
		var noteshapeAPI orm.NoteShapeAPI

		// insertion point for updating fields
		noteshapeAPI.ID = noteshapeDB.ID
		noteshapeDB.CopyBasicFieldsToNoteShape(&noteshapeAPI.NoteShape)
		noteshapeAPI.NoteShapePointersEnconding = noteshapeDB.NoteShapePointersEnconding
		noteshapeAPIs = append(noteshapeAPIs, noteshapeAPI)
	}

	c.JSON(http.StatusOK, noteshapeAPIs)
}

// PostNoteShape
//
// swagger:route POST /noteshapes noteshapes postNoteShape
//
// Creates a noteshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func PostNoteShape(c *gin.Context) {
	db := orm.BackRepo.BackRepoNoteShape.GetDB()

	// Validate input
	var input orm.NoteShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create noteshape
	noteshapeDB := orm.NoteShapeDB{}
	noteshapeDB.NoteShapePointersEnconding = input.NoteShapePointersEnconding
	noteshapeDB.CopyBasicFieldsFromNoteShape(&input.NoteShape)

	query := db.Create(&noteshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	orm.BackRepo.BackRepoNoteShape.CheckoutPhaseOneInstance(&noteshapeDB)
	noteshape := (*orm.BackRepo.BackRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr)[noteshapeDB.ID]

	if noteshape != nil {
		models.AfterCreateFromFront(&models.Stage, noteshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, noteshapeDB)
}

// GetNoteShape
//
// swagger:route GET /noteshapes/{ID} noteshapes getNoteShape
//
// Gets the details for a noteshape.
//
// Responses:
// default: genericError
//
//	200: noteshapeDBResponse
func GetNoteShape(c *gin.Context) {
	db := orm.BackRepo.BackRepoNoteShape.GetDB()

	// Get noteshapeDB in DB
	var noteshapeDB orm.NoteShapeDB
	if err := db.First(&noteshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var noteshapeAPI orm.NoteShapeAPI
	noteshapeAPI.ID = noteshapeDB.ID
	noteshapeAPI.NoteShapePointersEnconding = noteshapeDB.NoteShapePointersEnconding
	noteshapeDB.CopyBasicFieldsToNoteShape(&noteshapeAPI.NoteShape)

	c.JSON(http.StatusOK, noteshapeAPI)
}

// UpdateNoteShape
//
// swagger:route PATCH /noteshapes/{ID} noteshapes updateNoteShape
//
// # Update a noteshape
//
// Responses:
// default: genericError
//
//	200: noteshapeDBResponse
func UpdateNoteShape(c *gin.Context) {
	db := orm.BackRepo.BackRepoNoteShape.GetDB()

	// Get model if exist
	var noteshapeDB orm.NoteShapeDB

	// fetch the noteshape
	query := db.First(&noteshapeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.NoteShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	noteshapeDB.CopyBasicFieldsFromNoteShape(&input.NoteShape)
	noteshapeDB.NoteShapePointersEnconding = input.NoteShapePointersEnconding

	query = db.Model(&noteshapeDB).Updates(noteshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	noteshapeNew := new(models.NoteShape)
	noteshapeDB.CopyBasicFieldsToNoteShape(noteshapeNew)

	// get stage instance from DB instance, and call callback function
	noteshapeOld := (*orm.BackRepo.BackRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr)[noteshapeDB.ID]
	if noteshapeOld != nil {
		models.AfterUpdateFromFront(&models.Stage, noteshapeOld, noteshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the noteshapeDB
	c.JSON(http.StatusOK, noteshapeDB)
}

// DeleteNoteShape
//
// swagger:route DELETE /noteshapes/{ID} noteshapes deleteNoteShape
//
// # Delete a noteshape
//
// default: genericError
//
//	200: noteshapeDBResponse
func DeleteNoteShape(c *gin.Context) {
	db := orm.BackRepo.BackRepoNoteShape.GetDB()

	// Get model if exist
	var noteshapeDB orm.NoteShapeDB
	if err := db.First(&noteshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&noteshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	noteshapeDeleted := new(models.NoteShape)
	noteshapeDB.CopyBasicFieldsToNoteShape(noteshapeDeleted)

	// get stage instance from DB instance, and call callback function
	noteshapeStaged := (*orm.BackRepo.BackRepoNoteShape.Map_NoteShapeDBID_NoteShapePtr)[noteshapeDB.ID]
	if noteshapeStaged != nil {
		models.AfterDeleteFromFront(&models.Stage, noteshapeStaged, noteshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}