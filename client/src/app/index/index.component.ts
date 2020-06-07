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

	id = 0;
	text = '123123\nabcdfeaf';

	constructor(public ns: NavService) {
		ns.setCategory('index');
	}

	ngOnInit(): void {
	}

	async test() {

		console.log('this.text', this.text);

		const x = await ApiOp.echo(this.text || 'empty');
		console.log('echo re', x);

		const x2 = await ApiOp.itemEdit(this.id, this.text || 'empty');
		console.log('echo itemEdit', x2);
	}
}
