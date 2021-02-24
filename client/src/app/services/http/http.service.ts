import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { iHttpResponse } from '../../interfaces/http.interface';
import { iPost } from '../../interfaces/post.interface';

@Injectable({
  providedIn: 'root'
})

export class HttpService {
  private environment: boolean = true; // production?
  private serverAddr: string;

  constructor(private httpClient: HttpClient) {
    this.serverAddr = this.environment ? 'https://paste.isak-tech.tk' : 'http://localhost:8082'
  }

  public submit(data): Observable<iHttpResponse> {
    return this.httpClient.post<iHttpResponse>(this.serverAddr + '/api/new', data);
  }

  public posts(): Observable<iHttpResponse> {
    return this.httpClient.get<iHttpResponse>(this.serverAddr + '/api/posts');
  }

  public post(ID: number): Observable<iHttpResponse> {
    return this.httpClient.get<iHttpResponse>(this.serverAddr + '/api/post/' + ID);
  }

  public signIn(password: string): Observable<iHttpResponse> {
    return this.httpClient.post<iHttpResponse>(this.serverAddr + '/api/sign-in', { password })
  }

  public pageinate(page: number, limit: number): Observable<iHttpResponse> {
    return this.httpClient.get<iHttpResponse>(this.serverAddr + '/api/paginate/' + page + '/' + limit)
  }
}
