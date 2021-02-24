import { Component, OnInit } from '@angular/core';
import { routes } from './routes';
import { Router } from '@angular/router';
import { HttpService } from '../../services/http/http.service';

@Component({
  selector: 'app-api',
  templateUrl: './api.component.html',
  styleUrls: ['./api.component.scss']
})
export class APIComponent implements OnInit {
  public routes;
  public render: boolean = true;
  public password: string;
  
  constructor(
    private router: Router, 
    private httpService: HttpService) {
    this.routes = routes;
  }

  ngOnInit(): void {

  }

  signIn(password) {
    this.httpService.signIn(password)
    .subscribe(response => response.success ?
    this.render = true : 
    this.router.navigate(["/"]));
  }
}