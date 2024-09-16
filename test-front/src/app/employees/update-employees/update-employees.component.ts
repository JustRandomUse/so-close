// import { Component, OnInit} from '@angular/core';
// import { ActivatedRoute, Router } from '@angular/router';
// import { Employees } from '../../../employees';
// import { EmloyeesService } from '../../../employees.service';
// import internal from 'stream';

// @Component({
//   selector: 'app-update-employees',
//   templateUrl: './update-employees.component.html',
//   styleUrl: './update-employees.component.css'
// })
// export class UpdateEmployeesComponent implements OnInit{
//   id!: string;
//   employee!: Employees;
//   updated = false;

//   constructor(private route: ActivatedRoute,
//     private router: Router,
//      private employeesService: EmloyeesService) { }

//   ngOnInit(){
//     this.employee = new Employees();

//     this.id = this.route.snapshot.params['id'];
    
//     this.employeesService.getEmployee(this.id)
//       .subscribe((data: Employees) => {
//         this.employee = data;
//       }, (error: any) => console.log(error));
    
//   }

//   updateEmployee() {
//     this.employeesService.updateEmployee(this.id, this.employee)
//       .subscribe((data: any) => {
//         console.log(data);
//         this.employee = new Employees();
//         this.gotoList();
//       }, (error: any) => console.log(error));
//   }

//   onSubmit() {
//     this.updateEmployee(); 
//     this.updated = true;   
//   }

//   gotoList() {
//     this.router.navigate(['/employees']);
//   }
// }
