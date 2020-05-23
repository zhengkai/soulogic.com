import { Component, OnInit } from '@angular/core';
import { Api } from '../../api';

import { NavService } from '../common/nav.service';

@Component({
	selector: 'app-thoughts-item',
	templateUrl: './item.component.html',
	styleUrls: ['./item.component.scss']
})
export class ThoughtsItemComponent implements OnInit {

	isLoad = false;

	item = {
		id: 0,
		time: 0,
		content: '',
	};

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
		const id = 321;

		const one = list?.[7];
		if (!one) {
			return;
		}

		const s = Api.markdown(one);
		this.item = {
			id,
			time: Date.now(),
			content: s,
		};
	}
}
