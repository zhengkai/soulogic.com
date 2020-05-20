import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './common/routing.module';
import { BootstrapComponent } from './common/bootstrap.component';
import { IndexComponent } from './index/index.component';
import { NavComponent } from './common/nav.component';
import { AboutMeComponent } from './about-me/about-me.component';
import { ThoughtsComponent } from './thoughts/thoughts.component';
import { Error451Component } from './error/error451.component';
import { ThoughtsSideComponent } from './thoughts/side.component';

@NgModule({
	declarations: [
		BootstrapComponent,
		IndexComponent,
		NavComponent,
		AboutMeComponent,
		ThoughtsComponent,
		Error451Component,
		ThoughtsSideComponent,
	],
	imports: [
		BrowserModule,
		AppRoutingModule,
	],
	providers: [],
	bootstrap: [
		BootstrapComponent,
	],
})
export class AppModule { }
