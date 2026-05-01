import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AreaCompliance } from './area-compliance';

describe('AreaCompliance', () => {
  let component: AreaCompliance;
  let fixture: ComponentFixture<AreaCompliance>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AreaCompliance]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AreaCompliance);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
