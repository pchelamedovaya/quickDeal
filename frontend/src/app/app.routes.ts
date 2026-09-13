import {Routes} from '@angular/router';

import {authGuard} from './auth/auth.guard';
import {Home} from './pages/home/home';
import {Login} from './pages/login/login';
import {Register} from './pages/register/register';

export const routes: Routes = [
  {path: '', redirectTo: 'register', pathMatch: 'full'},
  {path: 'register', component: Register},
  {path: 'login', component: Login},
  {path: 'home', component: Home, canActivate: [authGuard]},
];
