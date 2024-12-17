// generated code - do not edit

import { DummyDataAPI } from './dummydata-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { AnotherDummyData } from './anotherdummydata'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DummyData {

	static GONGSTRUCT_NAME = "DummyData"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	DummyString: string = ""
	DummyInt: number = 0
	DummyFloat: number = 0
	DummyBool: boolean = false
	DummyEnumString: string = ""
	DummyEnumInt: number = 0
	DummyTime: Date = new Date
	DummyDuration: number = 0

	// insertion point for pointers and slices of pointers declarations
	DummyEnumInt_string?: string
	DummyDuration_string?: string
	DummyPointerToGongStruct?: AnotherDummyData

}

export function CopyDummyDataToDummyDataAPI(dummydata: DummyData, dummydataAPI: DummyDataAPI) {

	dummydataAPI.CreatedAt = dummydata.CreatedAt
	dummydataAPI.DeletedAt = dummydata.DeletedAt
	dummydataAPI.ID = dummydata.ID

	// insertion point for basic fields copy operations
	dummydataAPI.Name = dummydata.Name
	dummydataAPI.DummyString = dummydata.DummyString
	dummydataAPI.DummyInt = dummydata.DummyInt
	dummydataAPI.DummyFloat = dummydata.DummyFloat
	dummydataAPI.DummyBool = dummydata.DummyBool
	dummydataAPI.DummyEnumString = dummydata.DummyEnumString
	dummydataAPI.DummyEnumInt = dummydata.DummyEnumInt
	dummydataAPI.DummyTime = dummydata.DummyTime
	dummydataAPI.DummyDuration = dummydata.DummyDuration

	// insertion point for pointer fields encoding
	dummydataAPI.DummyDataPointersEncoding.DummyPointerToGongStructID.Valid = true
	if (dummydata.DummyPointerToGongStruct != undefined) {
		dummydataAPI.DummyDataPointersEncoding.DummyPointerToGongStructID.Int64 = dummydata.DummyPointerToGongStruct.ID  
	} else {
		dummydataAPI.DummyDataPointersEncoding.DummyPointerToGongStructID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyDummyDataAPIToDummyData update basic, pointers and slice of pointers fields of dummydata
// from respectively the basic fields and encoded fields of pointers and slices of pointers of dummydataAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDummyDataAPIToDummyData(dummydataAPI: DummyDataAPI, dummydata: DummyData, frontRepo: FrontRepo) {

	dummydata.CreatedAt = dummydataAPI.CreatedAt
	dummydata.DeletedAt = dummydataAPI.DeletedAt
	dummydata.ID = dummydataAPI.ID

	// insertion point for basic fields copy operations
	dummydata.Name = dummydataAPI.Name
	dummydata.DummyString = dummydataAPI.DummyString
	dummydata.DummyInt = dummydataAPI.DummyInt
	dummydata.DummyFloat = dummydataAPI.DummyFloat
	dummydata.DummyBool = dummydataAPI.DummyBool
	dummydata.DummyEnumString = dummydataAPI.DummyEnumString
	dummydata.DummyEnumInt = dummydataAPI.DummyEnumInt
	dummydata.DummyTime = dummydataAPI.DummyTime
	dummydata.DummyDuration = dummydataAPI.DummyDuration

	// insertion point for pointer fields encoding
	dummydata.DummyPointerToGongStruct = frontRepo.map_ID_AnotherDummyData.get(dummydataAPI.DummyDataPointersEncoding.DummyPointerToGongStructID.Int64)

	// insertion point for slice of pointers fields encoding
}
