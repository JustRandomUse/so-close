export class EmployeesComponent implements OnInit {
    employees: Employee[] = [];
  
    constructor(private employeeService: EmployeeService) {}
  
    ngOnInit() {
      this.loadEmployees();
    }
  
    loadEmployees() {
      this.employeeService.getEmployees().subscribe(data => {
        this.employees = data;
      });
    }
  }