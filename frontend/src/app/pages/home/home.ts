import {CurrencyPipe, DatePipe} from '@angular/common';
import {Component, inject, OnInit, signal} from '@angular/core';
import {Router} from '@angular/router';

import {AdResponse} from '../../ads/ad.models';
import {AdsService} from '../../ads/ads.service';
import {AuthService} from '../../auth/auth.service';
import {CreateAdForm} from '../../components/create-ad-form/create-ad-form';
import {LocalizationService} from '../../localization/localization.service';
import {TranslatePipe} from '../../localization/translate.pipe';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CreateAdForm, CurrencyPipe, DatePipe, TranslatePipe],
  templateUrl: './home.html',
  styleUrl: './home.css',
})
export class Home implements OnInit {
  private readonly authService = inject(AuthService);
  private readonly adsService = inject(AdsService);
  private readonly router = inject(Router);

  protected readonly localization = inject(LocalizationService);
  protected readonly ads = signal<AdResponse[]>([]);

  ngOnInit(): void {
    this.adsService.list().subscribe((ads) => this.ads.set(ads));

    this.adsService.createdAd$.subscribe((ad) => {
      this.ads.update((current) => [ad, ...current]);
    });
  }

  logout(): void {
    this.authService.logout().subscribe({
      next: () => this.router.navigateByUrl('/login'),
      error: () => this.router.navigateByUrl('/login'),
    });
  }
}
