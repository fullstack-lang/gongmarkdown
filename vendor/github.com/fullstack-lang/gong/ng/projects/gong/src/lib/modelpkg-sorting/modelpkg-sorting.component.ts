// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { ModelPkgDB } from '../modelpkg-db'
import { ModelPkgService } from '../modelpkg.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-modelpkg-sorting',
  templateUrl: './modelpkg-sorting.component.html',
  styleUrls: ['./modelpkg-sorting.component.css']
})
export class ModelPkgSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of ModelPkg instances that are in the association
  associatedModelPkgs = new Array<ModelPkgDB>();

  constructor(
    private modelpkgService: ModelPkgService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of modelpkg instances
    public dialogRef: MatDialogRef<ModelPkgSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getModelPkgs()
  }

  getModelPkgs(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let modelpkg of this.frontRepo.ModelPkgs_array) {
          let ID = this.dialogData.ID
          let revPointerID = modelpkg[this.dialogData.ReversePointer as keyof ModelPkgDB] as unknown as NullInt64
          let revPointerID_Index = modelpkg[this.dialogData.ReversePointer + "_Index" as keyof ModelPkgDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedModelPkgs.push(modelpkg)
          }
        }

        // sort associated modelpkg according to order
        this.associatedModelPkgs.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedModelPkgs, event.previousIndex, event.currentIndex);

    // set the order of ModelPkg instances
    let index = 0

    for (let modelpkg of this.associatedModelPkgs) {
      let revPointerID_Index = modelpkg[this.dialogData.ReversePointer + "_Index" as keyof ModelPkgDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedModelPkgs.forEach(
      modelpkg => {
        this.modelpkgService.updateModelPkg(modelpkg)
          .subscribe(modelpkg => {
            this.modelpkgService.ModelPkgServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer +' done');
  }
}