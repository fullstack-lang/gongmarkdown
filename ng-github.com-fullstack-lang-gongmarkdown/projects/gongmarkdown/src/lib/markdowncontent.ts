// generated code - do not edit

import { MarkdownContentAPI } from './markdowncontent-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Element } from './element'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MarkdownContent {

	static GONGSTRUCT_NAME = "MarkdownContent"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for pointers and slices of pointers declarations
	Root?: Element

}

export function CopyMarkdownContentToMarkdownContentAPI(markdowncontent: MarkdownContent, markdowncontentAPI: MarkdownContentAPI) {

	markdowncontentAPI.CreatedAt = markdowncontent.CreatedAt
	markdowncontentAPI.DeletedAt = markdowncontent.DeletedAt
	markdowncontentAPI.ID = markdowncontent.ID

	// insertion point for basic fields copy operations
	markdowncontentAPI.Name = markdowncontent.Name
	markdowncontentAPI.Content = markdowncontent.Content

	// insertion point for pointer fields encoding
	markdowncontentAPI.MarkdownContentPointersEncoding.RootID.Valid = true
	if (markdowncontent.Root != undefined) {
		markdowncontentAPI.MarkdownContentPointersEncoding.RootID.Int64 = markdowncontent.Root.ID  
	} else {
		markdowncontentAPI.MarkdownContentPointersEncoding.RootID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyMarkdownContentAPIToMarkdownContent update basic, pointers and slice of pointers fields of markdowncontent
// from respectively the basic fields and encoded fields of pointers and slices of pointers of markdowncontentAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMarkdownContentAPIToMarkdownContent(markdowncontentAPI: MarkdownContentAPI, markdowncontent: MarkdownContent, frontRepo: FrontRepo) {

	markdowncontent.CreatedAt = markdowncontentAPI.CreatedAt
	markdowncontent.DeletedAt = markdowncontentAPI.DeletedAt
	markdowncontent.ID = markdowncontentAPI.ID

	// insertion point for basic fields copy operations
	markdowncontent.Name = markdowncontentAPI.Name
	markdowncontent.Content = markdowncontentAPI.Content

	// insertion point for pointer fields encoding
	markdowncontent.Root = frontRepo.map_ID_Element.get(markdowncontentAPI.MarkdownContentPointersEncoding.RootID.Int64)

	// insertion point for slice of pointers fields encoding
}
