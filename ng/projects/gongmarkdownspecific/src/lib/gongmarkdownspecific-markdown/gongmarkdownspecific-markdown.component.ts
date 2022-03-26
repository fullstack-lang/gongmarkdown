import { Component, OnInit } from '@angular/core';

import * as gongmarkdown from 'gongmarkdown'

@Component({
  selector: 'lib-gongmarkdownspecific-markdown',
  templateUrl: './gongmarkdownspecific-markdown.component.html',
  styleUrls: ['./gongmarkdownspecific-markdown.component.css']
})
export class GongmarkdownspecificMarkdownComponent implements OnInit {

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
