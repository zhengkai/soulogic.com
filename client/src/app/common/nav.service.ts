import { Injectable } from '@angular/core';

@Injectable({
	providedIn: 'root'
})
export class NavService {

	enable = true;

	category = [
		{
			name: 'index',
			icon: 'home',
			link: '/',
		},
		{
			name: 'thoughts',
			icon: 'speaker_notes',
			link: '/thoughts',
		},
		{
			name: 'about me',
			icon: 'person_pin_circle',
			link: '/about-me',
		},
	];

	categorySelect = '';

	constructor() {
		this.categorySelect = this.category[0].name;
	}

	setCategory(name: string) {
		this.categorySelect = name;
	}

	setEnable(enable: boolean) {
		this.enable = !!enable;
	}
}
