import { Component, Input, OnChanges } from '@angular/core';
import { pb } from '../../pb/pb';

@Component({
	selector: 'app-item',
	templateUrl: './item.component.html',
	styleUrls: ['./item.component.scss']
})
export class ItemComponent implements OnChanges {

	@Input() data: pb.Item;

	constructor() { }

	ngOnChanges(): void {
	}
}
