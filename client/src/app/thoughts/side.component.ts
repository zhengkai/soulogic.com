import { Component, OnInit } from '@angular/core';

@Component({
	selector: 'app-thoughts-side',
	templateUrl: './side.component.html',
	styleUrls: ['./side.component.scss']
})
export class ThoughtsSideComponent implements OnInit {

	list = [];

	constructor() {

		for (let i = 0; i < 12; i++) {
			this.list.push({
				year: 2020 - i,
				count: Math.floor(Math.random() * 800),
			});
		}
	}

	ngOnInit(): void {

	}
}
