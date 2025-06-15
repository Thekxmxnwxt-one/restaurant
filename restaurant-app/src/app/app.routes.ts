import { Routes } from '@angular/router';
import { Login } from './pages/login/login';
import { CustomerMenu } from './pages/customer-menu/customer-menu';

export const routes: Routes = [
    {
        path: '',
        component: Login
    },
    {
        path: 'management-menu',
        component: CustomerMenu
    },
];
