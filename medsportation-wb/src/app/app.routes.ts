import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home';
import { SportsMedicineComponent } from './pages/sports-medicine/sports-medicine';
import { RequestQuoteComponent } from './pages/request-quote/request-quote';
import { AreaComplianceComponent } from './pages/area-compliance/area-compliance';
import { CourierServicesComponent } from './pages/courier-services/courier-services';
import { PartnerComponent } from './pages/partner/partner';
import { AdminComponent } from './pages/admin/admin';
import { LoginComponent } from './pages/login/login';
import { authGuard } from './guards/auth.guard';

export const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'sports-medicine', component: SportsMedicineComponent },
  { path: 'request-quote', component: RequestQuoteComponent },
  { path: 'area-compliance', component: AreaComplianceComponent },
  { path: 'courier-services', component: CourierServicesComponent },
  { path: 'partner', component: PartnerComponent },
  { path: 'admin', component: AdminComponent, canActivate: [authGuard] },
  { path: 'login', component: LoginComponent },
  { path: '**', redirectTo: '' }
];
