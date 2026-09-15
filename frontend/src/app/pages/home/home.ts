import {Component, inject, signal} from '@angular/core';
import {Router} from '@angular/router';

import {AuthService} from '../../auth/auth.service';
import {CreateAdModal} from '../../components/create-ad-modal/create-ad-modal';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CreateAdModal],
  templateUrl: './home.html',
  styleUrl: './home.css',
})
export class Home {
  private readonly authService = inject(AuthService);
  private readonly router = inject(Router);

  protected readonly showCreateAdModal = signal(false);

  logout(): void {
    this.authService.logout().subscribe({
      next: () => this.router.navigateByUrl('/login'),
      error: () => this.router.navigateByUrl('/login'),
    });
  }
}
