import {Component, inject} from '@angular/core';

import {
  Language,
  LocalizationService,
  SUPPORTED_LANGUAGES,
} from '../../localization/localization.service';
import {TranslatePipe} from '../../localization/translate.pipe';
import {TranslationKey} from '../../localization/translation-keys';

@Component({
  selector: 'app-language-switcher',
  standalone: true,
  imports: [TranslatePipe],
  templateUrl: './language-switcher.html',
  styleUrl: './language-switcher.css',
})
export class LanguageSwitcher {
  protected readonly localization = inject(LocalizationService);
  protected readonly languages = SUPPORTED_LANGUAGES;

  protected labelKey(language: Language): TranslationKey {
    return `language.${language}`;
  }
}
