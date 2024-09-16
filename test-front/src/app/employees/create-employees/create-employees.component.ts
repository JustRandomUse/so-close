// import { Component, OnInit } from '@angular/core';
// import { Router } from '@angular/router';
// import { Employees } from '../employees';
// import { EmloyeesService } from '../../../employees.service';

// @Component({
//   selector: 'app-create-employees',
//   standalone: true,
//   templateUrl: './create-employees.component.html',
//   styleUrl: './create-employees.component.css'
// })
// export class CreateEmployeesComponent implements OnInit{

//   employee: Employees = new Employees();
//   submitted = false;

//   constructor(private EmloyeesService: EmloyeesService,
//     private router: Router) { }

//   ngOnInit() {
//   }

//   newUser(): void {
//     this.submitted = false;
//     this.employee = new Employees();
//   }

//   save() {
//     this.EmloyeesService
//     .createUser(this.employee).subscribe(data => {
//       console.log(data)
//       this.employee = new Employees();
//       this.gotoList();
//     }, 
//     error => console.log(error));
//   }

//   onSubmit() {
//     this.submitted = true;
//     this.save();    
//   }

//   gotoList() {
//     this.router.navigate(['/employees']);
//   }

// }
