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

	isLoad = false;

	constructor() { }

	ngOnChanges(): void {

	}
}
