import { Component, OnInit, Output, EventEmitter, Input } from '@angular/core';
import { StateService } from '../../services/state/state.service';

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.scss']
})
export class NavigationComponent implements OnInit {
  public menuOpen: boolean;

  constructor(public stateService: StateService) {}

  ngOnInit(): void {
    this.stateService.getSideNavState()
    .subscribe(state => this.menuOpen = state);
  }

}
