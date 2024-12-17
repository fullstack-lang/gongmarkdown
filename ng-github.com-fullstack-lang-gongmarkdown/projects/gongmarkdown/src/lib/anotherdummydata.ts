// generated code - do not edit

import { AnotherDummyDataAPI } from './anotherdummydata-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AnotherDummyData {

	static GONGSTRUCT_NAME = "AnotherDummyData"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyAnotherDummyDataToAnotherDummyDataAPI(anotherdummydata: AnotherDummyData, anotherdummydataAPI: AnotherDummyDataAPI) {

	anotherdummydataAPI.CreatedAt = anotherdummydata.CreatedAt
	anotherdummydataAPI.DeletedAt = anotherdummydata.DeletedAt
	anotherdummydataAPI.ID = anotherdummydata.ID

	// insertion point for basic fields copy operations
	anotherdummydataAPI.Name = anotherdummydata.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyAnotherDummyDataAPIToAnotherDummyData update basic, pointers and slice of pointers fields of anotherdummydata
// from respectively the basic fields and encoded fields of pointers and slices of pointers of anotherdummydataAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAnotherDummyDataAPIToAnotherDummyData(anotherdummydataAPI: AnotherDummyDataAPI, anotherdummydata: AnotherDummyData, frontRepo: FrontRepo) {

	anotherdummydata.CreatedAt = anotherdummydataAPI.CreatedAt
	anotherdummydata.DeletedAt = anotherdummydataAPI.DeletedAt
	anotherdummydata.ID = anotherdummydataAPI.ID

	// insertion point for basic fields copy operations
	anotherdummydata.Name = anotherdummydataAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
