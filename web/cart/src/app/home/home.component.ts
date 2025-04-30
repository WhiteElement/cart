import { Component, OnInit } from '@angular/core';
import { ShoppingList } from '../models/shopping-list';
import { ShoppinglistService } from '../shoppinglist.service';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit {
  allLists: ShoppingList[] | null;
  toggled = false;
  newListInput: string = '';

  constructor(private shoppingListService: ShoppinglistService) {
    this.allLists = null;
  }

  ngOnInit() {
    this.shoppingListService.getAll().subscribe(res => {
      const statusCode = res.status.toString();
      if (statusCode.startsWith("2")) {
        this.allLists = res.body;
      }
      else {
        console.error("Could not get shoppinglists", res.body);
      }
    });
  }

  newListToggle(): void {
    if (this.toggled) {
      this.toggled = false;
    } else {
      this.toggled = true;
    }
  }

  createNewList(): void {
    if (this.newListInput === '') {
      return;
    }

    this.shoppingListService.createNew(this.newListInput).subscribe(res => {
      const statusCode = res.status.toString();
      if (statusCode.startsWith("2")) {
        this.newListInput = '';
        this.toggled = false;

        this.ngOnInit();
      }
      else {
        console.error("Error creating new List", res.body)
      }
    });
  }
}
