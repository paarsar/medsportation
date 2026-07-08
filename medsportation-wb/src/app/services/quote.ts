import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { AuthService } from './auth';

@Injectable({
  providedIn: 'root'
})
export class QuoteService {
  private apiUrl = `${environment.apiUrl}/quote`;

  constructor(private http: HttpClient, private authService: AuthService) { }

  private getAuthHeaders(): HttpHeaders {
    return new HttpHeaders().set('Authorization', `Bearer ${this.authService.getToken()}`);
  }

  submitQuote(data: any): Observable<any> {
    return this.http.post(`${environment.apiUrl}/quote`, data);
  }

  submitConsultation(data: any): Observable<any> {
    return this.http.post(`${environment.apiUrl}/consultation`, data);
  }

  getAllQuotes(): Observable<any[]> {
    return this.http.get<any[]>(`${environment.apiUrl}/admin/quotes`, { headers: this.getAuthHeaders() });
  }

  getAllConsultations(): Observable<any[]> {
    return this.http.get<any[]>(`${environment.apiUrl}/admin/consultations`, { headers: this.getAuthHeaders() });
  }

  deleteQuote(id: number): Observable<any> {
    return this.http.delete(`${environment.apiUrl}/admin/quotes/${id}`, { headers: this.getAuthHeaders() });
  }
}
