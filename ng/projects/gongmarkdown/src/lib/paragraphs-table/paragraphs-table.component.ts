// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { ParagraphDB } from '../paragraph-db'
import { ParagraphService } from '../paragraph.service'

// insertion point for additional imports

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-paragraphstable',
  templateUrl: './paragraphs-table.component.html',
  styleUrls: ['./paragraphs-table.component.css'],
})
export class ParagraphsTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Paragraph instances
  selection: SelectionModel<ParagraphDB> = new (SelectionModel)
  initialSelection = new Array<ParagraphDB>()

  // the data source for the table
  paragraphs: ParagraphDB[] = []
  matTableDataSource: MatTableDataSource<ParagraphDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.paragraphs
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (paragraphDB: ParagraphDB, property: string) => {
      switch (property) {
        case 'ID':
          return paragraphDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return paragraphDB.Name;

        case 'Content':
          return paragraphDB.Content;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (paragraphDB: ParagraphDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the paragraphDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += paragraphDB.Name.toLowerCase()
      mergedContent += paragraphDB.Content.toLowerCase()

      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private paragraphService: ParagraphService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of paragraph instances
    public dialogRef: MatDialogRef<ParagraphsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.paragraphService.ParagraphServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getParagraphs()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Content",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Content",
      ]
      this.selection = new SelectionModel<ParagraphDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getParagraphs()
    this.matTableDataSource = new MatTableDataSource(this.paragraphs)
  }

  getParagraphs(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.paragraphs = this.frontRepo.Paragraphs_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let paragraph of this.paragraphs) {
            let ID = this.dialogData.ID
            let revPointer = paragraph[this.dialogData.ReversePointer as keyof ParagraphDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(paragraph)
            }
            this.selection = new SelectionModel<ParagraphDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, ParagraphDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as ParagraphDB[]
          for (let associationInstance of sourceField) {
            let paragraph = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as ParagraphDB
            this.initialSelection.push(paragraph)
          }

          this.selection = new SelectionModel<ParagraphDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.paragraphs
      }
    )
  }

  // newParagraph initiate a new paragraph
  // create a new Paragraph objet
  newParagraph() {
  }

  deleteParagraph(paragraphID: number, paragraph: ParagraphDB) {
    // list of paragraphs is truncated of paragraph before the delete
    this.paragraphs = this.paragraphs.filter(h => h !== paragraph);

    this.paragraphService.deleteParagraph(paragraphID).subscribe(
      paragraph => {
        this.paragraphService.ParagraphServiceChanged.next("delete")
      }
    );
  }

  editParagraph(paragraphID: number, paragraph: ParagraphDB) {

  }

  // display paragraph in router
  displayParagraphInRouter(paragraphID: number) {
    this.router.navigate(["gongmarkdown_go-" + "paragraph-display", paragraphID])
  }

  // set editor outlet
  setEditorRouterOutlet(paragraphID: number) {
    this.router.navigate([{
      outlets: {
        gongmarkdown_go_editor: ["gongmarkdown_go-" + "paragraph-detail", paragraphID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(paragraphID: number) {
    this.router.navigate([{
      outlets: {
        gongmarkdown_go_presentation: ["gongmarkdown_go-" + "paragraph-presentation", paragraphID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.paragraphs.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.paragraphs.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<ParagraphDB>()

      // reset all initial selection of paragraph that belong to paragraph
      for (let paragraph of this.initialSelection) {
        let index = paragraph[this.dialogData.ReversePointer as keyof ParagraphDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(paragraph)

      }

      // from selection, set paragraph that belong to paragraph
      for (let paragraph of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = paragraph[this.dialogData.ReversePointer as keyof ParagraphDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(paragraph)
      }


      // update all paragraph (only update selection & initial selection)
      for (let paragraph of toUpdate) {
        this.paragraphService.updateParagraph(paragraph)
          .subscribe(paragraph => {
            this.paragraphService.ParagraphServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, ParagraphDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedParagraph = new Set<number>()
      for (let paragraph of this.initialSelection) {
        if (this.selection.selected.includes(paragraph)) {
          // console.log("paragraph " + paragraph.Name + " is still selected")
        } else {
          console.log("paragraph " + paragraph.Name + " has been unselected")
          unselectedParagraph.add(paragraph.ID)
          console.log("is unselected " + unselectedParagraph.has(paragraph.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let paragraph = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as ParagraphDB
      if (unselectedParagraph.has(paragraph.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<ParagraphDB>) = new Array<ParagraphDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          paragraph => {
            if (!this.initialSelection.includes(paragraph)) {
              // console.log("paragraph " + paragraph.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + paragraph.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = paragraph.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = paragraph.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("paragraph " + paragraph.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<ParagraphDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}