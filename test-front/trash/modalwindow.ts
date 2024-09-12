openAddDialog() {
    const dialogRef = this.dialog.open(EmployeeFormDialog, {
      width: '400px',
      data: { employee: new Employee() }
    });
  
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.employeeService.addEmployee(result).subscribe(() => this.loadEmployees());
      }
    });
  }