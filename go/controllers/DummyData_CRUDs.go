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
var __DummyData__dummysDeclaration__ models.DummyData
var __DummyData_time__dummyDeclaration time.Duration

var mutexDummyData sync.Mutex

// An DummyDataID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDummyData updateDummyData deleteDummyData
type DummyDataID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DummyDataInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDummyData updateDummyData
type DummyDataInput struct {
	// The DummyData to submit or modify
	// in: body
	DummyData *orm.DummyDataAPI
}

// GetDummyDatas
//
// swagger:route GET /dummydatas dummydatas getDummyDatas
//
// # Get all dummydatas
//
// Responses:
// default: genericError
//
//	200: dummydataDBResponse
func (controller *Controller) GetDummyDatas(c *gin.Context) {

	// source slice
	var dummydataDBs []orm.DummyDataDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDummyDatas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDummyData.GetDB()

	_, err := db.Find(&dummydataDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	dummydataAPIs := make([]orm.DummyDataAPI, 0)

	// for each dummydata, update fields from the database nullable fields
	for idx := range dummydataDBs {
		dummydataDB := &dummydataDBs[idx]
		_ = dummydataDB
		var dummydataAPI orm.DummyDataAPI

		// insertion point for updating fields
		dummydataAPI.ID = dummydataDB.ID
		dummydataDB.CopyBasicFieldsToDummyData_WOP(&dummydataAPI.DummyData_WOP)
		dummydataAPI.DummyDataPointersEncoding = dummydataDB.DummyDataPointersEncoding
		dummydataAPIs = append(dummydataAPIs, dummydataAPI)
	}

	c.JSON(http.StatusOK, dummydataAPIs)
}

// PostDummyData
//
// swagger:route POST /dummydatas dummydatas postDummyData
//
// Creates a dummydata
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDummyData(c *gin.Context) {

	mutexDummyData.Lock()
	defer mutexDummyData.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDummyDatas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDummyData.GetDB()

	// Validate input
	var input orm.DummyDataAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create dummydata
	dummydataDB := orm.DummyDataDB{}
	dummydataDB.DummyDataPointersEncoding = input.DummyDataPointersEncoding
	dummydataDB.CopyBasicFieldsFromDummyData_WOP(&input.DummyData_WOP)

	_, err = db.Create(&dummydataDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDummyData.CheckoutPhaseOneInstance(&dummydataDB)
	dummydata := backRepo.BackRepoDummyData.Map_DummyDataDBID_DummyDataPtr[dummydataDB.ID]

	if dummydata != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), dummydata)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, dummydataDB)
}

// GetDummyData
//
// swagger:route GET /dummydatas/{ID} dummydatas getDummyData
//
// Gets the details for a dummydata.
//
// Responses:
// default: genericError
//
//	200: dummydataDBResponse
func (controller *Controller) GetDummyData(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDummyData", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDummyData.GetDB()

	// Get dummydataDB in DB
	var dummydataDB orm.DummyDataDB
	if _, err := db.First(&dummydataDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var dummydataAPI orm.DummyDataAPI
	dummydataAPI.ID = dummydataDB.ID
	dummydataAPI.DummyDataPointersEncoding = dummydataDB.DummyDataPointersEncoding
	dummydataDB.CopyBasicFieldsToDummyData_WOP(&dummydataAPI.DummyData_WOP)

	c.JSON(http.StatusOK, dummydataAPI)
}

// UpdateDummyData
//
// swagger:route PATCH /dummydatas/{ID} dummydatas updateDummyData
//
// # Update a dummydata
//
// Responses:
// default: genericError
//
//	200: dummydataDBResponse
func (controller *Controller) UpdateDummyData(c *gin.Context) {

	mutexDummyData.Lock()
	defer mutexDummyData.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDummyData", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDummyData.GetDB()

	// Validate input
	var input orm.DummyDataAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var dummydataDB orm.DummyDataDB

	// fetch the dummydata
	_, err := db.First(&dummydataDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	dummydataDB.CopyBasicFieldsFromDummyData_WOP(&input.DummyData_WOP)
	dummydataDB.DummyDataPointersEncoding = input.DummyDataPointersEncoding

	db, _ = db.Model(&dummydataDB)
	_, err = db.Updates(&dummydataDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	dummydataNew := new(models.DummyData)
	dummydataDB.CopyBasicFieldsToDummyData(dummydataNew)

	// redeem pointers
	dummydataDB.DecodePointers(backRepo, dummydataNew)

	// get stage instance from DB instance, and call callback function
	dummydataOld := backRepo.BackRepoDummyData.Map_DummyDataDBID_DummyDataPtr[dummydataDB.ID]
	if dummydataOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), dummydataOld, dummydataNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the dummydataDB
	c.JSON(http.StatusOK, dummydataDB)
}

// DeleteDummyData
//
// swagger:route DELETE /dummydatas/{ID} dummydatas deleteDummyData
//
// # Delete a dummydata
//
// default: genericError
//
//	200: dummydataDBResponse
func (controller *Controller) DeleteDummyData(c *gin.Context) {

	mutexDummyData.Lock()
	defer mutexDummyData.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDummyData", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDummyData.GetDB()

	// Get model if exist
	var dummydataDB orm.DummyDataDB
	if _, err := db.First(&dummydataDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&dummydataDB)

	// get an instance (not staged) from DB instance, and call callback function
	dummydataDeleted := new(models.DummyData)
	dummydataDB.CopyBasicFieldsToDummyData(dummydataDeleted)

	// get stage instance from DB instance, and call callback function
	dummydataStaged := backRepo.BackRepoDummyData.Map_DummyDataDBID_DummyDataPtr[dummydataDB.ID]
	if dummydataStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), dummydataStaged, dummydataDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
