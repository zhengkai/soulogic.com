import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './common/routing.module';
import { BootstrapComponent } from './common/bootstrap.component';
import { IndexComponent } from './index/index.component';
import { NavComponent } from './common/nav.component';

@NgModule({
	declarations: [
		BootstrapComponent,
		IndexComponent,
		NavComponent,
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
