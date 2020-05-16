import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './common/routing.module';
import { BootstrapComponent } from './common/bootstrap.component';
import { IndexComponent } from './index/index.component';
import { NavComponent } from './common/nav.component';
import { AboutMeComponent } from './about-me/about-me.component';
import { ThoughtsComponent } from './thoughts/thoughts.component';

@NgModule({
	declarations: [
		BootstrapComponent,
		IndexComponent,
		NavComponent,
		AboutMeComponent,
		ThoughtsComponent,
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
