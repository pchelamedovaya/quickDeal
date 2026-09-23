import {registerLocaleData} from '@angular/common';
import localeEn from '@angular/common/locales/en';
import localeRu from '@angular/common/locales/ru';

export function registerAppLocales(): void {
  registerLocaleData(localeEn, 'en-US');
  registerLocaleData(localeRu, 'ru-RU');
}
