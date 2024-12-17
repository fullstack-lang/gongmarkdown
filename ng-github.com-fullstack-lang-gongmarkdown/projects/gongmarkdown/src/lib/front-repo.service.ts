import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { AnotherDummyDataAPI } from './anotherdummydata-api'
import { AnotherDummyData, CopyAnotherDummyDataAPIToAnotherDummyData } from './anotherdummydata'
import { AnotherDummyDataService } from './anotherdummydata.service'

import { CellAPI } from './cell-api'
import { Cell, CopyCellAPIToCell } from './cell'
import { CellService } from './cell.service'

import { DummyDataAPI } from './dummydata-api'
import { DummyData, CopyDummyDataAPIToDummyData } from './dummydata'
import { DummyDataService } from './dummydata.service'

import { ElementAPI } from './element-api'
import { Element, CopyElementAPIToElement } from './element'
import { ElementService } from './element.service'

import { MarkdownContentAPI } from './markdowncontent-api'
import { MarkdownContent, CopyMarkdownContentAPIToMarkdownContent } from './markdowncontent'
import { MarkdownContentService } from './markdowncontent.service'

import { RowAPI } from './row-api'
import { Row, CopyRowAPIToRow } from './row'
import { RowService } from './row.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gongmarkdown/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_AnotherDummyDatas = new Array<AnotherDummyData>() // array of front instances
	map_ID_AnotherDummyData = new Map<number, AnotherDummyData>() // map of front instances

	array_Cells = new Array<Cell>() // array of front instances
	map_ID_Cell = new Map<number, Cell>() // map of front instances

	array_DummyDatas = new Array<DummyData>() // array of front instances
	map_ID_DummyData = new Map<number, DummyData>() // map of front instances

	array_Elements = new Array<Element>() // array of front instances
	map_ID_Element = new Map<number, Element>() // map of front instances

	array_MarkdownContents = new Array<MarkdownContent>() // array of front instances
	map_ID_MarkdownContent = new Map<number, MarkdownContent>() // map of front instances

	array_Rows = new Array<Row>() // array of front instances
	map_ID_Row = new Map<number, Row>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'AnotherDummyData':
				return this.array_AnotherDummyDatas as unknown as Array<Type>
			case 'Cell':
				return this.array_Cells as unknown as Array<Type>
			case 'DummyData':
				return this.array_DummyDatas as unknown as Array<Type>
			case 'Element':
				return this.array_Elements as unknown as Array<Type>
			case 'MarkdownContent':
				return this.array_MarkdownContents as unknown as Array<Type>
			case 'Row':
				return this.array_Rows as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'AnotherDummyData':
				return this.map_ID_AnotherDummyData as unknown as Map<number, Type>
			case 'Cell':
				return this.map_ID_Cell as unknown as Map<number, Type>
			case 'DummyData':
				return this.map_ID_DummyData as unknown as Map<number, Type>
			case 'Element':
				return this.map_ID_Element as unknown as Map<number, Type>
			case 'MarkdownContent':
				return this.map_ID_MarkdownContent as unknown as Map<number, Type>
			case 'Row':
				return this.map_ID_Row as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized");
		}
	}
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
	ID: number = 0 // ID of the calling instance

	// the reverse pointer is the name of the generated field on the destination
	// struct of the ONE-MANY association
	ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
	OrderingMode: boolean = false // if true, this is for ordering items

	// there are different selection mode : ONE_MANY or MANY_MANY
	SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

	// used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
	//
	// In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
	// 
	// in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
	// at the end of the ONE-MANY association
	SourceStruct: string = ""	// The "Aclass"
	SourceField: string = "" // the "AnarrayofbUse"
	IntermediateStruct: string = "" // the "AclassBclassUse" 
	IntermediateStructField: string = "" // the "Bclass" as field
	NextAssociationStruct: string = "" // the "Bclass"

	GONG__StackPath: string = ""
}

export enum SelectionMode {
	ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
	MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
	providedIn: 'root'
})
export class FrontRepoService {

	GONG__StackPath: string = ""
	private socket: WebSocket | undefined

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	};

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	constructor(
		private http: HttpClient, // insertion point sub template 
		private anotherdummydataService: AnotherDummyDataService,
		private cellService: CellService,
		private dummydataService: DummyDataService,
		private elementService: ElementService,
		private markdowncontentService: MarkdownContentService,
		private rowService: RowService,
	) { }

	// postService provides a post function for each struct name
	postService(structName: string, instanceToBePosted: any) {
		let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
		let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

		servicePostFunction(instanceToBePosted).subscribe(
			instance => {
				let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("post")
			}
		);
	}

	// deleteService provides a delete function for each struct name
	deleteService(structName: string, instanceToBeDeleted: any) {
		let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
		let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

		serviceDeleteFunction(instanceToBeDeleted).subscribe(
			instance => {
				let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("delete")
			}
		);
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<AnotherDummyDataAPI[]>,
		Observable<CellAPI[]>,
		Observable<DummyDataAPI[]>,
		Observable<ElementAPI[]>,
		Observable<MarkdownContentAPI[]>,
		Observable<RowAPI[]>,
	] = [
			// Using "combineLatest" with a placeholder observable.
			//
			// This allows the typescript compiler to pass when no GongStruct is present in the front API
			//
			// The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
			// This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
			// expectation for a non-empty array of observables.
			of(null), // 
			// insertion point sub template
			this.anotherdummydataService.getAnotherDummyDatas(this.GONG__StackPath, this.frontRepo),
			this.cellService.getCells(this.GONG__StackPath, this.frontRepo),
			this.dummydataService.getDummyDatas(this.GONG__StackPath, this.frontRepo),
			this.elementService.getElements(this.GONG__StackPath, this.frontRepo),
			this.markdowncontentService.getMarkdownContents(this.GONG__StackPath, this.frontRepo),
			this.rowService.getRows(this.GONG__StackPath, this.frontRepo),
		];

	//
	// pull performs a GET on all struct of the stack and redeem association pointers 
	//
	// This is an observable. Therefore, the control flow forks with
	// - pull() return immediatly the observable
	// - the observable observer, if it subscribe, is called when all GET calls are performs
	pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath

		this.observableFrontRepo = [
			of(null), // see above for justification
			// insertion point sub template
			this.anotherdummydataService.getAnotherDummyDatas(this.GONG__StackPath, this.frontRepo),
			this.cellService.getCells(this.GONG__StackPath, this.frontRepo),
			this.dummydataService.getDummyDatas(this.GONG__StackPath, this.frontRepo),
			this.elementService.getElements(this.GONG__StackPath, this.frontRepo),
			this.markdowncontentService.getMarkdownContents(this.GONG__StackPath, this.frontRepo),
			this.rowService.getRows(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						anotherdummydatas_,
						cells_,
						dummydatas_,
						elements_,
						markdowncontents_,
						rows_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var anotherdummydatas: AnotherDummyDataAPI[]
						anotherdummydatas = anotherdummydatas_ as AnotherDummyDataAPI[]
						var cells: CellAPI[]
						cells = cells_ as CellAPI[]
						var dummydatas: DummyDataAPI[]
						dummydatas = dummydatas_ as DummyDataAPI[]
						var elements: ElementAPI[]
						elements = elements_ as ElementAPI[]
						var markdowncontents: MarkdownContentAPI[]
						markdowncontents = markdowncontents_ as MarkdownContentAPI[]
						var rows: RowAPI[]
						rows = rows_ as RowAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_AnotherDummyDatas = []
						this.frontRepo.map_ID_AnotherDummyData.clear()

						anotherdummydatas.forEach(
							anotherdummydataAPI => {
								let anotherdummydata = new AnotherDummyData
								this.frontRepo.array_AnotherDummyDatas.push(anotherdummydata)
								this.frontRepo.map_ID_AnotherDummyData.set(anotherdummydataAPI.ID, anotherdummydata)
							}
						)

						// init the arrays
						this.frontRepo.array_Cells = []
						this.frontRepo.map_ID_Cell.clear()

						cells.forEach(
							cellAPI => {
								let cell = new Cell
								this.frontRepo.array_Cells.push(cell)
								this.frontRepo.map_ID_Cell.set(cellAPI.ID, cell)
							}
						)

						// init the arrays
						this.frontRepo.array_DummyDatas = []
						this.frontRepo.map_ID_DummyData.clear()

						dummydatas.forEach(
							dummydataAPI => {
								let dummydata = new DummyData
								this.frontRepo.array_DummyDatas.push(dummydata)
								this.frontRepo.map_ID_DummyData.set(dummydataAPI.ID, dummydata)
							}
						)

						// init the arrays
						this.frontRepo.array_Elements = []
						this.frontRepo.map_ID_Element.clear()

						elements.forEach(
							elementAPI => {
								let element = new Element
								this.frontRepo.array_Elements.push(element)
								this.frontRepo.map_ID_Element.set(elementAPI.ID, element)
							}
						)

						// init the arrays
						this.frontRepo.array_MarkdownContents = []
						this.frontRepo.map_ID_MarkdownContent.clear()

						markdowncontents.forEach(
							markdowncontentAPI => {
								let markdowncontent = new MarkdownContent
								this.frontRepo.array_MarkdownContents.push(markdowncontent)
								this.frontRepo.map_ID_MarkdownContent.set(markdowncontentAPI.ID, markdowncontent)
							}
						)

						// init the arrays
						this.frontRepo.array_Rows = []
						this.frontRepo.map_ID_Row.clear()

						rows.forEach(
							rowAPI => {
								let row = new Row
								this.frontRepo.array_Rows.push(row)
								this.frontRepo.map_ID_Row.set(rowAPI.ID, row)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						anotherdummydatas.forEach(
							anotherdummydataAPI => {
								let anotherdummydata = this.frontRepo.map_ID_AnotherDummyData.get(anotherdummydataAPI.ID)
								CopyAnotherDummyDataAPIToAnotherDummyData(anotherdummydataAPI, anotherdummydata!, this.frontRepo)
							}
						)

						// fill up front objects
						cells.forEach(
							cellAPI => {
								let cell = this.frontRepo.map_ID_Cell.get(cellAPI.ID)
								CopyCellAPIToCell(cellAPI, cell!, this.frontRepo)
							}
						)

						// fill up front objects
						dummydatas.forEach(
							dummydataAPI => {
								let dummydata = this.frontRepo.map_ID_DummyData.get(dummydataAPI.ID)
								CopyDummyDataAPIToDummyData(dummydataAPI, dummydata!, this.frontRepo)
							}
						)

						// fill up front objects
						elements.forEach(
							elementAPI => {
								let element = this.frontRepo.map_ID_Element.get(elementAPI.ID)
								CopyElementAPIToElement(elementAPI, element!, this.frontRepo)
							}
						)

						// fill up front objects
						markdowncontents.forEach(
							markdowncontentAPI => {
								let markdowncontent = this.frontRepo.map_ID_MarkdownContent.get(markdowncontentAPI.ID)
								CopyMarkdownContentAPIToMarkdownContent(markdowncontentAPI, markdowncontent!, this.frontRepo)
							}
						)

						// fill up front objects
						rows.forEach(
							rowAPI => {
								let row = this.frontRepo.map_ID_Row.get(rowAPI.ID)
								CopyRowAPIToRow(rowAPI, row!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	public connectToWebSocket(GONG__StackPath: string): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath


		let params = new HttpParams().set("GONG__StackPath", this.GONG__StackPath)
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gongmarkdown/go/v1/ws/stage'
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`
		this.socket = new WebSocket(url)

		return new Observable(observer => {
			this.socket!.onmessage = event => {


				const backRepoData = new BackRepoData(JSON.parse(event.data))

				let frontRepo = new (FrontRepo)

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_AnotherDummyDatas = []
				frontRepo.map_ID_AnotherDummyData.clear()

				backRepoData.AnotherDummyDataAPIs.forEach(
					anotherdummydataAPI => {
						let anotherdummydata = new AnotherDummyData
						frontRepo.array_AnotherDummyDatas.push(anotherdummydata)
						frontRepo.map_ID_AnotherDummyData.set(anotherdummydataAPI.ID, anotherdummydata)
					}
				)

				// init the arrays
				frontRepo.array_Cells = []
				frontRepo.map_ID_Cell.clear()

				backRepoData.CellAPIs.forEach(
					cellAPI => {
						let cell = new Cell
						frontRepo.array_Cells.push(cell)
						frontRepo.map_ID_Cell.set(cellAPI.ID, cell)
					}
				)

				// init the arrays
				frontRepo.array_DummyDatas = []
				frontRepo.map_ID_DummyData.clear()

				backRepoData.DummyDataAPIs.forEach(
					dummydataAPI => {
						let dummydata = new DummyData
						frontRepo.array_DummyDatas.push(dummydata)
						frontRepo.map_ID_DummyData.set(dummydataAPI.ID, dummydata)
					}
				)

				// init the arrays
				frontRepo.array_Elements = []
				frontRepo.map_ID_Element.clear()

				backRepoData.ElementAPIs.forEach(
					elementAPI => {
						let element = new Element
						frontRepo.array_Elements.push(element)
						frontRepo.map_ID_Element.set(elementAPI.ID, element)
					}
				)

				// init the arrays
				frontRepo.array_MarkdownContents = []
				frontRepo.map_ID_MarkdownContent.clear()

				backRepoData.MarkdownContentAPIs.forEach(
					markdowncontentAPI => {
						let markdowncontent = new MarkdownContent
						frontRepo.array_MarkdownContents.push(markdowncontent)
						frontRepo.map_ID_MarkdownContent.set(markdowncontentAPI.ID, markdowncontent)
					}
				)

				// init the arrays
				frontRepo.array_Rows = []
				frontRepo.map_ID_Row.clear()

				backRepoData.RowAPIs.forEach(
					rowAPI => {
						let row = new Row
						frontRepo.array_Rows.push(row)
						frontRepo.map_ID_Row.set(rowAPI.ID, row)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.AnotherDummyDataAPIs.forEach(
					anotherdummydataAPI => {
						let anotherdummydata = frontRepo.map_ID_AnotherDummyData.get(anotherdummydataAPI.ID)
						CopyAnotherDummyDataAPIToAnotherDummyData(anotherdummydataAPI, anotherdummydata!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CellAPIs.forEach(
					cellAPI => {
						let cell = frontRepo.map_ID_Cell.get(cellAPI.ID)
						CopyCellAPIToCell(cellAPI, cell!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.DummyDataAPIs.forEach(
					dummydataAPI => {
						let dummydata = frontRepo.map_ID_DummyData.get(dummydataAPI.ID)
						CopyDummyDataAPIToDummyData(dummydataAPI, dummydata!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ElementAPIs.forEach(
					elementAPI => {
						let element = frontRepo.map_ID_Element.get(elementAPI.ID)
						CopyElementAPIToElement(elementAPI, element!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.MarkdownContentAPIs.forEach(
					markdowncontentAPI => {
						let markdowncontent = frontRepo.map_ID_MarkdownContent.get(markdowncontentAPI.ID)
						CopyMarkdownContentAPIToMarkdownContent(markdowncontentAPI, markdowncontent!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.RowAPIs.forEach(
					rowAPI => {
						let row = frontRepo.map_ID_Row.get(rowAPI.ID)
						CopyRowAPIToRow(rowAPI, row!, frontRepo)
					}
				)



				observer.next(frontRepo)
			}
			this.socket!.onerror = event => {
				observer.error(event)
			}
			this.socket!.onclose = event => {
				observer.complete()
			}

			return () => {
				this.socket!.close()
			}
		})
	}
}

// insertion point for get unique ID per struct 
export function getAnotherDummyDataUniqueID(id: number): number {
	return 31 * id
}
export function getCellUniqueID(id: number): number {
	return 37 * id
}
export function getDummyDataUniqueID(id: number): number {
	return 41 * id
}
export function getElementUniqueID(id: number): number {
	return 43 * id
}
export function getMarkdownContentUniqueID(id: number): number {
	return 47 * id
}
export function getRowUniqueID(id: number): number {
	return 53 * id
}
