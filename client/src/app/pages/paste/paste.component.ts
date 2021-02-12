import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { iPost } from '../../interfaces/post.interface';
import { HttpService } from '../../services/http/http.service';
import { iHttpResponse } from '../../interfaces/http.interface';

@Component({
  selector: 'app-paste',
  templateUrl: './paste.component.html',
  styleUrls: ['./paste.component.scss']
})
export class PasteComponent implements OnInit {
  public post: iPost;

  constructor(
    private httpService: HttpService,
    private activatedRoute: ActivatedRoute, 
    private router: Router
    ) {}

  ngOnInit(): void {
    this.activatedRoute.params
    .subscribe(routerData => 
    isNaN(parseInt(routerData.ID)) ? 
    this.router.navigate(['/']) : 
    this.makeAPIcall(parseInt(routerData.ID)));
  }


  makeAPIcall(ID: number) {
    //console.log('Calling API with ID ' + ID)
    this.httpService.post(ID)
    .subscribe((response: iHttpResponse) => 
    this.post = response.success ? response.data : { _id: 1337, body: 'what the hell you up to?', date: new Date()});
  }
}
