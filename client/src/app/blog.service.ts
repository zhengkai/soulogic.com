import { Injectable } from '@angular/core';

import { ApiOp, ApiPublic } from '../api';
import { pb } from '../pb/pb';

@Injectable({
	providedIn: 'root'
})
export class BlogService {

	recent: pb.Item[];

	constructor() {
	}

	async getRecent(force = false) {
		if (force || !this.recent?.length) {
			const x = await ApiPublic.blogIndex();
			if (x?.list?.length) {
				this.recent = x.list;
			}
		}
		return this.recent;
	}
}
