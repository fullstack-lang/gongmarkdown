// insertion point for imports
import { RowAPI } from './row-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ElementAPI {

	static GONGSTRUCT_NAME = "Element"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""
	Type: string = ""

	// insertion point for other decls

	ElementPointersEncoding: ElementPointersEncoding = new ElementPointersEncoding
}

export class ElementPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	SubElements: number[] = []
	Rows: number[] = []
}
