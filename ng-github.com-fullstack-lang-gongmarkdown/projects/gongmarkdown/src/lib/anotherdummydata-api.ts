// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AnotherDummyDataAPI {

	static GONGSTRUCT_NAME = "AnotherDummyData"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	AnotherDummyDataPointersEncoding: AnotherDummyDataPointersEncoding = new AnotherDummyDataPointersEncoding
}

export class AnotherDummyDataPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
