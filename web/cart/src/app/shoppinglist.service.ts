import { Injectable, inject, signal } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { ShoppingList } from './models/shopping-list';
import { Observable, catchError, of } from 'rxjs';
import { toSignal } from '@angular/core/rxjs-interop';

@Injectable({
  providedIn: 'root'
})
export class ShoppinglistService {
  private http = inject(HttpClient);
  private baseUrl = "api/shoppinglist";

  private getAll(): Observable<ShoppingList[]> {
    return this.http.get<ShoppingList[]>(this.baseUrl);
  }

  readonly allLists = toSignal(this.getAll());
  private currentList = signal<ShoppingList | undefined>(undefined);
  readonly activeList = this.currentList.asReadonly();

  updateList(id: string): void {
    this.http.get<ShoppingList>(`${this.baseUrl}/${id}`).pipe(
      catchError((err) => {
        console.error(`Error getting shoppinglist: ${err}`);
        return of(undefined);
      })
    ).subscribe(res => { this.currentList.set(res) });
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

  patch(list: ShoppingList): Observable<HttpResponse<string>> {
    return this.http.patch(this.baseUrl, JSON.stringify(list), {
      headers: {
        'Content-Type': 'application/json',
      },
      responseType: 'text',
      observe: 'response',
    });
  }
}
