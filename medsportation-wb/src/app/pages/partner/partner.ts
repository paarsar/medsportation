import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { QuoteService } from '../../services/quote';

@Component({
  selector: 'app-partner',
  standalone: true,
  imports: [RouterLink, ReactiveFormsModule, CommonModule],
  templateUrl: './partner.html',
  styleUrl: './partner.css'
})
export class PartnerComponent {
  consultationForm: FormGroup;
  isSubmitting = false;
  submitSuccess = false;
  submitError = false;

  constructor(private fb: FormBuilder, private quoteService: QuoteService) {
    this.consultationForm = this.fb.group({
      name: ['', Validators.required],
      organization: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      phone: ['', Validators.required],
      interestedService: ['Route Optimization Audit', Validators.required],
      message: ['']
    });
  }

  onSubmit() {
    if (this.consultationForm.valid) {
      this.isSubmitting = true;
      this.submitSuccess = false;
      this.submitError = false;

      this.quoteService.submitConsultation(this.consultationForm.value).subscribe({
        next: (response) => {
          this.isSubmitting = false;
          this.submitSuccess = true;
          this.consultationForm.reset({
            interestedService: 'Route Optimization Audit'
          });
        },
        error: (error) => {
          this.isSubmitting = false;
          this.submitError = true;
          console.error('Error submitting consultation:', error);
        }
      });
    } else {
      this.markFormGroupTouched(this.consultationForm);
    }
  }

  private markFormGroupTouched(formGroup: FormGroup) {
    Object.values(formGroup.controls).forEach(control => {
      control.markAsTouched();
    });
  }
}
