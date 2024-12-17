import { TestBed } from '@angular/core/testing';

import { GongmarkdownService } from './gongmarkdown.service';

describe('GongmarkdownService', () => {
  let service: GongmarkdownService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongmarkdownService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
