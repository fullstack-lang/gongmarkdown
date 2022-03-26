import { NgModule } from '@angular/core';
import { GongmarkdownspecificComponent } from './gongmarkdownspecific.component';
import { GongmarkdownspecificMarkdownComponent } from './gongmarkdownspecific-markdown/gongmarkdownspecific-markdown.component';

import { MarkdownModule, MarkedOptions } from 'ngx-markdown';

@NgModule({
  declarations: [
    GongmarkdownspecificComponent,
    GongmarkdownspecificMarkdownComponent
  ],
  imports: [
    MarkdownModule.forRoot()
  ],
  exports: [
    GongmarkdownspecificComponent,
    GongmarkdownspecificMarkdownComponent
  ]
})
export class GongmarkdownspecificModule { }
