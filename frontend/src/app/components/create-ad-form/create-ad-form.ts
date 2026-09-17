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

  protected readonly form = this.fb.group({
    title: this.fb.nonNullable.control('', [Validators.required]),
    description: this.fb.nonNullable.control(''),
    price: this.fb.control<number | null>(null, [
      Validators.required,
      Validators.min(0.01),
    ]),
  });

  submit(): void {
    if (this.form.invalid || this.submitting()) {
      this.form.markAllAsTouched();
      return;
    }

    this.submitting.set(true);
    this.errorMessage.set(null);

    const {title, description, price} = this.form.getRawValue();

    this.adsService.create({title, description, price: price ?? 0}).subscribe({
      next: () => {
        this.submitting.set(false);
        this.form.reset({title: '', description: '', price: null});
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
