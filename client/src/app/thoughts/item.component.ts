import { Component, Input, OnChanges } from '@angular/core';
import { Api } from '../../api';
import { pb } from '../../pb/pb';

@Component({
	selector: 'app-thoughts-item',
	templateUrl: './item.component.html',
	styleUrls: ['./item.component.scss']
})
export class ThoughtsItemComponent implements OnChanges {

	@Input() data: pb.Item;
	@Input() link: boolean;

	html = '';

	isLoad = false;

	constructor() { }

	ngOnChanges(): void {

		const raw = this.data?.revision?.raw;
		if (!raw) {
			return;
		}

		this.html = Api.markdown(raw);
		// console.log(this.html);
	}
}
