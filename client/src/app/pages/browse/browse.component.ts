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

  public paginationData = {
    page: 0,
    limit: 25,
    pages: 0,
    count: 0
  }

  constructor(
    private httpService: HttpService, 
    private stateService: StateService) {}

  ngOnInit(): void {
    this.paginate(this.paginationData.page, this.paginationData.limit);

    this.render = true;
    setTimeout(() => console.log(this.posts), 3000)
  }

  paginate(page, limit) {

    this.httpService.pageinate(page, limit)
    .subscribe((response: iHttpResponse) => {
      if(response.success) {
        this.paginationData.page = page;
        this.paginationData.pages = response.data.pages;
        this.paginationData.count = response.data.count;
        this.posts = response.data.posts; 
      }   
    });
  }
}
