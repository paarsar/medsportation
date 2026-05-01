import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class QuoteService {
  // Update this URL once your Cloud Function is deployed
  private apiUrl = 'http://localhost:8080/RequestQuote';

  constructor(private http: HttpClient) { }

  submitQuote(data: any): Observable<any> {
    return this.http.post(this.apiUrl, data);
  }
}
