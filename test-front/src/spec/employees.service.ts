// import { Injectable } from '@angular/core';
// import { HttpClient } from '@angular/common/http';
// import { Observable } from 'rxjs';

// @Injectable({
//   providedIn: 'root'
// })
// export class EmloyeesService {
//   updateEmployee(id: string, employee: import("./app/employees/employees").Employees) {
//     throw new Error('Method not implemented.');
//   }
//   getEmployee(id: string) {
//     throw new Error('Method not implemented.');
//   }

//   private baseUrl = 'http://localhost:8080/employees';

//   constructor(private http: HttpClient) { }

//   getUser(id: string): Observable<any> {
//     return this.http.get(`${this.baseUrl}/${id}`);
//   }

//   createUser(employee: Object): Observable<Object> {
//     return this.http.post(`${this.baseUrl}`, employee);
//   }

//   updateUser(id: string, value: any): 
//       Observable<Object> {
//     return this.http.put(`${this.baseUrl}/${id}`, 
//        value);
//   }

//   deleteUser(id: string): Observable<any> {
//     return this.http.delete(`${this.baseUrl}/${id}`, 
//        { responseType: 'text' });
//   }

//   getUsersList(): Observable<any> {
//     return this.http.get(`${this.baseUrl}`);
//   }
// }