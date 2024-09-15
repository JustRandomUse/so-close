// import { Injectable } from '@angular/core';
// import { HttpClient, HttpParams } from '@angular/common/http';
// import { Observable } from 'rxjs';

// @Injectable({
//   providedIn: 'root'
// })
// export class EmployeesService {
//   private apiUrl = 'http://localhost:8080/employees'; // Ваш сервер

//   constructor(private http: HttpClient) {}

//   // Получение списка сотрудников с пагинацией и фильтрацией
//   getEmployees(page: number, size: number, searchQuery: string = ''): Observable<any> {
//     let params = new HttpParams()
//       .set('page', page.toString())
//       .set('size', size.toString());

//     if (searchQuery) {
//       params = params.set('search', searchQuery);
//     }

//     return this.http.get<any>(this.apiUrl, { params });
//   }

//   // Создание нового сотрудника
//   addEmployee(employee: any): Observable<any> {
//     return this.http.post<any>(this.apiUrl, employee);
//   }

//   // Удаление сотрудника
//   deleteEmployee(id: number): Observable<any> {
//     return this.http.delete<any>(`${this.apiUrl}/${id}`);
//   }

//   // Обновление данных сотрудника
//   updateEmployee(id: number, employee: any): Observable<any> {
//     return this.http.put<any>(`${this.apiUrl}/${id}`, employee);
//   }
// }