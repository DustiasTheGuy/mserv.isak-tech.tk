import { Component, OnInit } from '@angular/core';
import { StateService } from '../../services/state/state.service';

@Component({
  selector: 'app-sidenav',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.scss']
})
export class SidenavComponent implements OnInit {
  public menuOpen: boolean = false;

  constructor(public stateService: StateService) {}

  ngOnInit(): void {
    this.stateService.getSideNavState()
    .subscribe(state => this.menuOpen = state);
  }
}
