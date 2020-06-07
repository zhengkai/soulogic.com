import { Component, OnInit } from '@angular/core';

import { NavService } from '../common/nav.service';
import { BlogService } from '../blog.service';

import { ApiOp, ApiPublic } from '../../api';
import { pb } from '../../pb/pb';

@Component({
	selector: 'app-index',
	templateUrl: './index.component.html',
	styleUrls: ['./index.component.scss']
})
export class IndexComponent implements OnInit {

	id = 0;
	text = '123123\nabcdfeaf';

	recent: pb.Item[];

	constructor(
		public ns: NavService,
		private bs: BlogService,
	) {
		ns.setCategory('index');
	}

	ngOnInit(): void {
		this.index();
	}

	async index() {
		this.recent = await this.bs.getRecent();
		console.log('index', this.recent);
	}

	async test() {

		const x = await ApiOp.echo(this.text || 'empty');
		console.log('echo re', x);

		const x2 = await ApiOp.itemEdit(this.id, this.text || 'empty');
		console.log('echo itemEdit', x2);

		this.recent = await this.bs.getRecent(true);
	}
}
