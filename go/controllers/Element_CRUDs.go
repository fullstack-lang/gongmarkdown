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
var __Element__dummysDeclaration__ models.Element
var __Element_time__dummyDeclaration time.Duration

var mutexElement sync.Mutex

// An ElementID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getElement updateElement deleteElement
type ElementID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ElementInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postElement updateElement
type ElementInput struct {
	// The Element to submit or modify
	// in: body
	Element *orm.ElementAPI
}

// GetElements
//
// swagger:route GET /elements elements getElements
//
// # Get all elements
//
// Responses:
// default: genericError
//
//	200: elementDBResponse
func (controller *Controller) GetElements(c *gin.Context) {

	// source slice
	var elementDBs []orm.ElementDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetElements", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoElement.GetDB()

	_, err := db.Find(&elementDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	elementAPIs := make([]orm.ElementAPI, 0)

	// for each element, update fields from the database nullable fields
	for idx := range elementDBs {
		elementDB := &elementDBs[idx]
		_ = elementDB
		var elementAPI orm.ElementAPI

		// insertion point for updating fields
		elementAPI.ID = elementDB.ID
		elementDB.CopyBasicFieldsToElement_WOP(&elementAPI.Element_WOP)
		elementAPI.ElementPointersEncoding = elementDB.ElementPointersEncoding
		elementAPIs = append(elementAPIs, elementAPI)
	}

	c.JSON(http.StatusOK, elementAPIs)
}

// PostElement
//
// swagger:route POST /elements elements postElement
//
// Creates a element
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostElement(c *gin.Context) {

	mutexElement.Lock()
	defer mutexElement.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostElements", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoElement.GetDB()

	// Validate input
	var input orm.ElementAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create element
	elementDB := orm.ElementDB{}
	elementDB.ElementPointersEncoding = input.ElementPointersEncoding
	elementDB.CopyBasicFieldsFromElement_WOP(&input.Element_WOP)

	_, err = db.Create(&elementDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoElement.CheckoutPhaseOneInstance(&elementDB)
	element := backRepo.BackRepoElement.Map_ElementDBID_ElementPtr[elementDB.ID]

	if element != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), element)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, elementDB)
}

// GetElement
//
// swagger:route GET /elements/{ID} elements getElement
//
// Gets the details for a element.
//
// Responses:
// default: genericError
//
//	200: elementDBResponse
func (controller *Controller) GetElement(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetElement", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoElement.GetDB()

	// Get elementDB in DB
	var elementDB orm.ElementDB
	if _, err := db.First(&elementDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var elementAPI orm.ElementAPI
	elementAPI.ID = elementDB.ID
	elementAPI.ElementPointersEncoding = elementDB.ElementPointersEncoding
	elementDB.CopyBasicFieldsToElement_WOP(&elementAPI.Element_WOP)

	c.JSON(http.StatusOK, elementAPI)
}

// UpdateElement
//
// swagger:route PATCH /elements/{ID} elements updateElement
//
// # Update a element
//
// Responses:
// default: genericError
//
//	200: elementDBResponse
func (controller *Controller) UpdateElement(c *gin.Context) {

	mutexElement.Lock()
	defer mutexElement.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateElement", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoElement.GetDB()

	// Validate input
	var input orm.ElementAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var elementDB orm.ElementDB

	// fetch the element
	_, err := db.First(&elementDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	elementDB.CopyBasicFieldsFromElement_WOP(&input.Element_WOP)
	elementDB.ElementPointersEncoding = input.ElementPointersEncoding

	db, _ = db.Model(&elementDB)
	_, err = db.Updates(&elementDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	elementNew := new(models.Element)
	elementDB.CopyBasicFieldsToElement(elementNew)

	// redeem pointers
	elementDB.DecodePointers(backRepo, elementNew)

	// get stage instance from DB instance, and call callback function
	elementOld := backRepo.BackRepoElement.Map_ElementDBID_ElementPtr[elementDB.ID]
	if elementOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), elementOld, elementNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the elementDB
	c.JSON(http.StatusOK, elementDB)
}

// DeleteElement
//
// swagger:route DELETE /elements/{ID} elements deleteElement
//
// # Delete a element
//
// default: genericError
//
//	200: elementDBResponse
func (controller *Controller) DeleteElement(c *gin.Context) {

	mutexElement.Lock()
	defer mutexElement.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteElement", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoElement.GetDB()

	// Get model if exist
	var elementDB orm.ElementDB
	if _, err := db.First(&elementDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&elementDB)

	// get an instance (not staged) from DB instance, and call callback function
	elementDeleted := new(models.Element)
	elementDB.CopyBasicFieldsToElement(elementDeleted)

	// get stage instance from DB instance, and call callback function
	elementStaged := backRepo.BackRepoElement.Map_ElementDBID_ElementPtr[elementDB.ID]
	if elementStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), elementStaged, elementDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
