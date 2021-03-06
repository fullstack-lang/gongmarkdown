// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __ModelPkg__dummysDeclaration__ models.ModelPkg
var __ModelPkg_time__dummyDeclaration time.Duration

// An ModelPkgID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getModelPkg updateModelPkg deleteModelPkg
type ModelPkgID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ModelPkgInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postModelPkg updateModelPkg
type ModelPkgInput struct {
	// The ModelPkg to submit or modify
	// in: body
	ModelPkg *orm.ModelPkgAPI
}

// GetModelPkgs
//
// swagger:route GET /modelpkgs modelpkgs getModelPkgs
//
// Get all modelpkgs
//
// Responses:
//    default: genericError
//        200: modelpkgDBsResponse
func GetModelPkgs(c *gin.Context) {
	db := orm.BackRepo.BackRepoModelPkg.GetDB()

	// source slice
	var modelpkgDBs []orm.ModelPkgDB
	query := db.Find(&modelpkgDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	modelpkgAPIs := make([]orm.ModelPkgAPI, 0)

	// for each modelpkg, update fields from the database nullable fields
	for idx := range modelpkgDBs {
		modelpkgDB := &modelpkgDBs[idx]
		_ = modelpkgDB
		var modelpkgAPI orm.ModelPkgAPI

		// insertion point for updating fields
		modelpkgAPI.ID = modelpkgDB.ID
		modelpkgDB.CopyBasicFieldsToModelPkg(&modelpkgAPI.ModelPkg)
		modelpkgAPI.ModelPkgPointersEnconding = modelpkgDB.ModelPkgPointersEnconding
		modelpkgAPIs = append(modelpkgAPIs, modelpkgAPI)
	}

	c.JSON(http.StatusOK, modelpkgAPIs)
}

// PostModelPkg
//
// swagger:route POST /modelpkgs modelpkgs postModelPkg
//
// Creates a modelpkg
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: modelpkgDBResponse
func PostModelPkg(c *gin.Context) {
	db := orm.BackRepo.BackRepoModelPkg.GetDB()

	// Validate input
	var input orm.ModelPkgAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create modelpkg
	modelpkgDB := orm.ModelPkgDB{}
	modelpkgDB.ModelPkgPointersEnconding = input.ModelPkgPointersEnconding
	modelpkgDB.CopyBasicFieldsFromModelPkg(&input.ModelPkg)

	query := db.Create(&modelpkgDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, modelpkgDB)
}

// GetModelPkg
//
// swagger:route GET /modelpkgs/{ID} modelpkgs getModelPkg
//
// Gets the details for a modelpkg.
//
// Responses:
//    default: genericError
//        200: modelpkgDBResponse
func GetModelPkg(c *gin.Context) {
	db := orm.BackRepo.BackRepoModelPkg.GetDB()

	// Get modelpkgDB in DB
	var modelpkgDB orm.ModelPkgDB
	if err := db.First(&modelpkgDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var modelpkgAPI orm.ModelPkgAPI
	modelpkgAPI.ID = modelpkgDB.ID
	modelpkgAPI.ModelPkgPointersEnconding = modelpkgDB.ModelPkgPointersEnconding
	modelpkgDB.CopyBasicFieldsToModelPkg(&modelpkgAPI.ModelPkg)

	c.JSON(http.StatusOK, modelpkgAPI)
}

// UpdateModelPkg
//
// swagger:route PATCH /modelpkgs/{ID} modelpkgs updateModelPkg
//
// Update a modelpkg
//
// Responses:
//    default: genericError
//        200: modelpkgDBResponse
func UpdateModelPkg(c *gin.Context) {
	db := orm.BackRepo.BackRepoModelPkg.GetDB()

	// Get model if exist
	var modelpkgDB orm.ModelPkgDB

	// fetch the modelpkg
	query := db.First(&modelpkgDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.ModelPkgAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	modelpkgDB.CopyBasicFieldsFromModelPkg(&input.ModelPkg)
	modelpkgDB.ModelPkgPointersEnconding = input.ModelPkgPointersEnconding

	query = db.Model(&modelpkgDB).Updates(modelpkgDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the modelpkgDB
	c.JSON(http.StatusOK, modelpkgDB)
}

// DeleteModelPkg
//
// swagger:route DELETE /modelpkgs/{ID} modelpkgs deleteModelPkg
//
// Delete a modelpkg
//
// Responses:
//    default: genericError
func DeleteModelPkg(c *gin.Context) {
	db := orm.BackRepo.BackRepoModelPkg.GetDB()

	// Get model if exist
	var modelpkgDB orm.ModelPkgDB
	if err := db.First(&modelpkgDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&modelpkgDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
