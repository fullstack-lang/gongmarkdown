// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongmarkdown/go/models"
	"github.com/fullstack-lang/gongmarkdown/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __AnotherDummyData__dummysDeclaration__ models.AnotherDummyData
var __AnotherDummyData_time__dummyDeclaration time.Duration

var mutexAnotherDummyData sync.Mutex

// An AnotherDummyDataID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAnotherDummyData updateAnotherDummyData deleteAnotherDummyData
type AnotherDummyDataID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AnotherDummyDataInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAnotherDummyData updateAnotherDummyData
type AnotherDummyDataInput struct {
	// The AnotherDummyData to submit or modify
	// in: body
	AnotherDummyData *orm.AnotherDummyDataAPI
}

// GetAnotherDummyDatas
//
// swagger:route GET /anotherdummydatas anotherdummydatas getAnotherDummyDatas
//
// # Get all anotherdummydatas
//
// Responses:
// default: genericError
//
//	200: anotherdummydataDBResponse
func (controller *Controller) GetAnotherDummyDatas(c *gin.Context) {

	// source slice
	var anotherdummydataDBs []orm.AnotherDummyDataDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAnotherDummyDatas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnotherDummyData.GetDB()

	_, err := db.Find(&anotherdummydataDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	anotherdummydataAPIs := make([]orm.AnotherDummyDataAPI, 0)

	// for each anotherdummydata, update fields from the database nullable fields
	for idx := range anotherdummydataDBs {
		anotherdummydataDB := &anotherdummydataDBs[idx]
		_ = anotherdummydataDB
		var anotherdummydataAPI orm.AnotherDummyDataAPI

		// insertion point for updating fields
		anotherdummydataAPI.ID = anotherdummydataDB.ID
		anotherdummydataDB.CopyBasicFieldsToAnotherDummyData_WOP(&anotherdummydataAPI.AnotherDummyData_WOP)
		anotherdummydataAPI.AnotherDummyDataPointersEncoding = anotherdummydataDB.AnotherDummyDataPointersEncoding
		anotherdummydataAPIs = append(anotherdummydataAPIs, anotherdummydataAPI)
	}

	c.JSON(http.StatusOK, anotherdummydataAPIs)
}

// PostAnotherDummyData
//
// swagger:route POST /anotherdummydatas anotherdummydatas postAnotherDummyData
//
// Creates a anotherdummydata
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAnotherDummyData(c *gin.Context) {

	mutexAnotherDummyData.Lock()
	defer mutexAnotherDummyData.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAnotherDummyDatas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnotherDummyData.GetDB()

	// Validate input
	var input orm.AnotherDummyDataAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create anotherdummydata
	anotherdummydataDB := orm.AnotherDummyDataDB{}
	anotherdummydataDB.AnotherDummyDataPointersEncoding = input.AnotherDummyDataPointersEncoding
	anotherdummydataDB.CopyBasicFieldsFromAnotherDummyData_WOP(&input.AnotherDummyData_WOP)

	_, err = db.Create(&anotherdummydataDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAnotherDummyData.CheckoutPhaseOneInstance(&anotherdummydataDB)
	anotherdummydata := backRepo.BackRepoAnotherDummyData.Map_AnotherDummyDataDBID_AnotherDummyDataPtr[anotherdummydataDB.ID]

	if anotherdummydata != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), anotherdummydata)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, anotherdummydataDB)
}

// GetAnotherDummyData
//
// swagger:route GET /anotherdummydatas/{ID} anotherdummydatas getAnotherDummyData
//
// Gets the details for a anotherdummydata.
//
// Responses:
// default: genericError
//
//	200: anotherdummydataDBResponse
func (controller *Controller) GetAnotherDummyData(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAnotherDummyData", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnotherDummyData.GetDB()

	// Get anotherdummydataDB in DB
	var anotherdummydataDB orm.AnotherDummyDataDB
	if _, err := db.First(&anotherdummydataDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var anotherdummydataAPI orm.AnotherDummyDataAPI
	anotherdummydataAPI.ID = anotherdummydataDB.ID
	anotherdummydataAPI.AnotherDummyDataPointersEncoding = anotherdummydataDB.AnotherDummyDataPointersEncoding
	anotherdummydataDB.CopyBasicFieldsToAnotherDummyData_WOP(&anotherdummydataAPI.AnotherDummyData_WOP)

	c.JSON(http.StatusOK, anotherdummydataAPI)
}

// UpdateAnotherDummyData
//
// swagger:route PATCH /anotherdummydatas/{ID} anotherdummydatas updateAnotherDummyData
//
// # Update a anotherdummydata
//
// Responses:
// default: genericError
//
//	200: anotherdummydataDBResponse
func (controller *Controller) UpdateAnotherDummyData(c *gin.Context) {

	mutexAnotherDummyData.Lock()
	defer mutexAnotherDummyData.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAnotherDummyData", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnotherDummyData.GetDB()

	// Validate input
	var input orm.AnotherDummyDataAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var anotherdummydataDB orm.AnotherDummyDataDB

	// fetch the anotherdummydata
	_, err := db.First(&anotherdummydataDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	anotherdummydataDB.CopyBasicFieldsFromAnotherDummyData_WOP(&input.AnotherDummyData_WOP)
	anotherdummydataDB.AnotherDummyDataPointersEncoding = input.AnotherDummyDataPointersEncoding

	db, _ = db.Model(&anotherdummydataDB)
	_, err = db.Updates(&anotherdummydataDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	anotherdummydataNew := new(models.AnotherDummyData)
	anotherdummydataDB.CopyBasicFieldsToAnotherDummyData(anotherdummydataNew)

	// redeem pointers
	anotherdummydataDB.DecodePointers(backRepo, anotherdummydataNew)

	// get stage instance from DB instance, and call callback function
	anotherdummydataOld := backRepo.BackRepoAnotherDummyData.Map_AnotherDummyDataDBID_AnotherDummyDataPtr[anotherdummydataDB.ID]
	if anotherdummydataOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), anotherdummydataOld, anotherdummydataNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the anotherdummydataDB
	c.JSON(http.StatusOK, anotherdummydataDB)
}

// DeleteAnotherDummyData
//
// swagger:route DELETE /anotherdummydatas/{ID} anotherdummydatas deleteAnotherDummyData
//
// # Delete a anotherdummydata
//
// default: genericError
//
//	200: anotherdummydataDBResponse
func (controller *Controller) DeleteAnotherDummyData(c *gin.Context) {

	mutexAnotherDummyData.Lock()
	defer mutexAnotherDummyData.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAnotherDummyData", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnotherDummyData.GetDB()

	// Get model if exist
	var anotherdummydataDB orm.AnotherDummyDataDB
	if _, err := db.First(&anotherdummydataDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&anotherdummydataDB)

	// get an instance (not staged) from DB instance, and call callback function
	anotherdummydataDeleted := new(models.AnotherDummyData)
	anotherdummydataDB.CopyBasicFieldsToAnotherDummyData(anotherdummydataDeleted)

	// get stage instance from DB instance, and call callback function
	anotherdummydataStaged := backRepo.BackRepoAnotherDummyData.Map_AnotherDummyDataDBID_AnotherDummyDataPtr[anotherdummydataDB.ID]
	if anotherdummydataStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), anotherdummydataStaged, anotherdummydataDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
