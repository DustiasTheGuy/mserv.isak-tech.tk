import { Component, OnInit } from '@angular/core';
import { HttpService } from '../../services/http/http.service';
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

  constructor(private httpService: HttpService) { }

  ngOnInit(): void {
    this.httpService.posts()
    .subscribe((response: iHttpResponse) => 
    this.posts = response.success ? response.data : [],
    (err) => console.error(err), () => this.render = true);
  }

}
