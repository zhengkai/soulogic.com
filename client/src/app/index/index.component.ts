import { Component, OnInit } from '@angular/core';

import { NavService } from '../common/nav.service';

import { ApiOp } from '../../api';
import { pb } from '../../pb/pb';

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

	async test() {

		const x = await ApiOp.echo('abc ' + Date.now());
		console.log('echo re', x);
	}
}
