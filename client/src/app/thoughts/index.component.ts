import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Api } from '../../api';
import { pb } from '../../pb/pb';

import { NavService } from '../common/nav.service';

@Component({
	selector: 'app-thoughts',
	templateUrl: './index.component.html',
	styleUrls: ['./index.component.scss'],
})
export class ThoughtsComponent implements OnInit {

	isLoad = false;

	short: pb.Item[] = [];
	long: pb.Item[] = [];

	post = [];

	constructor(public ns: NavService) {
		ns.setCategory('thoughts');

		this.post.push({
			time: '2019-06-08',
			title: '智慧和道德观念',
		});
		this.post.push({
			time: '2019-06-07',
			title: '关于功利主义',
		});
		this.post.push({
			time: '2018-01-14',
			title: '恶魔之名',
		});
	}

	ngOnInit(): void {

		if (this.isLoad) {
			return;
		}
		this.isLoad = true;

		this.load();
	}

	async load() {

		const list = await Api.fetch();
		let ID = 10;

		list.forEach(v => {
			ID++;
			const s = Api.markdown(v);

			const o = {} as pb.Item;
			o.ID = ID;
			o.tsCreate = Date.now();
			o.revision = {
				raw: s,
			} as pb.Revision,

			this.short.push(o);
		});

		list.reverse().forEach(v => {
			ID++;
			const s = Api.markdown(v);

			const o = {} as pb.Item;
			o.ID = ID;
			o.tsCreate = Date.now();
			o.revision = {
				raw: s,
			} as pb.Revision,

			this.long.push(o);
		});
	}
}
