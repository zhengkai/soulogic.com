import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { AppRoutingModule } from './common/routing.module';
import { BootstrapComponent } from './common/bootstrap.component';
import { IndexComponent } from './index/index.component';
import { NavComponent } from './common/nav.component';
import { AboutMeComponent } from './about-me/about-me.component';
import { ThoughtsComponent } from './thoughts/index.component';
import { Error451Component } from './error/error451.component';
import { ThoughtsSideComponent } from './thoughts/side.component';
import { ThoughtsItemComponent } from './thoughts/item.component';
import { ThoughtsViewComponent } from './thoughts/view.component';
import { ItemComponent } from './item/item.component';

@NgModule({
	declarations: [
		BootstrapComponent,
		IndexComponent,
		NavComponent,
		AboutMeComponent,
		ThoughtsComponent,
		Error451Component,
		ThoughtsSideComponent,
		ThoughtsItemComponent,
		ThoughtsViewComponent,
		ItemComponent,
	],
	imports: [
		BrowserModule,
		AppRoutingModule,
		FormsModule,
	],
	providers: [],
	bootstrap: [
		BootstrapComponent,
	],
})
export class AppModule { }
