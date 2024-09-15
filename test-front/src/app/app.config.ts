import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter, Routes } from '@angular/router';
// import { EmployeesComponent } from './employees/employees.component';
import { TechniquesComponent } from './techniques/techniques.component';
import { TechniqueTypesComponent } from './technique-types/technique-types.component';
import { EmployeeEquipmentComponent } from './employee-equipment/employee-equipment.component';
import { InventoryComponent } from './inventory/inventory.component';

export const routes: Routes = [
  // { path: 'employees', component: EmployeesComponent },
  { path: 'techniques', component: TechniquesComponent },
  { path: 'technique-types', component: TechniqueTypesComponent },
  { path: 'employee-equipment', component: EmployeeEquipmentComponent },
  { path: 'inventory', component: InventoryComponent },
  { path: '', redirectTo: '/employees', pathMatch: 'full' },
  { path: '**', redirectTo: '/employees' }
];

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(routes),
  ]
};
