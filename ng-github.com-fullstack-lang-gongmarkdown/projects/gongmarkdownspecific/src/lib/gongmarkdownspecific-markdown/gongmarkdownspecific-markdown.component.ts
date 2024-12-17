import { Component, Input, OnInit, importProvidersFrom } from '@angular/core';
import * as gongmarkdown from '../../../../gongmarkdown/src/public-api'
import { Observable, timer } from 'rxjs';
import { CommonModule } from '@angular/common';
import { MarkdownModule } from 'ngx-markdown';
import { HttpClientModule } from '@angular/common/http';

@Component({
  selector: 'lib-gongmarkdownspecific-markdown',
  templateUrl: './gongmarkdownspecific-markdown.component.html',
  styleUrls: ['./gongmarkdownspecific-markdown.component.css'],
  standalone: true,
  imports: [CommonModule,
    MarkdownModule, HttpClientModule  ],
    providers: [],

})
export class GongmarkdownspecificMarkdownComponent implements OnInit {

  @Input() GONG__StackPath: string = ""
  frontRepo: gongmarkdown.FrontRepo | undefined

  markdownContent = ``

  constructor(
    private frontRepoService: gongmarkdown.FrontRepoService,
  ) { }

  ngOnInit(): void {
    this.frontRepoService.connectToWebSocket(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.refresh()
      }
    )
  }

  refresh(): void {
    if (this.frontRepo === undefined) {
      return
    }
    if (this.frontRepo.getFrontArray(gongmarkdown.MarkdownContent.GONGSTRUCT_NAME).length == 1) {
      let markdownContentObj = this.frontRepo.getFrontArray<gongmarkdown.MarkdownContent>(gongmarkdown.MarkdownContent.GONGSTRUCT_NAME)[0]
      this.markdownContent = markdownContentObj.Content
    }
  }
}