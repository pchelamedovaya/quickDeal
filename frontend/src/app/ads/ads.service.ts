import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable, Subject, tap} from 'rxjs';

import {AdResponse, CreateAdRequest} from './ad.models';

@Injectable({providedIn: 'root'})
export class AdsService {
  constructor(private readonly http: HttpClient) {}

  private readonly createdAdSubject = new Subject<AdResponse>();
  readonly createdAd$ = this.createdAdSubject.asObservable();

  list(): Observable<AdResponse[]> {
    return this.http.get<AdResponse[]>('/api/ads');
  }

  create(request: CreateAdRequest): Observable<AdResponse> {
    return this.http
      .post<AdResponse>('/api/ads', request)
      .pipe(tap((ad) => this.createdAdSubject.next(ad)));
  }
}
