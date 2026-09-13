import {HttpErrorResponse, HttpInterceptorFn} from '@angular/common/http';
import {inject} from '@angular/core';
import {Router} from '@angular/router';
import {catchError, switchMap, throwError} from 'rxjs';

import {AuthService} from './auth.service';
import {TokenStorageService} from './token-storage.service';

function isAuthEndpoint(url: string): boolean {
  return url.includes('/api/auth/');
}

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const tokenStorage = inject(TokenStorageService);
  const authService = inject(AuthService);
  const router = inject(Router);

  const accessToken = tokenStorage.getAccessToken();
  const authorizedReq =
    accessToken && !isAuthEndpoint(req.url)
      ? req.clone({setHeaders: {Authorization: `Bearer ${accessToken}`}})
      : req;

  return next(authorizedReq).pipe(
    catchError((error: unknown) => {
      if (
        !(error instanceof HttpErrorResponse) ||
        error.status !== 401 ||
        isAuthEndpoint(req.url)
      ) {
        return throwError(() => error);
      }

      return authService.refresh().pipe(
        switchMap((tokens) =>
          next(
            req.clone({
              setHeaders: {Authorization: `Bearer ${tokens.access_token}`},
            }),
          ),
        ),
        catchError((refreshError: unknown) => {
          tokenStorage.clear();
          void router.navigateByUrl('/login');
          return throwError(() => refreshError);
        }),
      );
    }),
  );
};
