import { Component, Input, OnInit, importProvidersFrom } from '@angular/core';
import * as gongmarkdown from '../../../../gongmarkdown/src/public-api'
import { Observable, timer } from 'rxjs';
import { CommonModule } from '@angular/common';
import { MarkdownModule } from 'ngx-markdown';
import { HttpClientModule } from '@angular/common/http';

@Component({
  selector: 'lib-textarea',
  standalone: true,
  imports: [CommonModule, MarkdownModule, HttpClientModule],
  templateUrl: './textarea.component.html',
  styleUrl: './textarea.component.css'
})
export class TextareaComponent {
  @Input() GONG__StackPath: string = ""
  frontRepo: gongmarkdown.FrontRepo | undefined

  content = ``

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
    if (this.frontRepo.getFrontArray(gongmarkdown.Content.GONGSTRUCT_NAME).length == 1) {
      let content = this.frontRepo.getFrontArray<gongmarkdown.Content>(gongmarkdown.Content.GONGSTRUCT_NAME)[0]
      this.content = content.Content
    }
  }
}
