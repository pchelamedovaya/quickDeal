import {inject, Pipe, PipeTransform} from '@angular/core';

import {LocalizationService} from './localization.service';
import {TranslationKey} from './translation-keys';

@Pipe({name: 'translate', standalone: true, pure: false})
export class TranslatePipe implements PipeTransform {
  private readonly localization = inject(LocalizationService);

  transform(key: TranslationKey): string {
    return this.localization.translate(key);
  }
}
