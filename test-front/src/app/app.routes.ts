import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { EmployeesComponent } from './employees/employees.component';
import { TechniquesComponent } from './techniques/techniques.component';
import { TechniqueTypesComponent } from './technique-types/technique-types.component';
import { EmployeeEquipmentComponent } from './employee-equipment/employee-equipment.component';

const routes: Routes = [
  { path: 'employees', component: EmployeesComponent },
  { path: 'employee_equipment', component: EmployeeEquipmentComponent },
  { path: 'technique_types', component: TechniqueTypesComponent },
  { path: 'techniques', component: TechniquesComponent },
  // { path: '', redirectTo: '/employees', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }