import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { ShoppingList } from './models/shopping-list';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ShoppinglistService {
  private baseUrl = "api/shoppinglist";
  constructor(private http: HttpClient) { }

  getAll(): Observable<HttpResponse<ShoppingList[]>> {
    return this.http.get<ShoppingList[]>(this.baseUrl, {
      observe: 'response'
    });
  }
}
