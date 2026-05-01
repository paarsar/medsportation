import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { QuoteService } from '../../services/quote';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [RouterLink, ReactiveFormsModule, CommonModule],
  templateUrl: './home.html',
  styleUrl: './home.css'
})
export class HomeComponent {
  quickQuoteForm: FormGroup;
  isSubmitting = false;
  submitSuccess = false;
  submitError = false;

  constructor(private fb: FormBuilder, private quoteService: QuoteService) {
    this.quickQuoteForm = this.fb.group({
      organizationName: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      serviceType: ['Medical Courier', Validators.required]
    });
  }

  onQuickSubmit() {
    if (this.quickQuoteForm.valid) {
      this.isSubmitting = true;
      this.submitSuccess = false;
      this.submitError = false;

      this.quoteService.submitQuote(this.quickQuoteForm.value).subscribe({
        next: () => {
          this.isSubmitting = false;
          this.submitSuccess = true;
          this.quickQuoteForm.reset({
            serviceType: 'Medical Courier'
          });
        },
        error: (error) => {
          this.isSubmitting = false;
          this.submitError = true;
          console.error('Error submitting quick quote:', error);
        }
      });
    } else {
      Object.values(this.quickQuoteForm.controls).forEach(control => {
        control.markAsTouched();
      });
    }
  }
}
