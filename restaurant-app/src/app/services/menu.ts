import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Menu {
  id: number;
  name: string;
  image_url: string;
  description: string;
  price: number;
  category: string;
  available: boolean;
}

@Injectable({
  providedIn: 'root'
})
export class MenuService {
  private apiUrl = '/api/v1/menu';

  constructor(private http: HttpClient) { }

  getMenus(): Observable<Menu[]> {
    return this.http.get<Menu[]>(this.apiUrl);
  }

  createMenu(menuData: Partial<Menu>): Observable<Menu> {
    return this.http.post<Menu>(this.apiUrl, menuData);
  }

  updateMenu(id: number, data: any): Observable<any> {
    return this.http.patch(`${this.apiUrl}/${id}`, data);
  }


  deleteMenu(id: number): Observable<any> {
    return this.http.delete(`${this.apiUrl}/${id}`);
  }
}
