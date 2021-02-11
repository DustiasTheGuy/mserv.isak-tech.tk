import { Component, OnInit } from '@angular/core';
import { HttpService } from '../../services/http/http.service';
import { iHttpResponse } from '../../interfaces/http.interface';

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
    this.httpService.submit({ _id: 0, body: this.body, date: null })
    .subscribe((response: iHttpResponse) => 
    response.success ? this.body = undefined : console.log(response));
  }
}
