import { Component, OnInit } from '@angular/core';
import { HttpService } from '../../services/http/http.service';

@Component({
  selector: 'app-new',
  templateUrl: './new.component.html',
  styleUrls: ['./new.component.scss']
})
export class NewComponent implements OnInit {
  public body: string;

  constructor(private httpService: HttpService) { }

  ngOnInit(): void {
  }

  submit() {
    this.httpService.submit({ 
      _id: 0,
      body: this.body,
      date: null 
    }).subscribe(response => console.log(response));
  }
}
