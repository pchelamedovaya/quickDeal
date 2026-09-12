import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';

import {RegisterRequest, UserResponse} from './auth.models';

@Injectable({providedIn: 'root'})
export class AuthService {
  constructor(private readonly http: HttpClient) {}

  register(request: RegisterRequest): Observable<UserResponse> {
    return this.http.post<UserResponse>('/api/auth/register', request);
  }
}
