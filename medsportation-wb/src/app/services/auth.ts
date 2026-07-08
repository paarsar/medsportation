import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, tap } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private apiUrl = environment.apiUrl;
  private tokenKey = 'auth_token';

  constructor(private http: HttpClient) { }

  login(credentials: any): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/login`, credentials).pipe(
      tap(res => {
        if (res.token) {
          localStorage.setItem(this.tokenKey, res.token);
        }
      })
    );
  }

  logout(): void {
    localStorage.removeItem(this.tokenKey);
  }

  getToken(): string | null {
    return localStorage.getItem(this.tokenKey);
  }

  isLoggedIn(): boolean {
    return !!this.getToken();
  }

  private getAuthHeaders(): HttpHeaders {
    return new HttpHeaders().set('Authorization', `Bearer ${this.getToken()}`);
  }

  // User Management
  getAllUsers(): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/admin/users`, { headers: this.getAuthHeaders() });
  }

  createUser(user: any): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/admin/users`, user, { headers: this.getAuthHeaders() });
  }

  deleteUser(id: number): Observable<any> {
    return this.http.delete<any>(`${this.apiUrl}/admin/users/${id}`, { headers: this.getAuthHeaders() });
  }
}
