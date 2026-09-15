import {HttpErrorResponse} from '@angular/common/http';
import {Component, inject, output, signal} from '@angular/core';
import {FormBuilder, ReactiveFormsModule, Validators} from '@angular/forms';

import {AdsService} from '../../ads/ads.service';

@Component({
  selector: 'app-create-ad-modal',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './create-ad-modal.html',
  styleUrl: './create-ad-modal.css',
})
export class CreateAdModal {
  private readonly fb = inject(FormBuilder);
  private readonly adsService = inject(AdsService);

  readonly close = output<void>();

  protected readonly submitting = signal(false);
  protected readonly errorMessage = signal<string | null>(null);
  protected readonly success = signal(false);

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

    this.adsService.create(this.form.getRawValue()).subscribe({
      next: () => {
        this.submitting.set(false);
        this.success.set(true);
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
