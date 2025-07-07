import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

// for angular & angular material
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatRadioModule } from '@angular/material/radio';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import * as gongmarkdown from '../../projects/gongmarkdown/src/public-api'

import { GongmarkdownSpecificComponent, TextareaComponent } from '../../projects/gongmarkdownspecific/src/public-api'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatRadioModule,
    MatButtonModule,
    MatIconModule,
    AngularSplitModule,

    SplitSpecificComponent,

    TextareaComponent,
    GongmarkdownSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  gongmarkdown = 'Gongmarkdown'
  probe = 'Gongmarkdown Data/Model'
  view = this.gongmarkdown

  views: string[] = [this.gongmarkdown, this.probe];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "gongmarkdown"
  StackType = gongmarkdown.StackType

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
