import { Component } from '@angular/core';
import { StateService } from './services/state/state.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent {
  public menuOpen: boolean = false;

  constructor(private stateService: StateService) {
    this.stateService.getSideNavState()
    .subscribe(state => this.menuOpen = state);
  }
}
