import { Component, Input, OnInit, importProvidersFrom } from '@angular/core';
import * as gongmarkdown from '../../../../gongmarkdown/src/public-api'
import { Observable, timer } from 'rxjs';
import { CommonModule } from '@angular/common';
import { MarkdownModule, provideMarkdown } from 'ngx-markdown';

@Component({
  selector: 'lib-gongmarkdown-specific',
  standalone: true,
  imports: [CommonModule, MarkdownModule],
  providers: [provideMarkdown()],
  templateUrl: './gongmarkdown-specific.component.html',
  styleUrl: './gongmarkdown-specific.component.css'
})
export class GongmarkdownSpecificComponent {
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
