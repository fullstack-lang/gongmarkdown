import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { ParagraphsTableComponent } from './paragraphs-table/paragraphs-table.component'
import { ParagraphDetailComponent } from './paragraph-detail/paragraph-detail.component'
import { ParagraphPresentationComponent } from './paragraph-presentation/paragraph-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'gongmarkdown_go-paragraphs', component: ParagraphsTableComponent, outlet: 'gongmarkdown_go_table' },
	{ path: 'gongmarkdown_go-paragraph-adder', component: ParagraphDetailComponent, outlet: 'gongmarkdown_go_editor' },
	{ path: 'gongmarkdown_go-paragraph-adder/:id/:originStruct/:originStructFieldName', component: ParagraphDetailComponent, outlet: 'gongmarkdown_go_editor' },
	{ path: 'gongmarkdown_go-paragraph-detail/:id', component: ParagraphDetailComponent, outlet: 'gongmarkdown_go_editor' },
	{ path: 'gongmarkdown_go-paragraph-presentation/:id', component: ParagraphPresentationComponent, outlet: 'gongmarkdown_go_presentation' },
	{ path: 'gongmarkdown_go-paragraph-presentation-special/:id', component: ParagraphPresentationComponent, outlet: 'gongmarkdown_goparagraphpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
