import { Component, OnInit } from '@angular/core';
import { HttpService } from '../../services/http/http.service';
import { StateService } from '../../services/state/state.service';
import { iPost } from '../../interfaces/post.interface';
import { iHttpResponse } from '../../interfaces/http.interface';

@Component({
  selector: 'app-browse',
  templateUrl: './browse.component.html',
  styleUrls: ['./browse.component.scss']
})

export class BrowseComponent implements OnInit {
  public posts: iPost[];
  public render: boolean = false;
  public settings = {
    rowPadding: 0
  }

  constructor(
    private httpService: HttpService, 
    private stateService: StateService) {}

  ngOnInit(): void {
    this.stateService.getPostsState()
    .subscribe(response => this.posts = response);
    this.stateService.updatePostsState();
    this.render = true;
  }

}
