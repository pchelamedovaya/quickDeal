import {DOCUMENT} from '@angular/common';
import {computed, effect, inject, Injectable, signal} from '@angular/core';

import {TranslationKey, Translations} from './translation-keys';
import {en} from './translations.en';
import {ru} from './translations.ru';

export type Language = 'en' | 'ru';
export const SUPPORTED_LANGUAGES: readonly Language[] = ['en', 'ru'];

const LANGUAGE_KEY = 'quickdeal_language';
const DEFAULT_LANGUAGE: Language = 'ru';

const DICTIONARIES: Record<Language, Translations> = {en, ru};
const LOCALES: Record<Language, string> = {
  en: 'en-US',
  ru: 'ru-RU',
};

function isLanguage(value: unknown): value is Language {
  return SUPPORTED_LANGUAGES.includes(value as Language);
}

@Injectable({providedIn: 'root'})
export class LocalizationService {
  private readonly document = inject(DOCUMENT);

  private readonly currentLanguage = signal<Language>(this.detectLanguage());
  readonly language = this.currentLanguage.asReadonly();
  readonly locale = computed(() => LOCALES[this.currentLanguage()]);

  constructor() {
    effect(() => {
      this.document.documentElement.lang = this.currentLanguage();
    });
  }

  setLanguage(language: Language): void {
    if (language === this.currentLanguage()) {
      return;
    }

    localStorage.setItem(LANGUAGE_KEY, language);
    this.currentLanguage.set(language);
  }

  private detectLanguage(): Language {
    const saved = localStorage.getItem(LANGUAGE_KEY);
    if (isLanguage(saved)) {
      return saved;
    }

    const browser = navigator.language?.slice(0, 2).toLowerCase();
    if (isLanguage(browser)) {
      return browser;
    }

    return DEFAULT_LANGUAGE;
  }

  translate(key: TranslationKey): string {
    let node: unknown = DICTIONARIES[this.currentLanguage()];

    for (const part of key.split('.')) {
      node = (node as Record<string, unknown>)[part];
    }

    return typeof node === 'string' ? node : key;
  }
}
