import {HttpErrorResponse} from '@angular/common/http';
import {Component, inject, signal} from '@angular/core';
import {FormBuilder, ReactiveFormsModule, Validators} from '@angular/forms';

import {AuthService} from '../auth/auth.service';

const ERROR_MESSAGES: Record<string, string> = {
  'email already taken': 'Этот email уже занят',
  'username already taken': 'Это имя пользователя уже занято',
};

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './register.html',
  styleUrl: './register.css',
})
export class Register {
  private readonly fb = inject(FormBuilder);
  private readonly authService = inject(AuthService);

  protected readonly submitting = signal(false);
  protected readonly successMessage = signal<string | null>(null);
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
    this.successMessage.set(null);
    this.errorMessage.set(null);

    this.authService.register(this.form.getRawValue()).subscribe({
      next: (user) => {
        this.submitting.set(false);
        this.successMessage.set(
          `Пользователь ${user.username} успешно зарегистрирован`,
        );
        this.form.reset();
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
