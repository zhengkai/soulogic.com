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

	constructor(public ns: NavService) {
		ns.setCategory('thoughts');
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
		list.forEach(v => {
			const s = Api.markdown(v);
			this.short.push({
				time: Date.now(),
				content: s,
			});
		});

		list.reverse().forEach(v => {
			const s = Api.markdown(v);
			this.long.push({
				time: Date.now(),
				content: s,
			});
		});
	}
}
