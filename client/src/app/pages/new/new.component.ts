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
  public tag: string;
  public msg: string;
  public formFields = {
    title: '',
    body: '',
    tags: []
  }

  constructor(
    private router: Router,
    private httpService: HttpService) {}

  ngOnInit(): void {
  }

  submit(): void {
    if(this.formFields.tags.length <= 0)
    return alert('You must enter at least one tag')

    if(this.formFields.body.length <= 1)
    return alert('You must enter a body');

    this.formFields.tags = this.formFields.tags.slice(0, 5);
    this.formFields.title = this.formFields.title || 'Empty';
    this.httpService.submit(this.formFields)
    .subscribe((response: iHttpResponse) => this.onResponse(response));
  }

  onResponse(response: iHttpResponse): void {
    response.success ? this.router.navigate(['/paste/' + response.data]) :
    alert(response.message);
  }

  appendTag(tag: string): boolean {
    if(this.formFields.tags.length >= 5) {
      window.location.reload(); 
      return false; 
    }
    
    this.formFields.tags.push(tag);
    this.tag = undefined;
    return true;
  }
}
