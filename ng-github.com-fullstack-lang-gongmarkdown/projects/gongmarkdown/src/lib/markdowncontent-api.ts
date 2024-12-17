// insertion point for imports
import { ElementAPI } from './element-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MarkdownContentAPI {

	static GONGSTRUCT_NAME = "MarkdownContent"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for other decls

	MarkdownContentPointersEncoding: MarkdownContentPointersEncoding = new MarkdownContentPointersEncoding
}

export class MarkdownContentPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	RootID: NullInt64 = new NullInt64 // if pointer is null, Root.ID = 0

}
