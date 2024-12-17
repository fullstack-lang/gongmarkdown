import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongmarkdownComponent } from './gongmarkdown.component';

describe('GongmarkdownComponent', () => {
  let component: GongmarkdownComponent;
  let fixture: ComponentFixture<GongmarkdownComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongmarkdownComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongmarkdownComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
