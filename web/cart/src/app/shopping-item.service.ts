import { HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ShoppingItem } from './models/shopping-item';

@Injectable({
  providedIn: 'root'
})
export class ShoppingItemService {
  baseUrl = '/api/shoppingitem';

  constructor(private http: HttpClient) { }

  newItem(name: string, listId: number, checked: boolean): Observable<HttpResponse<string>> {
    const item: ShoppingItem = { Id: null, Name: name, ListId: listId, Checked: checked, Updated: null };

    return this.http.post(this.baseUrl, JSON.stringify(item), {
      headers: {
        'Content-Type': 'application/json'
      },
      responseType: 'text',
      observe: 'response'
    });
  }

  patchItem(item: ShoppingItem): Observable<HttpResponse<string>> {
    return this.http.patch(this.baseUrl, JSON.stringify(item), {
      headers: {
        'Content-Type': 'application/json'
      },
      responseType: 'text',
      observe: 'response'
    });
  }

}
