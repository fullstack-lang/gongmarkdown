// generated code - do not edit

import { ElementAPI } from './element-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Row } from './row'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Element {

	static GONGSTRUCT_NAME = "Element"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""
	Type: string = ""

	// insertion point for pointers and slices of pointers declarations
	SubElements: Array<Element> = []
	Rows: Array<Row> = []
}

export function CopyElementToElementAPI(element: Element, elementAPI: ElementAPI) {

	elementAPI.CreatedAt = element.CreatedAt
	elementAPI.DeletedAt = element.DeletedAt
	elementAPI.ID = element.ID

	// insertion point for basic fields copy operations
	elementAPI.Name = element.Name
	elementAPI.Content = element.Content
	elementAPI.Type = element.Type

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	elementAPI.ElementPointersEncoding.SubElements = []
	for (let _element of element.SubElements) {
		elementAPI.ElementPointersEncoding.SubElements.push(_element.ID)
	}

	elementAPI.ElementPointersEncoding.Rows = []
	for (let _row of element.Rows) {
		elementAPI.ElementPointersEncoding.Rows.push(_row.ID)
	}

}

// CopyElementAPIToElement update basic, pointers and slice of pointers fields of element
// from respectively the basic fields and encoded fields of pointers and slices of pointers of elementAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyElementAPIToElement(elementAPI: ElementAPI, element: Element, frontRepo: FrontRepo) {

	element.CreatedAt = elementAPI.CreatedAt
	element.DeletedAt = elementAPI.DeletedAt
	element.ID = elementAPI.ID

	// insertion point for basic fields copy operations
	element.Name = elementAPI.Name
	element.Content = elementAPI.Content
	element.Type = elementAPI.Type

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(elementAPI.ElementPointersEncoding.SubElements)) {
		console.error('Rects is not an array:', elementAPI.ElementPointersEncoding.SubElements);
		return;
	}

	element.SubElements = new Array<Element>()
	for (let _id of elementAPI.ElementPointersEncoding.SubElements) {
		let _element = frontRepo.map_ID_Element.get(_id)
		if (_element != undefined) {
			element.SubElements.push(_element!)
		}
	}
	if (!Array.isArray(elementAPI.ElementPointersEncoding.Rows)) {
		console.error('Rects is not an array:', elementAPI.ElementPointersEncoding.Rows);
		return;
	}

	element.Rows = new Array<Row>()
	for (let _id of elementAPI.ElementPointersEncoding.Rows) {
		let _row = frontRepo.map_ID_Row.get(_id)
		if (_row != undefined) {
			element.Rows.push(_row!)
		}
	}
}
