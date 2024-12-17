// insertion point for imports
import { AnotherDummyDataAPI } from './anotherdummydata-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DummyDataAPI {

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

	// insertion point for other decls
	DummyEnumInt_string?: string
	DummyDuration_string?: string

	DummyDataPointersEncoding: DummyDataPointersEncoding = new DummyDataPointersEncoding
}

export class DummyDataPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	DummyPointerToGongStructID: NullInt64 = new NullInt64 // if pointer is null, DummyPointerToGongStruct.ID = 0

}
