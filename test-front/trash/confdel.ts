confirmDelete(employeeId: number) {
    const dialogRef = this.dialog.open(DeleteConfirmationDialog, {
      width: '250px',
    });
  
    dialogRef.afterClosed().subscribe(result => {
      if (result === true) {
        this.employeeService.deleteEmployee(employeeId).subscribe(() => this.loadEmployees());
      }
    });
  }