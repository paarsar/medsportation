import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { QuoteService } from '../../services/quote';
import { AuthService } from '../../services/auth';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-admin',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './admin.html',
  styleUrl: './admin.css'
})
export class AdminComponent implements OnInit {
  quotes: any[] = [];
  users: any[] = [];
  loading = true;
  error = false;
  view: 'quotes' | 'users' = 'quotes';

  // New User Form
  newUser = { username: '', password: '' };
  userError = '';

  constructor(
    private quoteService: QuoteService,
    private authService: AuthService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.fetchQuotes();
    this.fetchUsers();
  }

  fetchQuotes(): void {
    this.loading = true;
    this.quoteService.getAllQuotes().subscribe({
      next: (data) => {
        this.quotes = data;
        this.loading = false;
      },
      error: (err) => {
        console.error('Error fetching quotes:', err);
        this.error = true;
        this.loading = false;
      }
    });
  }

  deleteQuote(id: number): void {
    if (confirm('Are you sure you want to delete this quote request?')) {
      this.quoteService.deleteQuote(id).subscribe({
        next: () => {
          this.quotes = this.quotes.filter(q => q.id !== id);
        },
        error: (err) => alert('Failed to delete quote')
      });
    }
  }

  exportToCSV(): void {
    if (this.quotes.length === 0) return;

    const headers = ['Date', 'Organization', 'Type', 'Contact', 'Email', 'Phone', 'Service', 'Pickup', 'Delivery', 'Requirements', 'Notes'];
    const rows = this.quotes.map(q => [
      new Date(q.createdAt).toLocaleDateString(),
      `"${q.organizationName.replace(/"/g, '""')}"`,
      q.organizationType,
      `"${q.contactPerson.replace(/"/g, '""')}"`,
      q.email,
      q.phone,
      q.serviceType,
      `"${q.pickupAddress.replace(/"/g, '""')}"`,
      `"${q.deliveryAddress.replace(/"/g, '""')}"`,
      `"${(q.specialRequirements || '').replace(/"/g, '""')}"`,
      `"${(q.additionalNotes || '').replace(/"/g, '""')}"`
    ]);

    const csvContent = [headers.join(','), ...rows.map(r => r.join(','))].join('\n');
    const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
    const link = document.createElement('a');
    const url = URL.createObjectURL(blob);
    
    link.setAttribute('href', url);
    link.setAttribute('download', `medsportation_quotes_${new Date().toISOString().split('T')[0]}.csv`);
    link.style.visibility = 'hidden';
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  }

  // User Management
  fetchUsers(): void {
    this.authService.getAllUsers().subscribe({
      next: (data) => this.users = data,
      error: (err) => console.error('Error fetching users:', err)
    });
  }

  onCreateUser(): void {
    this.userError = '';
    this.authService.createUser(this.newUser).subscribe({
      next: (user) => {
        this.users.push(user);
        this.newUser = { username: '', password: '' };
      },
      error: (err) => this.userError = 'Failed to create user. Username might already exist.'
    });
  }

  deleteUser(id: number): void {
    if (confirm('Are you sure you want to delete this admin user?')) {
      this.authService.deleteUser(id).subscribe({
        next: () => {
          this.users = this.users.filter(u => u.id !== id);
        },
        error: (err) => alert('Failed to delete user')
      });
    }
  }

  logout(): void {
    this.authService.logout();
    this.router.navigate(['/login']);
  }
}
