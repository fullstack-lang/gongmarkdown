import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ParagraphDB } from '../paragraph-db'
import { ParagraphService } from '../paragraph.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface paragraphDummyElement {
}

const ELEMENT_DATA: paragraphDummyElement[] = [
];

@Component({
	selector: 'app-paragraph-presentation',
	templateUrl: './paragraph-presentation.component.html',
	styleUrls: ['./paragraph-presentation.component.css'],
})
export class ParagraphPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	paragraph: ParagraphDB = new (ParagraphDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private paragraphService: ParagraphService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getParagraph();

		// observable for changes in 
		this.paragraphService.ParagraphServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getParagraph()
				}
			}
		)
	}

	getParagraph(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.paragraph = this.frontRepo.Paragraphs.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				gongmarkdown_go_presentation: ["gongmarkdown_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				gongmarkdown_go_editor: ["gongmarkdown_go-" + "paragraph-detail", ID]
			}
		}]);
	}
}
