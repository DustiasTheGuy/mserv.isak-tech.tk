import { Injectable } from '@angular/core';
import { Subject, Observable } from 'rxjs';
import { iPost } from '../../interfaces/post.interface';
import { HttpService } from '../http/http.service';
import { iHttpResponse } from '../../interfaces/http.interface';

@Injectable({
  providedIn: 'root'
})

export class StateService {
  public sideNavSubject = new Subject<boolean>();
  public postsSubject = new Subject<iPost[]>();


  constructor(private httpService: HttpService) {}

  public updateSideNavState(newState: boolean): void {
    this.sideNavSubject.next(newState);
  }

  public getSideNavState(): Observable<boolean> {
    return this.sideNavSubject.asObservable();
  }

  public updatePostsState() {
    this.httpService.posts()
    .subscribe((response: iHttpResponse) => response.success ? 
    this.postsSubject.next(response.data) : this.postsSubject.next([]));
  }

  public getPostsState(): Observable<iPost[]> {
    return this.postsSubject.asObservable();
  }
}
