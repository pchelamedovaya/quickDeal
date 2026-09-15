import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';

import {AdResponse, CreateAdRequest} from './ad.models';

@Injectable({providedIn: 'root'})
export class AdsService {
  constructor(private readonly http: HttpClient) {}

  create(request: CreateAdRequest): Observable<AdResponse> {
    return this.http.post<AdResponse>('/api/ads', request);
  }
}
