import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { NewComponent } from './pages/new/new.component';
import { PasteComponent } from './pages/paste/paste.component';
import { BrowseComponent } from './pages/browse/browse.component';

const routes: Routes = [
  { path: '', component: BrowseComponent },
  { path: 'new', component: NewComponent },
  { path: 'paste/:id', component: PasteComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
