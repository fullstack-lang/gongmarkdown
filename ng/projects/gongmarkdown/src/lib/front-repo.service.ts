import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { MarkdownContentDB } from './markdowncontent-db'
import { MarkdownContentService } from './markdowncontent.service'

import { ParagraphDB } from './paragraph-db'
import { ParagraphService } from './paragraph.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  MarkdownContents_array = new Array<MarkdownContentDB>(); // array of repo instances
  MarkdownContents = new Map<number, MarkdownContentDB>(); // map of repo instances
  MarkdownContents_batch = new Map<number, MarkdownContentDB>(); // same but only in last GET (for finding repo instances to delete)
  Paragraphs_array = new Array<ParagraphDB>(); // array of repo instances
  Paragraphs = new Map<number, ParagraphDB>(); // map of repo instances
  Paragraphs_batch = new Map<number, ParagraphDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

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
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"
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

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private markdowncontentService: MarkdownContentService,
    private paragraphService: ParagraphService,
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
  observableFrontRepo: [ // insertion point sub template 
    Observable<MarkdownContentDB[]>,
    Observable<ParagraphDB[]>,
  ] = [ // insertion point sub template 
      this.markdowncontentService.getMarkdownContents(),
      this.paragraphService.getParagraphs(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            markdowncontents_,
            paragraphs_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var markdowncontents: MarkdownContentDB[]
            markdowncontents = markdowncontents_ as MarkdownContentDB[]
            var paragraphs: ParagraphDB[]
            paragraphs = paragraphs_ as ParagraphDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.MarkdownContents_array = markdowncontents

            // clear the map that counts MarkdownContent in the GET
            FrontRepoSingloton.MarkdownContents_batch.clear()

            markdowncontents.forEach(
              markdowncontent => {
                FrontRepoSingloton.MarkdownContents.set(markdowncontent.ID, markdowncontent)
                FrontRepoSingloton.MarkdownContents_batch.set(markdowncontent.ID, markdowncontent)
              }
            )

            // clear markdowncontents that are absent from the batch
            FrontRepoSingloton.MarkdownContents.forEach(
              markdowncontent => {
                if (FrontRepoSingloton.MarkdownContents_batch.get(markdowncontent.ID) == undefined) {
                  FrontRepoSingloton.MarkdownContents.delete(markdowncontent.ID)
                }
              }
            )

            // sort MarkdownContents_array array
            FrontRepoSingloton.MarkdownContents_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Paragraphs_array = paragraphs

            // clear the map that counts Paragraph in the GET
            FrontRepoSingloton.Paragraphs_batch.clear()

            paragraphs.forEach(
              paragraph => {
                FrontRepoSingloton.Paragraphs.set(paragraph.ID, paragraph)
                FrontRepoSingloton.Paragraphs_batch.set(paragraph.ID, paragraph)
              }
            )

            // clear paragraphs that are absent from the batch
            FrontRepoSingloton.Paragraphs.forEach(
              paragraph => {
                if (FrontRepoSingloton.Paragraphs_batch.get(paragraph.ID) == undefined) {
                  FrontRepoSingloton.Paragraphs.delete(paragraph.ID)
                }
              }
            )

            // sort Paragraphs_array array
            FrontRepoSingloton.Paragraphs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            markdowncontents.forEach(
              markdowncontent => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            paragraphs.forEach(
              paragraph => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // MarkdownContentPull performs a GET on MarkdownContent of the stack and redeem association pointers 
  MarkdownContentPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.markdowncontentService.getMarkdownContents()
        ]).subscribe(
          ([ // insertion point sub template 
            markdowncontents,
          ]) => {
            // init the array
            FrontRepoSingloton.MarkdownContents_array = markdowncontents

            // clear the map that counts MarkdownContent in the GET
            FrontRepoSingloton.MarkdownContents_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            markdowncontents.forEach(
              markdowncontent => {
                FrontRepoSingloton.MarkdownContents.set(markdowncontent.ID, markdowncontent)
                FrontRepoSingloton.MarkdownContents_batch.set(markdowncontent.ID, markdowncontent)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear markdowncontents that are absent from the GET
            FrontRepoSingloton.MarkdownContents.forEach(
              markdowncontent => {
                if (FrontRepoSingloton.MarkdownContents_batch.get(markdowncontent.ID) == undefined) {
                  FrontRepoSingloton.MarkdownContents.delete(markdowncontent.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // ParagraphPull performs a GET on Paragraph of the stack and redeem association pointers 
  ParagraphPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.paragraphService.getParagraphs()
        ]).subscribe(
          ([ // insertion point sub template 
            paragraphs,
          ]) => {
            // init the array
            FrontRepoSingloton.Paragraphs_array = paragraphs

            // clear the map that counts Paragraph in the GET
            FrontRepoSingloton.Paragraphs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            paragraphs.forEach(
              paragraph => {
                FrontRepoSingloton.Paragraphs.set(paragraph.ID, paragraph)
                FrontRepoSingloton.Paragraphs_batch.set(paragraph.ID, paragraph)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear paragraphs that are absent from the GET
            FrontRepoSingloton.Paragraphs.forEach(
              paragraph => {
                if (FrontRepoSingloton.Paragraphs_batch.get(paragraph.ID) == undefined) {
                  FrontRepoSingloton.Paragraphs.delete(paragraph.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getMarkdownContentUniqueID(id: number): number {
  return 31 * id
}
export function getParagraphUniqueID(id: number): number {
  return 37 * id
}