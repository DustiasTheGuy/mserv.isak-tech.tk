import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { NewComponent } from './pages/new/new.component';
import { PasteComponent } from './pages/paste/paste.component';
import { BrowseComponent } from './pages/browse/browse.component';
import { APIComponent } from './pages/api/api.component';

const routes: Routes = [
  { path: '', component: BrowseComponent },
  { path: 'new', component: NewComponent },
  { path: 'paste/:ID', component: PasteComponent },
  { path: 'api-instructions', component: APIComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
