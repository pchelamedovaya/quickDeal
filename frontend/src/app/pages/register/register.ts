import {HttpErrorResponse} from '@angular/common/http';
import {Component, inject, signal} from '@angular/core';
import {FormBuilder, ReactiveFormsModule, Validators} from '@angular/forms';
import {Router, RouterLink} from '@angular/router';
import {switchMap} from 'rxjs';

import {AuthService} from '../../auth/auth.service';

const ERROR_MESSAGES: Record<string, string> = {
  'email already taken': 'Этот email уже занят',
  'username already taken': 'Это имя пользователя уже занято',
};

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [ReactiveFormsModule, RouterLink],
  templateUrl: './register.html',
  styleUrl: './register.css',
})
export class Register {
  private readonly fb = inject(FormBuilder);
  private readonly authService = inject(AuthService);
  private readonly router = inject(Router);

  protected readonly submitting = signal(false);
  protected readonly errorMessage = signal<string | null>(null);

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
    this.errorMessage.set(null);

    const {email, password} = this.form.getRawValue();

    this.authService.register(this.form.getRawValue())
      .pipe(switchMap(() => this.authService.login({email, password})))
      .subscribe({
        next: () => {
          this.submitting.set(false);
          void this.router.navigateByUrl('/home');
        },
        error: (err: HttpErrorResponse) => {
          this.submitting.set(false);
          const backendMessage: string | undefined = err.error?.error;
          this.errorMessage.set(
            (backendMessage && ERROR_MESSAGES[backendMessage]) ??
            backendMessage ??
            'Не удалось зарегистрироваться, попробуйте еще раз',
          );
        },
      });
  }
}
