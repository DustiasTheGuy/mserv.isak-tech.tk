import { Injectable } from '@angular/core';
import { Subject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class StateService {
  public sideNavSubject = new Subject<boolean>();

  constructor() {}

  public updateSideNavState(newState: boolean): void {
    this.sideNavSubject.next(newState);
  }

  public getSideNavState(): Observable<boolean> {
    return this.sideNavSubject.asObservable();
  }
}
