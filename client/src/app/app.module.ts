import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { NavigationComponent } from './components/navigation/navigation.component';
import { SidenavComponent } from './components/sidenav/sidenav.component';
import { NewComponent } from './pages/new/new.component';
import { PasteComponent } from './pages/paste/paste.component';
import { BrowseComponent } from './pages/browse/browse.component';


@NgModule({
  declarations: [
    AppComponent,
    NavigationComponent,
    SidenavComponent,
    NewComponent,
    PasteComponent,
    BrowseComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [ AppComponent ]
})

export class AppModule { }
 