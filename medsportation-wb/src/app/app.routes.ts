import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home';
import { SportsMedicineComponent } from './pages/sports-medicine/sports-medicine';
import { RequestQuoteComponent } from './pages/request-quote/request-quote';
import { AreaComplianceComponent } from './pages/area-compliance/area-compliance';
import { CourierServicesComponent } from './pages/courier-services/courier-services';
import { PartnerComponent } from './pages/partner/partner';

export const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'sports-medicine', component: SportsMedicineComponent },
  { path: 'request-quote', component: RequestQuoteComponent },
  { path: 'area-compliance', component: AreaComplianceComponent },
  { path: 'courier-services', component: CourierServicesComponent },
  { path: 'partner', component: PartnerComponent },
  { path: '**', redirectTo: '' }
];
