import {HttpErrorResponse} from '@angular/common/http';
import {Component, inject, signal} from '@angular/core';
import {FormBuilder, ReactiveFormsModule, Validators} from '@angular/forms';

import {AdsService} from '../../ads/ads.service';

@Component({
  selector: 'app-create-ad-form',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './create-ad-form.html',
  styleUrl: './create-ad-form.css',
})
export class CreateAdForm {
  private readonly fb = inject(FormBuilder);
  private readonly adsService = inject(AdsService);

  protected readonly submitting = signal(false);
  protected readonly errorMessage = signal<string | null>(null);
  protected readonly success = signal(false);

  private successTimeoutId?: ReturnType<typeof setTimeout>;

  protected readonly form = this.fb.nonNullable.group({
    title: ['', [Validators.required]],
    description: [''],
    price: [0, [Validators.required, Validators.min(0.01)]],
  });

  submit(): void {
    if (this.form.invalid || this.submitting()) {
      this.form.markAllAsTouched();
      return;
    }

    this.submitting.set(true);
    this.errorMessage.set(null);
    this.success.set(false);

    this.adsService.create(this.form.getRawValue()).subscribe({
      next: () => {
        this.submitting.set(false);
        this.success.set(true);
        this.form.reset({title: '', description: '', price: 0});

        clearTimeout(this.successTimeoutId);
        this.successTimeoutId = setTimeout(() => this.success.set(false), 1500);
      },
      error: (err: HttpErrorResponse) => {
        this.submitting.set(false);
        this.errorMessage.set(
          err.error?.error ??
            'Не удалось создать объявление, попробуйте еще раз',
        );
      },
    });
  }
}
