import { NgModule } from '@angular/core';
import { GongmarkdownspecificComponent } from './gongmarkdownspecific.component';
import { GongmarkdownspecificMarkdownComponent } from './gongmarkdownspecific-markdown/gongmarkdownspecific-markdown.component';



@NgModule({
  declarations: [
    GongmarkdownspecificComponent,
    GongmarkdownspecificMarkdownComponent
  ],
  imports: [
  ],
  exports: [
    GongmarkdownspecificComponent,
    GongmarkdownspecificMarkdownComponent
  ]
})
export class GongmarkdownspecificModule { }
