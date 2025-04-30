import { Component, OnInit } from '@angular/core';
import { ShoppingList } from '../models/shopping-list';
import { ShoppinglistService } from '../shoppinglist.service';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit {
  allLists: ShoppingList[] | null;

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

}
