import {en} from './translations.en';

export type Translations = typeof en;

type DotPaths<T> = {
  [K in keyof T & string]: T[K] extends string ? K : `${K}.${DotPaths<T[K]>}`;
}[keyof T & string];

export type TranslationKey = DotPaths<Translations>;

export type ApiErrorCode = keyof Translations['errors'];
