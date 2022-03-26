// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ElementDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""
	IsTitle: boolean = false

	// insertion point for other declarations
	SubElements?: Array<ElementDB>
	Element_SubElementsDBID: NullInt64 = new NullInt64
	Element_SubElementsDBID_Index: NullInt64  = new NullInt64 // store the index of the element instance in Element.SubElements
	Element_SubElements_reverse?: ElementDB 

}
