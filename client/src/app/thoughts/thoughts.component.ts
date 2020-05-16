import { Component, OnInit } from '@angular/core';

import { NavService } from '../common/nav.service';

@Component({
	selector: 'app-thoughts',
	templateUrl: './thoughts.component.html',
	styleUrls: ['./thoughts.component.scss']
})
export class ThoughtsComponent implements OnInit {

	constructor(public ns: NavService) {
		ns.setCategory('thoughts');
	}

	ngOnInit(): void {
	}
}
