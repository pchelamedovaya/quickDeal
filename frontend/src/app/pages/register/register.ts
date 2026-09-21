import {Component, inject, signal} from '@angular/core';
import {FormBuilder, ReactiveFormsModule, Validators} from '@angular/forms';
import {Router, RouterLink} from '@angular/router';
import {switchMap} from 'rxjs';

import {AuthService} from '../../auth/auth.service';
import {apiErrorKey} from '../../localization/api-error';
import {TranslatePipe} from '../../localization/translate.pipe';
import {TranslationKey} from '../../localization/translation-keys';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [ReactiveFormsModule, RouterLink, TranslatePipe],
  templateUrl: './register.html',
  styleUrl: './register.css',
})
export class Register {
  private readonly fb = inject(FormBuilder);
  private readonly authService = inject(AuthService);
  private readonly router = inject(Router);

  protected readonly submitting = signal(false);
  protected readonly errorKey = signal<TranslationKey | null>(null);

  protected readonly form = this.fb.nonNullable.group({
    email: ['', [Validators.required, Validators.email]],
    username: ['', [Validators.required]],
    password: ['', [Validators.required, Validators.minLength(6)]],
  });

  submit(): void {
    if (this.form.invalid || this.submitting()) {
      this.form.markAllAsTouched();
      return;
    }

    this.submitting.set(true);
    this.errorKey.set(null);

    const {email, password} = this.form.getRawValue();

    this.authService
      .register(this.form.getRawValue())
      .pipe(switchMap(() => this.authService.login({email, password})))
      .subscribe({
        next: () => {
          this.submitting.set(false);
          void this.router.navigateByUrl('/home');
        },
        error: (err: unknown) => {
          this.submitting.set(false);
          this.errorKey.set(apiErrorKey(err));
        },
      });
  }
}
