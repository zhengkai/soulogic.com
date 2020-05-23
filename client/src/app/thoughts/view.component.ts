import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Api } from '../../api';
import { pb } from '../../pb/pb';

import { NavService } from '../common/nav.service';

@Component({
	selector: 'app-thoughts-view',
	templateUrl: './view.component.html',
	styleUrls: ['./view.component.scss']
})
export class ThoughtsViewComponent implements OnInit {

	isLoad = false;

	id = 0;

	item = { } as pb.Post;

	constructor(
		public ns: NavService,
		private route: ActivatedRoute,
	) {
		ns.setCategory('thoughts');

		route.params.subscribe(() => {
			this.ngOnInit();
		});
	}

	ngOnInit(): void {

		if (this.id) {
			return;
		}

		const params = this.route.snapshot.params;
		this.id = parseInt(params.id, 10) || 0;
		console.log(this.id);
		if (!this.id) {
			return;
		}

		if (this.isLoad) {
			return;
		}
		this.isLoad = true;

		this.load();
	}

	async load() {

		const list = await Api.fetch();

		const one = list?.[7];
		if (!one) {
			return;
		}

		this.item.ID = 321;
		this.item.tsCreate = Date.now(),
		this.item.raw = one;
	}
}
