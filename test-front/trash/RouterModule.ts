const routes: Routes = [
    { path: 'employees', component: EmployeesComponent },
    { path: 'equipment', component: EquipmentComponent },
    { path: 'equipment-types', component: EquipmentTypesComponent },
    { path: 'employee-equipment', component: EmployeeEquipmentComponent },
    { path: 'inventory', component: InventoryComponent },
    { path: '', redirectTo: 'employees', pathMatch: 'full' }
  ];
  
  @NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
  })
  export class AppRoutingModule { }