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

  createNew(name: string): Observable<HttpResponse<string>> {
    const list: ShoppingList = {
      Name: name,
      Id: null,
      Items: null,
      Archived: null,
      Created: null,
      Updated: null,
    };

    return this.http.post(this.baseUrl, JSON.stringify(list), {
      headers: {
        'Content-Type': 'application/json',
      },
      responseType: 'text',
      observe: 'response',
    });
  }
}
