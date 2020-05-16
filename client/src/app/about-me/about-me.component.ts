import { Component, OnInit } from '@angular/core';

import { NavService } from '../common/nav.service';

@Component({
	selector: 'app-about-me',
	templateUrl: './about-me.component.html',
	styleUrls: ['./about-me.component.scss']
})
export class AboutMeComponent implements OnInit {

	constructor(public ns: NavService) {
		ns.setCategory('about me');
	}

	ngOnInit(): void {
	}
}
