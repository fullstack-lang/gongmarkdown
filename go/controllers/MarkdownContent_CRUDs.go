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
var __MarkdownContent__dummysDeclaration__ models.MarkdownContent
var __MarkdownContent_time__dummyDeclaration time.Duration

var mutexMarkdownContent sync.Mutex

// An MarkdownContentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMarkdownContent updateMarkdownContent deleteMarkdownContent
type MarkdownContentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MarkdownContentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMarkdownContent updateMarkdownContent
type MarkdownContentInput struct {
	// The MarkdownContent to submit or modify
	// in: body
	MarkdownContent *orm.MarkdownContentAPI
}

// GetMarkdownContents
//
// swagger:route GET /markdowncontents markdowncontents getMarkdownContents
//
// # Get all markdowncontents
//
// Responses:
// default: genericError
//
//	200: markdowncontentDBResponse
func (controller *Controller) GetMarkdownContents(c *gin.Context) {

	// source slice
	var markdowncontentDBs []orm.MarkdownContentDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMarkdownContents", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMarkdownContent.GetDB()

	_, err := db.Find(&markdowncontentDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	markdowncontentAPIs := make([]orm.MarkdownContentAPI, 0)

	// for each markdowncontent, update fields from the database nullable fields
	for idx := range markdowncontentDBs {
		markdowncontentDB := &markdowncontentDBs[idx]
		_ = markdowncontentDB
		var markdowncontentAPI orm.MarkdownContentAPI

		// insertion point for updating fields
		markdowncontentAPI.ID = markdowncontentDB.ID
		markdowncontentDB.CopyBasicFieldsToMarkdownContent_WOP(&markdowncontentAPI.MarkdownContent_WOP)
		markdowncontentAPI.MarkdownContentPointersEncoding = markdowncontentDB.MarkdownContentPointersEncoding
		markdowncontentAPIs = append(markdowncontentAPIs, markdowncontentAPI)
	}

	c.JSON(http.StatusOK, markdowncontentAPIs)
}

// PostMarkdownContent
//
// swagger:route POST /markdowncontents markdowncontents postMarkdownContent
//
// Creates a markdowncontent
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMarkdownContent(c *gin.Context) {

	mutexMarkdownContent.Lock()
	defer mutexMarkdownContent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMarkdownContents", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMarkdownContent.GetDB()

	// Validate input
	var input orm.MarkdownContentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create markdowncontent
	markdowncontentDB := orm.MarkdownContentDB{}
	markdowncontentDB.MarkdownContentPointersEncoding = input.MarkdownContentPointersEncoding
	markdowncontentDB.CopyBasicFieldsFromMarkdownContent_WOP(&input.MarkdownContent_WOP)

	_, err = db.Create(&markdowncontentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMarkdownContent.CheckoutPhaseOneInstance(&markdowncontentDB)
	markdowncontent := backRepo.BackRepoMarkdownContent.Map_MarkdownContentDBID_MarkdownContentPtr[markdowncontentDB.ID]

	if markdowncontent != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), markdowncontent)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, markdowncontentDB)
}

// GetMarkdownContent
//
// swagger:route GET /markdowncontents/{ID} markdowncontents getMarkdownContent
//
// Gets the details for a markdowncontent.
//
// Responses:
// default: genericError
//
//	200: markdowncontentDBResponse
func (controller *Controller) GetMarkdownContent(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMarkdownContent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMarkdownContent.GetDB()

	// Get markdowncontentDB in DB
	var markdowncontentDB orm.MarkdownContentDB
	if _, err := db.First(&markdowncontentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var markdowncontentAPI orm.MarkdownContentAPI
	markdowncontentAPI.ID = markdowncontentDB.ID
	markdowncontentAPI.MarkdownContentPointersEncoding = markdowncontentDB.MarkdownContentPointersEncoding
	markdowncontentDB.CopyBasicFieldsToMarkdownContent_WOP(&markdowncontentAPI.MarkdownContent_WOP)

	c.JSON(http.StatusOK, markdowncontentAPI)
}

// UpdateMarkdownContent
//
// swagger:route PATCH /markdowncontents/{ID} markdowncontents updateMarkdownContent
//
// # Update a markdowncontent
//
// Responses:
// default: genericError
//
//	200: markdowncontentDBResponse
func (controller *Controller) UpdateMarkdownContent(c *gin.Context) {

	mutexMarkdownContent.Lock()
	defer mutexMarkdownContent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMarkdownContent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMarkdownContent.GetDB()

	// Validate input
	var input orm.MarkdownContentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var markdowncontentDB orm.MarkdownContentDB

	// fetch the markdowncontent
	_, err := db.First(&markdowncontentDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	markdowncontentDB.CopyBasicFieldsFromMarkdownContent_WOP(&input.MarkdownContent_WOP)
	markdowncontentDB.MarkdownContentPointersEncoding = input.MarkdownContentPointersEncoding

	db, _ = db.Model(&markdowncontentDB)
	_, err = db.Updates(&markdowncontentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	markdowncontentNew := new(models.MarkdownContent)
	markdowncontentDB.CopyBasicFieldsToMarkdownContent(markdowncontentNew)

	// redeem pointers
	markdowncontentDB.DecodePointers(backRepo, markdowncontentNew)

	// get stage instance from DB instance, and call callback function
	markdowncontentOld := backRepo.BackRepoMarkdownContent.Map_MarkdownContentDBID_MarkdownContentPtr[markdowncontentDB.ID]
	if markdowncontentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), markdowncontentOld, markdowncontentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the markdowncontentDB
	c.JSON(http.StatusOK, markdowncontentDB)
}

// DeleteMarkdownContent
//
// swagger:route DELETE /markdowncontents/{ID} markdowncontents deleteMarkdownContent
//
// # Delete a markdowncontent
//
// default: genericError
//
//	200: markdowncontentDBResponse
func (controller *Controller) DeleteMarkdownContent(c *gin.Context) {

	mutexMarkdownContent.Lock()
	defer mutexMarkdownContent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMarkdownContent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmarkdown/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMarkdownContent.GetDB()

	// Get model if exist
	var markdowncontentDB orm.MarkdownContentDB
	if _, err := db.First(&markdowncontentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&markdowncontentDB)

	// get an instance (not staged) from DB instance, and call callback function
	markdowncontentDeleted := new(models.MarkdownContent)
	markdowncontentDB.CopyBasicFieldsToMarkdownContent(markdowncontentDeleted)

	// get stage instance from DB instance, and call callback function
	markdowncontentStaged := backRepo.BackRepoMarkdownContent.Map_MarkdownContentDBID_MarkdownContentPtr[markdowncontentDB.ID]
	if markdowncontentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), markdowncontentStaged, markdowncontentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
