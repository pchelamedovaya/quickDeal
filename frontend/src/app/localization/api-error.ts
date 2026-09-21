import {HttpErrorResponse} from '@angular/common/http';

import {ApiErrorCode, TranslationKey} from './translation-keys';
import {en} from './translations.en';

function isApiErrorCode(value: unknown): value is ApiErrorCode {
  return typeof value === 'string' && value in en.errors;
}

export function apiErrorKey(err: unknown): TranslationKey {
  const code =
    err instanceof HttpErrorResponse ? err.error?.error_code : undefined;

  return isApiErrorCode(code) ? `errors.${code}` : 'errors.UNKNOWN';
}
