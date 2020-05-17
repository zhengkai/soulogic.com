import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { IndexComponent } from '../index/index.component';
import { AboutMeComponent } from '../about-me/about-me.component';
import { ThoughtsComponent } from '../thoughts/thoughts.component';
import { Error451Component } from '../error/error451.component';

const routes: Routes = [
	{ path: '', component: IndexComponent },
	{ path: 'about-me', component: AboutMeComponent },
	{ path: 'thoughts', component: ThoughtsComponent },
	{ path: 'error-451', component: Error451Component },
	{ path: '**', redirectTo: '/' },
];

@NgModule({
	imports: [
		RouterModule.forRoot(routes, {
			scrollPositionRestoration: 'enabled',
		}),
	],
	exports: [
		RouterModule,
	],
})
export class AppRoutingModule { }
