import { Component, OnInit } from '@angular/core';
import { EmployeesService } from './employees.service.spec'; // Импорт сервиса
import { employee } from './employees';

@Component({
  selector: 'app-employees',
  standalone: true,
  templateUrl: './employees.component.html',
  styleUrls: ['./employees.component.css']
})
export class EmployeesComponent implements OnInit {
  employees: employee = new employee;
  submitted = false;
  paginatedEmployees: any[] = [];
  currentPage: number = 1;
  itemsPerPage: number = 5;
  totalItems: number = 0;
  searchQuery: string = ''; 

  constructor(private employeesService: EmployeesService) {}

  ngOnInit() {
    this.loadEmployees();
  }

  onSearch(event: any) {
    this.searchQuery = event.target.value;
    this.loadEmployees();
  }

  loadEmployees() {
    this.employeesService.getEmployees(this.currentPage, this.itemsPerPage, this.searchQuery).subscribe(response => {
      this.employees = response.class; // Предположим, что данные приходят в виде "content"
      this.totalItems = response.totalElements; // Общее количество элементов
    });
  }

  addEmployee(Employees: any) {
    this.employeesService.addEmployee(this.employees).subscribe(() => {
      this.loadEmployees();
    });
    (error: any) => console.log(error);
  }

  editEmployee(employee: any) {
    this.employeesService.updateEmployee(employee.id, employee).subscribe(() => {
      this.loadEmployees();
    });
    (error: any) => console.log(error);
  }

  deleteEmployee(employee: any) {
    this.employeesService.deleteEmployee(employee.id).subscribe(() => {
      this.loadEmployees();
    });
    (error: any) => console.log(error);
  }

  previousPage() {
    if (this.currentPage > 1) {
      this.currentPage--;
      this.loadEmployees();
    }
  }

  nextPage() {
    if (this.currentPage * this.itemsPerPage < this.totalItems) {
      this.currentPage++;
      this.loadEmployees();
    }
  }

}