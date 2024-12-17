// generated code - do not edit

//insertion point for imports
import { AnotherDummyDataAPI } from './anotherdummydata-api'

import { CellAPI } from './cell-api'

import { DummyDataAPI } from './dummydata-api'

import { ElementAPI } from './element-api'

import { MarkdownContentAPI } from './markdowncontent-api'

import { RowAPI } from './row-api'


export class BackRepoData {
	// insertion point for declarations
	AnotherDummyDataAPIs = new Array<AnotherDummyDataAPI>()

	CellAPIs = new Array<CellAPI>()

	DummyDataAPIs = new Array<DummyDataAPI>()

	ElementAPIs = new Array<ElementAPI>()

	MarkdownContentAPIs = new Array<MarkdownContentAPI>()

	RowAPIs = new Array<RowAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AnotherDummyDataAPIs = data?.AnotherDummyDataAPIs || [];

		this.CellAPIs = data?.CellAPIs || [];

		this.DummyDataAPIs = data?.DummyDataAPIs || [];

		this.ElementAPIs = data?.ElementAPIs || [];

		this.MarkdownContentAPIs = data?.MarkdownContentAPIs || [];

		this.RowAPIs = data?.RowAPIs || [];

	}

}