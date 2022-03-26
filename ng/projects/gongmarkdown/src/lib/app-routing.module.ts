import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { ElementsTableComponent } from './elements-table/elements-table.component'
import { ElementDetailComponent } from './element-detail/element-detail.component'
import { ElementPresentationComponent } from './element-presentation/element-presentation.component'

import { MarkdownContentsTableComponent } from './markdowncontents-table/markdowncontents-table.component'
import { MarkdownContentDetailComponent } from './markdowncontent-detail/markdowncontent-detail.component'
import { MarkdownContentPresentationComponent } from './markdowncontent-presentation/markdowncontent-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'gongmarkdown_go-elements', component: ElementsTableComponent, outlet: 'gongmarkdown_go_table' },
	{ path: 'gongmarkdown_go-element-adder', component: ElementDetailComponent, outlet: 'gongmarkdown_go_editor' },
	{ path: 'gongmarkdown_go-element-adder/:id/:originStruct/:originStructFieldName', component: ElementDetailComponent, outlet: 'gongmarkdown_go_editor' },
	{ path: 'gongmarkdown_go-element-detail/:id', component: ElementDetailComponent, outlet: 'gongmarkdown_go_editor' },
	{ path: 'gongmarkdown_go-element-presentation/:id', component: ElementPresentationComponent, outlet: 'gongmarkdown_go_presentation' },
	{ path: 'gongmarkdown_go-element-presentation-special/:id', component: ElementPresentationComponent, outlet: 'gongmarkdown_goelementpres' },

	{ path: 'gongmarkdown_go-markdowncontents', component: MarkdownContentsTableComponent, outlet: 'gongmarkdown_go_table' },
	{ path: 'gongmarkdown_go-markdowncontent-adder', component: MarkdownContentDetailComponent, outlet: 'gongmarkdown_go_editor' },
	{ path: 'gongmarkdown_go-markdowncontent-adder/:id/:originStruct/:originStructFieldName', component: MarkdownContentDetailComponent, outlet: 'gongmarkdown_go_editor' },
	{ path: 'gongmarkdown_go-markdowncontent-detail/:id', component: MarkdownContentDetailComponent, outlet: 'gongmarkdown_go_editor' },
	{ path: 'gongmarkdown_go-markdowncontent-presentation/:id', component: MarkdownContentPresentationComponent, outlet: 'gongmarkdown_go_presentation' },
	{ path: 'gongmarkdown_go-markdowncontent-presentation-special/:id', component: MarkdownContentPresentationComponent, outlet: 'gongmarkdown_gomarkdowncontentpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
