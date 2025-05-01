import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ShoppinglistService } from '../shoppinglist.service';
import { ShoppingItemService } from '../shopping-item.service';
import { ShoppingList } from '../models/shopping-list';

@Component({
  selector: 'app-detail',
  standalone: true,
  imports: [],
  templateUrl: './detail.component.html',
  styleUrl: './detail.component.css'
})
export class DetailComponent implements OnInit {
  list: ShoppingList = {} as ShoppingList;

  constructor(private route: ActivatedRoute, private shoppinglistService: ShoppinglistService, private shoppingItemService: ShoppingItemService) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id) {
        this.shoppinglistService.getOne(id).subscribe(res => {
          const statusCode = res.status.toString();
          if (statusCode.startsWith("2")) {
            if (res.body) {
              this.list = res.body;
            }
          }
        })
      }
    });
  }
}
