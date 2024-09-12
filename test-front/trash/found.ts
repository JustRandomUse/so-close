search() {
    this.employeeService.getEmployees().subscribe(data => {
      this.employees = data.filter(employee => employee.name.includes(this.searchQuery));
    });
  }