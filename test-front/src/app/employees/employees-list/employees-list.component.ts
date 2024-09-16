// import { Component, OnInit } from '@angular/core';
// import { Observable } from "rxjs";
// import { Router } from '@angular/router';
// import { Employees } from '../../../employees';
// import { EmloyeesService } from '../../../employees.service';

// @Component({
//   selector: 'app-employees-list',
//   standalone: true,
//   templateUrl: './employees-list.component.html',
//   styleUrl: './employees-list.component.css'
// })
// export class EmployeesListComponent implements OnInit {
  
//   employee!: Observable<Employees[]>;

//   constructor(private EmloyeesService: EmloyeesService,
//     private router: Router) {}

//   ngOnInit() {
//     this.reloadData();
//   }

//   reloadData() {
//     this.employee = this.EmloyeesService.getUsersList();
//   }

//   deleteEmployee(id: string) {
  
//     this.EmloyeesService.deleteUser(id)
//       .subscribe(
//         data => {
//           console.log(data);
//           this.reloadData();
//         },
//         error => console.log(error));
//   }
//   updateEmployee(id: string){
//     this.router.navigate(['update', id]);
//   }
//   EmployeeDetails(id: string){
//     this.router.navigate(['details', id]);
//   }

// }
