import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CustomerMenu } from './customer-menu';

describe('CustomerMenu', () => {
  let component: CustomerMenu;
  let fixture: ComponentFixture<CustomerMenu>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CustomerMenu]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CustomerMenu);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
