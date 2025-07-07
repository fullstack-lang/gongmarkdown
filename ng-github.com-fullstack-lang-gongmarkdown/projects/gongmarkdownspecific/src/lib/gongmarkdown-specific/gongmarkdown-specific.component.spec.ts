import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongmarkdownSpecificComponent } from './gongmarkdown-specific.component';

describe('GongmarkdownSpecificComponent', () => {
  let component: GongmarkdownSpecificComponent;
  let fixture: ComponentFixture<GongmarkdownSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongmarkdownSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongmarkdownSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
