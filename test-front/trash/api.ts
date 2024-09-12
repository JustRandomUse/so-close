@Injectable({
  providedIn: 'root',
})
export class EmployeeService {
  private apiUrl = 'http://your-api-url/employees';

  constructor(private http: HttpClient) {}

  getEmployees() {
    return this.http.get<Employee[]>(this.apiUrl);
  }

  addEmployee(employee: Employee) {
    return this.http.post(this.apiUrl, employee);
  }

  editEmployee(employee: Employee) {
    return this.http.put(`${this.apiUrl}/${employee.id}`, employee);
  }

  deleteEmployee(id: number) {
    return this.http.delete(`${this.apiUrl}/${id}`);
  }
}