import { Component, OnInit } from '@angular/core';

import { NavService } from '../common/nav.service';

@Component({
	selector: 'app-index',
	templateUrl: './index.component.html',
	styleUrls: ['./index.component.scss']
})
export class IndexComponent implements OnInit {

	constructor(public ns: NavService) {
		ns.setCategory('index');
	}

	ngOnInit(): void {
	}
}
