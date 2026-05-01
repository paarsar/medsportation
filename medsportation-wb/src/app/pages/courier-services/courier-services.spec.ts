import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CourierServices } from './courier-services';

describe('CourierServices', () => {
  let component: CourierServices;
  let fixture: ComponentFixture<CourierServices>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CourierServices]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CourierServices);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
