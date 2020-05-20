import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Api } from '../../api';

import { NavService } from '../common/nav.service';

@Component({
	selector: 'app-thoughts',
	templateUrl: './thoughts.component.html',
	styleUrls: ['./thoughts.component.scss'],
})
export class ThoughtsComponent implements OnInit {

	isLoad = false;

	short = [];
	long = [];

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
		let id = 0;
		list.forEach(v => {
			id++;
			const s = Api.markdown(v);
			this.short.push({
				id,
				time: Date.now(),
				content: s,
			});
		});

		list.reverse().forEach(v => {
			id++;
			const s = Api.markdown(v);
			this.long.push({
				id,
				time: Date.now(),
				content: s,
			});
		});
	}
}
