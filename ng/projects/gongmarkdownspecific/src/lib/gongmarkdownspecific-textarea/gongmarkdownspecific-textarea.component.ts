import { Component, OnInit } from '@angular/core';

import * as gongmarkdown from 'gongmarkdown'

@Component({
  selector: 'lib-gongmarkdownspecific-textarea',
  templateUrl: './gongmarkdownspecific-textarea.component.html',
  styleUrls: ['./gongmarkdownspecific-textarea.component.css']
})
export class GongmarkdownspecificTextareaComponent implements OnInit {

  markdownContent: string = ""

  markdownContentDB: gongmarkdown.MarkdownContentDB | undefined

  constructor(
    private gongmarkdownMarkdownContentService: gongmarkdown.MarkdownContentService
  ) { }

  ngOnInit(): void {

    // get the singloton
    this.gongmarkdownMarkdownContentService.getMarkdownContents().subscribe(
      markdownContentDBs => {
        if (markdownContentDBs.length == 1) {
          this.markdownContentDB = markdownContentDBs[0]
          this.markdownContent = this.markdownContentDB.Content
        } else {
          console.log("wrong number of markdown content " + markdownContentDBs.length)
        }
      }

    )

  }

}
