import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { QuoteService } from '../../services/quote';

@Component({
  selector: 'app-request-quote',
  standalone: true,
  imports: [RouterLink, ReactiveFormsModule, CommonModule],
  templateUrl: './request-quote.html',
  styleUrl: './request-quote.css'
})
export class RequestQuoteComponent {
  quoteForm: FormGroup;
  isSubmitting = false;
  submitSuccess = false;
  submitError = false;

  constructor(private fb: FormBuilder, private quoteService: QuoteService) {
    this.quoteForm = this.fb.group({
      organizationName: ['', Validators.required],
      organizationType: ['Hospital / Medical Center', Validators.required],
      contactPerson: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      phone: ['', Validators.required],
      serviceType: ['Stat / Rush', Validators.required],
      pickupAddress: ['', Validators.required],
      deliveryAddress: ['', Validators.required],
      specialRequirements: this.fb.group({
        tempControlled: [false],
        chainOfCustody: [false],
        hazmat: [false],
        waitAndReturn: [false]
      }),
      additionalNotes: ['']
    });
  }

  onSubmit() {
    if (this.quoteForm.valid) {
      this.isSubmitting = true;
      this.submitSuccess = false;
      this.submitError = false;

      // Extract special requirements into a list of strings
      const requirements = this.quoteForm.value.specialRequirements;
      const requirementsList = [];
      if (requirements.tempControlled) requirementsList.push('Temperature Controlled');
      if (requirements.chainOfCustody) requirementsList.push('Chain of Custody');
      if (requirements.hazmat) requirementsList.push('Hazmat');
      if (requirements.waitAndReturn) requirementsList.push('Wait & Return');

      const payload = {
        ...this.quoteForm.value,
        specialRequirements: requirementsList
      };

      this.quoteService.submitQuote(payload).subscribe({
        next: (response) => {
          this.isSubmitting = false;
          this.submitSuccess = true;
          this.quoteForm.reset({
            organizationType: 'Hospital / Medical Center',
            serviceType: 'Stat / Rush'
          });
        },
        error: (error) => {
          this.isSubmitting = false;
          this.submitError = true;
          console.error('Error submitting quote:', error);
        }
      });
    } else {
      this.markFormGroupTouched(this.quoteForm);
    }
  }

  private markFormGroupTouched(formGroup: FormGroup) {
    Object.values(formGroup.controls).forEach(control => {
      control.markAsTouched();
      if ((control as any).controls) {
        this.markFormGroupTouched(control as FormGroup);
      }
    });
  }
}
