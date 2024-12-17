import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent {

  // choices for the top radio button
  view = 'Markdown view'
  markdown = 'Markdown view'
  text = 'Text view'
  default = 'Default view'
  diagrams = 'Diagrams view'
  meta = 'Meta view'
  views: string[] = [this.markdown, this.text, this.default, this.diagrams, this.meta];
}
