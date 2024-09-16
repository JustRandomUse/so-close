// import { Component, OnInit, Input } from '@angular/core';
// import { Router, ActivatedRoute }  from '@angular/router';
// import { Employees } from '../../../employees';
// import { EmloyeesService } from '../../../employees.service';

// @Component({
//   selector: 'app-employees-details',
//   standalone: true,
//   templateUrl: './employees-details.component.html',
//   styleUrl: './employees-details.component.css'
// })
// export class EmployeesDetailsComponent implements OnInit {

//   id!: string;
//   employee!: Employees;

//   constructor(private route: ActivatedRoute,
//       private router: Router,
//         private EmloyeesService: EmloyeesService) { }

//   ngOnInit() {
//     this.employee = new Employees();

//     this.id = this.route.snapshot.params['id'];
//     console.log("gh"+this.id);
//     this.EmloyeesService.getUser(this.id)
//       .subscribe(data => {
//         console.log(data)
//         this.employee = data;
//       }, error => console.log(error));
//   }

//   list(){
//     this.router.navigate(['users']);
//   }

// }
