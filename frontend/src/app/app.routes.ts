import {Routes} from '@angular/router';

import {authGuard} from './auth/auth.guard';
import {Home} from './home/home';
import {Login} from './login/login';
import {Register} from './register/register';

export const routes: Routes = [
  {path: '', redirectTo: 'register', pathMatch: 'full'},
  {path: 'register', component: Register},
  {path: 'login', component: Login},
  {path: 'home', component: Home, canActivate: [authGuard]},
];
