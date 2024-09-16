import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';
import { EmployeeEquipmentComponent } from './employee-equipment/employee-equipment.component';
import { TechniquesComponent } from './techniques/techniques.component';
import { TechniqueTypesComponent } from './technique-types/technique-types.component';
import { InventoryComponent } from './inventory/inventory.component';
import { RouterOutlet } from '@angular/router';
import { EmployeesComponent } from './employees/employees.component'; // Убедитесь, что этот компонент является standalone

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,
    AppComponent,
    EmployeesComponent,
    EmployeeEquipmentComponent,
    TechniquesComponent,
    TechniqueTypesComponent,
    InventoryComponent
  ],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  selectedTab: string = 'techniques'; 

  selectTab(tab: string) {
    this.selectedTab = tab;
  }
}