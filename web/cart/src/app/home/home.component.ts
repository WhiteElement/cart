import { Component, OnInit, computed, inject, Signal, effect } from '@angular/core';
import { ShoppingList } from '../models/shopping-list';
import { ShoppinglistService } from '../shoppinglist.service';
import { FormsModule } from '@angular/forms';
import { DatePipe } from '@angular/common';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [FormsModule, DatePipe],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  private router = inject(Router);
  private shoppingListService = inject(ShoppinglistService);

  private allLists = this.shoppingListService.allLists;
  activeLists = computed(() => { return this.allLists()?.filter(l => !l.Archived); });
  archivedLists = computed(() => { return this.allLists()?.filter(l => l.Archived); });
  //toggled = false;
  newListInput: string = '';


  //ngOnInit() {
  //  this.shoppingListService.getAll().subscribe(res => {
  //    const statusCode = res.status.toString();
  //    if (statusCode.startsWith("2")) {
  //      if (res.body) {
  //        this.activeLists = res.body.filter(l => !l.Archived)
  //        this.archivedLists = res.body.filter(l => l.Archived)
  //      }
  //    }
  //    else {
  //      console.error("Could not get shoppinglists", res.body);
  //    }
  //  });
  //}

  //newListToggle(): void {
  //  if (this.toggled) {
  //    this.toggled = false;
  //  } else {
  //    this.toggled = true;
  //  }
  //}

  createNewList(): void {
    if (this.newListInput === '') {
      return;
    }

    this.shoppingListService.createNew(this.newListInput).subscribe(res => {
      const statusCode = res.status.toString();
      if (statusCode.startsWith("2")) {
        this.newListInput = '';
        //this.toggled = false;

        //this.ngOnInit();
      }
      else {
        console.error("Error creating new List", res.body)
      }
    });
  }

  // archive = delete
  archivedList(id: number | null): void {
    const archivedList: ShoppingList = { Id: id, Archived: true, Name: null, Created: null, Updated: null, Items: null };
    this.shoppingListService.patch(archivedList).subscribe(res => {
      const statusCode = res.status.toString();
      if (statusCode.startsWith("2")) {
        //this.ngOnInit();
      }
      else {
        console.error("Error creating new List", res.body)
      }
    });
  }

  toDetailsPage(id: number | null): void {
    this.router.navigate(['/list', id]);
  }
}
