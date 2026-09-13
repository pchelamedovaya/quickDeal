import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable, of} from 'rxjs';
import {tap} from 'rxjs/operators';

import {
  LoginRequest,
  RefreshTokenRequest,
  RegisterRequest,
  TokenResponse,
  UserResponse,
} from './auth.models';
import {TokenStorageService} from './token-storage.service';

@Injectable({providedIn: 'root'})
export class AuthService {
  constructor(
    private readonly http: HttpClient,
    private readonly tokenStorage: TokenStorageService,
  ) {}

  register(request: RegisterRequest): Observable<UserResponse> {
    return this.http.post<UserResponse>('/api/auth/register', request);
  }

  login(request: LoginRequest): Observable<TokenResponse> {
    return this.http
      .post<TokenResponse>('/api/auth/login', request)
      .pipe(tap((tokens) => this.storeTokens(tokens)));
  }

  refresh(): Observable<TokenResponse> {
    const request: RefreshTokenRequest = {
      refresh_token: this.tokenStorage.getRefreshToken() ?? '',
    };
    return this.http
      .post<TokenResponse>('/api/auth/refresh', request)
      .pipe(tap((tokens) => this.storeTokens(tokens)));
  }

  logout(): Observable<void> {
    const refreshToken = this.tokenStorage.getRefreshToken();
    this.tokenStorage.clear();

    if (!refreshToken) {
      return of(undefined);
    }

    const request: RefreshTokenRequest = {refresh_token: refreshToken};
    return this.http.post<void>('/api/auth/logout', request);
  }

  isLoggedIn(): boolean {
    return !!this.tokenStorage.getAccessToken();
  }

  private storeTokens(tokens: TokenResponse): void {
    this.tokenStorage.setTokens(tokens.access_token, tokens.refresh_token);
  }
}
