import { Component } from '@angular/core';
import { NavService } from '../common/nav.service';

@Component({
  selector: 'app-error451',
  templateUrl: './error451.component.html',
  styleUrls: ['./error451.component.scss']
})
export class Error451Component {

	constructor(public ns: NavService) {
		ns.setCategory(null);
		ns.setCategoryDisplay(false);
	}
}
