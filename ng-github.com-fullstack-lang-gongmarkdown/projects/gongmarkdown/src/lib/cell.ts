// generated code - do not edit

import { CellAPI } from './cell-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Cell {

	static GONGSTRUCT_NAME = "Cell"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyCellToCellAPI(cell: Cell, cellAPI: CellAPI) {

	cellAPI.CreatedAt = cell.CreatedAt
	cellAPI.DeletedAt = cell.DeletedAt
	cellAPI.ID = cell.ID

	// insertion point for basic fields copy operations
	cellAPI.Name = cell.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyCellAPIToCell update basic, pointers and slice of pointers fields of cell
// from respectively the basic fields and encoded fields of pointers and slices of pointers of cellAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCellAPIToCell(cellAPI: CellAPI, cell: Cell, frontRepo: FrontRepo) {

	cell.CreatedAt = cellAPI.CreatedAt
	cell.DeletedAt = cellAPI.DeletedAt
	cell.ID = cellAPI.ID

	// insertion point for basic fields copy operations
	cell.Name = cellAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
