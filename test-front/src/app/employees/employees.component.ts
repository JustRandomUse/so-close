// import { Component, OnInit } from '@angular/core';
// import { EmployeesService } from './employees.service.spec'; // Импорт сервиса

// @Component({
//   selector: 'app-employees',
//   templateUrl: './employees.component.html',
//   styleUrls: ['./employees.component.css']
// })
// export class EmployeesComponent implements OnInit {
//   employees: any[] = [];
//   paginatedEmployees: any[] = [];
//   currentPage: number = 1;
//   itemsPerPage: number = 5;
//   totalItems: number = 0;
//   searchQuery: string = '';

//   constructor(private employeesService: EmployeesService) {}

//   ngOnInit() {
//     this.loadEmployees();
//   }

//   loadEmployees() {
//     this.employeesService.getEmployees(this.currentPage, this.itemsPerPage, this.searchQuery).subscribe(response => {
//       this.employees = response.content; // Предположим, что данные приходят в виде "content"
//       this.totalItems = response.totalElements; // Общее количество элементов
//     });
//   }

//   onSearch(event: any) {
//     this.searchQuery = event.target.value;
//     this.loadEmployees();
//   }

//   previousPage() {
//     if (this.currentPage > 1) {
//       this.currentPage--;
//       this.loadEmployees();
//     }
//   }

//   nextPage() {
//     if (this.currentPage * this.itemsPerPage < this.totalItems) {
//       this.currentPage++;
//       this.loadEmployees();
//     }
//   }

//   addEmployee(employee: any) {
//     this.employeesService.addEmployee(employee).subscribe(() => {
//       this.loadEmployees();
//     });
//   }

//   deleteEmployee(employee: any) {
//     this.employeesService.deleteEmployee(employee.id).subscribe(() => {
//       this.loadEmployees();
//     });
//   }

//   editEmployee(employee: any) {
//     this.employeesService.updateEmployee(employee.id, employee).subscribe(() => {
//       this.loadEmployees();
//     });
//   }
// }