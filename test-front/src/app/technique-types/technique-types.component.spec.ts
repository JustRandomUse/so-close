import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TechniqueTypesComponent } from './technique-types.component';

describe('TechniqueTypesComponent', () => {
  let component: TechniqueTypesComponent;
  let fixture: ComponentFixture<TechniqueTypesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TechniqueTypesComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TechniqueTypesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
