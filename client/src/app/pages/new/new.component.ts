import { Component, OnInit } from '@angular/core';
import { HttpService } from '../../services/http/http.service';
import { iHttpResponse } from '../../interfaces/http.interface';
import { Router } from '@angular/router';

@Component({
  selector: 'app-new',
  templateUrl: './new.component.html',
  styleUrls: ['./new.component.scss']
})

export class NewComponent implements OnInit {
  public body: string;

  constructor(
    private router: Router,
    private httpService: HttpService) {}

  ngOnInit(): void {
  }

  submit() {
    this.httpService.submit({ _id: 0, body: this.body, date: null })
    .subscribe((response: iHttpResponse) => this.onResponse(response));
  }

  onResponse(response: iHttpResponse) {
    response.success ? this.router.navigate(['/paste/' + response.data])
    : console.log("Hmm")
  }
}
