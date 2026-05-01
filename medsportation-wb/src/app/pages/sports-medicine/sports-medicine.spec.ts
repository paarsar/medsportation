import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SportsMedicine } from './sports-medicine';

describe('SportsMedicine', () => {
  let component: SportsMedicine;
  let fixture: ComponentFixture<SportsMedicine>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SportsMedicine]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SportsMedicine);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
