import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class HttpService {
  private serverAddr: string = 'http://localhost:8082';

  constructor(private httpClient: HttpClient) {}

  public submit(data): Observable<any> {
    return this.httpClient.post(this.serverAddr + '/api/new', data);
  }
}
